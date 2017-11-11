package waechter

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// UserGenerateRefreshToken generates a new refresh-token and saves it in the database
func (waechter *Waechter) UserGenerateRefreshToken(userID string, expires int64) (string, *AuthError) {

	// Generate token

	token := generateRandomString(32)

	// Get the associated user

	user, err := waechter.getDBAdapter().GetUserByID(userID)

	if err != nil {
		return "", &AuthError{
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
		return "", CryptError(err)
	}

	return resString, nil

}

// UserCheckRefreshToken checks if a refresh token is valid. In case of invalidity a theft is assumed and the users sessions are nuked
func (waechter *Waechter) UserCheckRefreshToken(jwtToken string) (*User, *jwt.StandardClaims, error) {

	claims := jwt.StandardClaims{}
	_, parseError := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(waechter.JwtSecret), nil
	})

	refreshToken := claims.Id
	userID := claims.Subject

	if parseError != nil {
		return nil, nil, parseError
	}

	user, userErr := waechter.getDBAdapter().GetUserByID(userID)

	if userErr != nil {
		return nil, nil, userErr
	}

	// Hash token

	dk := hash(refreshToken, user.Salt)

	_, retrieveError := waechter.getDBAdapter().FindRefreshToken(userID, string(dk))

	if retrieveError != nil {
		return nil, nil, retrieveError
	}

	return user, &claims, nil

}

type accessClaims struct {
	jwt.StandardClaims
	Realm string `json:"rlm"`
}

// UserCheckAccessToken make sure the access token is valid
func (waechter *Waechter) UserCheckAccessToken(realm string, token string) (*jwt.StandardClaims, *AuthError) {
	claims := accessClaims{}
	_, parseError := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(waechter.JwtSecret), nil
	})

	if parseError != nil || claims.Realm != realm {
		return nil, &AuthError{
			ErrorCode:   "invalidAccessToken",
			Description: "The access token passed was invalid",
			IsInternal:  false,
		}
	}

	return &claims.StandardClaims, nil

}

// UserGenerateAccessToken issues a new access token based on a refresh token
func (waechter *Waechter) UserGenerateAccessToken(claims *jwt.StandardClaims, realm string, expires time.Duration) (string, *AuthError) {
	result, err := jwt.NewWithClaims(jwt.SigningMethodHS512, accessClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    waechter.JwtIssuer,
			IssuedAt:  time.Now().Unix(),
			Subject:   claims.Subject,
			ExpiresAt: time.Now().Add(expires).Unix(),
		},
		Realm: realm,
	}).SignedString([]byte(waechter.JwtSecret))

	if err != nil {
		return "", CryptError(err)
	}

	return result, nil

}
