package main

import (
	"time"

	"golang.org/x/crypto/scrypt"
)

// LoginEmailOrUsernameData is the required information for logging in
type LoginEmailOrUsernameData struct {
	UsernameOrEmail string `json:"id" binding:"required"`
	Password        string `json:"password" binding:"required"`
	RememberMe      bool   `json:"rememberMe" binding:"required"`
}

// LoginWithUsernameOrEmail logs a user in using email or username and password. Possible AuthErrors: unknownUser, cryptError, wrongPassword
func LoginWithUsernameOrEmail(parameters LoginEmailOrUsernameData) (*string, *AuthError) {

	// Retrieve the desired user
	user, err := getAdapter().GetUserByUsernameOrEmail(parameters.UsernameOrEmail)

	if err != nil {
		return nil, &AuthError{
			ErrorCode:   "unknownUser",
			Description: "The user requested was not found",
			IsInternal:  false,
		}
	}

	// Crypt the password using the salt of the user
	dk, cryptError := scrypt.Key([]byte(parameters.Password), []byte(user.Salt), 16384, 8, 1, 32)

	if cryptError != nil {
		return nil, &AuthError{
			ErrorCode:   "cryptError",
			Description: "Something went wrong while encrypting the password",
			IsInternal:  true,
		}
	}

	if string(dk) != user.PasswordHash {

		return nil, &AuthError{
			ErrorCode:   "wrongPassword",
			Description: "The passsword provided was not correct",
			IsInternal:  false,
		}

	}

	// Generate a refresh token

	var expires int64 = -1

	if parameters.RememberMe {
		expires = time.Now().Unix() + getParameters().sessionDuration
	}

	token, tokenError := GenerateRefreshToken(user.ID, expires)

	if tokenError != nil {
		return nil, tokenError
	}

	return token, nil

}
