package waechter

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
)

// GenerateRefreshToken generates a new refresh-token and saves it in the database
func GenerateRefreshToken(userID string, expires int64) (*string, *AuthError) {

	// Generate token
	token, err := generateRandomString(64)
	if err != nil {
		return nil, &AuthError{
			ErrorCode:   "randomError",
			Description: "Could not generate a random token",
			IsInternal:  true,
		}
	}

	// Get the associated user

	user, err := getAdapter().GetUserById(userID)

	if err != nil {
		return nil, &AuthError{
			ErrorCode:   "userNotFound",
			Description: "Could not find a user with the given id",
			IsInternal:  true,
		}
	}

	// Hash token
	derivedKey, err := scrypt.Key([]byte(token), []byte(user.Salt), 16384, 8, 1, 32)

	if err != nil {
		return nil, cryptError(err)
	}

	getAdapter().InsertRefreshToken(&RefreshToken{
		Expires: expires,
		Token:   string(derivedKey),
		UserID:  userID,
	})

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Issuer:   "das",
		Subject:  userID,
		IssuedAt: time.Now().Unix(),
		Id:       token,
	})

	resString, err := jwtToken.SignedString(getParameters().jwtSecret)
	if err != nil {
		return nil, cryptError(err)
	}

	return &resString, nil

}

// checkRefreshToken checks if a refresh token is valid. In case of invalidity a theft is assumed and the users sessions are nuked
func checkRefreshToken(jwtToken string) (*jwt.StandardClaims, error) {

	claims := jwt.StandardClaims{}
	_, parseError := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(getParameters().jwtSecret), nil
	})

	refreshToken := claims.Id
	userID := claims.Subject

	if parseError != nil {
		return nil, parseError
	}

	user, userErr := getAdapter().GetUserById(userID)

	if userErr != nil {
		return nil, userErr
	}

	// Hash token

	dk, cryptError := scrypt.Key([]byte(refreshToken), []byte(user.Salt), 16384, 8, 1, 32)

	if cryptError != nil {
		return nil, cryptError
	}

	_, retrieveError := getAdapter().FindRefreshToken(userID, string(dk))

	if retrieveError != nil {
		return nil, retrieveError
	}

	return &claims, nil

}

// GenerateAccessToken issues a new access token based on a refresh token
func GenerateAccessToken(refreshToken string) (*string, *AuthError) {

	claims, err := checkRefreshToken(refreshToken)
	if err != nil {
		return nil, &AuthError{
			ErrorCode:   "invalidRefreshToken",
			Description: "The refresh token passed was invalid",
			IsInternal:  false,
		}
	}

	result, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Issuer:    getParameters().jwtIssuer,
		IssuedAt:  time.Now().Unix(),
		Subject:   claims.Subject,
		ExpiresAt: time.Now().Add(getParameters().jwtAccessTokenLifetime).Unix(),
	}).SignedString(getParameters().jwtSecret)

	if err != nil {
		return nil, cryptError(err)
	}

	return &result, nil

}
