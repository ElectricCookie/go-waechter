package waechter

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateRefreshToken generates a new refresh-token and saves it in the database
func (waechter *Waechter) GenerateRefreshToken(userID string, expires int64) (*string, *AuthError) {

	// Generate token

	token := generateRandomString(32)

	// Get the associated user

	user, err := waechter.getDBAdapter().GetUserByID(userID)

	if err != nil {
		return nil, &AuthError{
			ErrorCode:   "unknownUser",
			Description: "Could not find a user with the given id",
			IsInternal:  true,
		}
	}

	// Hash token
	derivedKey := hash(token, user.Salt)

	waechter.getDBAdapter().InsertRefreshToken(&RefreshToken{
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

	resString, err := jwtToken.SignedString([]byte(waechter.JwtSecret))
	if err != nil {
		return nil, cryptError(err)
	}

	return &resString, nil

}

// CheckRefreshToken checks if a refresh token is valid. In case of invalidity a theft is assumed and the users sessions are nuked
func (waechter *Waechter) CheckRefreshToken(jwtToken string) (*jwt.StandardClaims, error) {

	claims := jwt.StandardClaims{}
	_, parseError := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(waechter.JwtSecret), nil
	})

	refreshToken := claims.Id
	userID := claims.Subject

	if parseError != nil {
		return nil, parseError
	}

	user, userErr := waechter.getDBAdapter().GetUserByID(userID)

	if userErr != nil {
		return nil, userErr
	}

	// Hash token

	dk := hash(refreshToken, user.Salt)

	_, retrieveError := waechter.getDBAdapter().FindRefreshToken(userID, string(dk))

	if retrieveError != nil {
		return nil, retrieveError
	}

	return &claims, nil

}

// GenerateAccessToken issues a new access token based on a refresh token
func (waechter *Waechter) GenerateAccessToken(refreshToken string, isAllowed func(*jwt.StandardClaims) bool) (*string, *AuthError) {

	claims, err := waechter.CheckRefreshToken(refreshToken)
	if err != nil {
		return nil, &AuthError{
			ErrorCode:   "invalidRefreshToken",
			Description: "The refresh token passed was invalid",
			IsInternal:  false,
		}
	}

	result, err := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.StandardClaims{
		Issuer:    waechter.JwtIssuer,
		IssuedAt:  time.Now().Unix(),
		Subject:   claims.Subject,
		ExpiresAt: time.Now().Add(0).Unix(),
	}).SignedString(waechter.JwtSecret)

	if err != nil {
		return nil, cryptError(err)
	}

	return &result, nil

}
