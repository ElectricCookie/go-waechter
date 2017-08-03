package waechter

import (
	"time"
)

// LoginEmailOrUsernameData is the required information for logging in
type LoginEmailOrUsernameData struct {
	UsernameOrEmail string `json:"id" binding:"required"`
	Password        string `json:"password" binding:"required"`
	RememberMe      bool   `json:"rememberMe"`
}

// LoginWithUsernameOrEmail logs a user in using email or username and password. Returns a new refresh token. Possible AuthErrors: unknownUser, cryptError, wrongPassword
func (waechter *Waechter) LoginWithUsernameOrEmail(parameters LoginEmailOrUsernameData) (*string, *AuthError) {

	// Retrieve the desired user
	user, err := waechter.getDBAdapter().GetUserByUsernameOrEmail(parameters.UsernameOrEmail)

	if err != nil {
		return nil, &AuthError{
			ErrorCode:   "unknownUser",
			Description: "The user requested was not found",
			IsInternal:  false,
		}
	}

	// Crypt the password using the salt of the user

	dk := hash(parameters.Password, user.Salt)

	if string(dk) != user.PasswordHash {

		return nil, &AuthError{
			ErrorCode:   "wrongPassword",
			Description: "The passsword provided was not correct",
			IsInternal:  false,
		}

	}

	if !user.EmailVerfied {
		return nil, &AuthError{
			ErrorCode:   "emailNotVerified",
			Description: "The email address is not verified",
			IsInternal:  false,
		}
	}

	// Generate a refresh token

	var expires int64

	if parameters.RememberMe {

		if waechter.SessionDurationRememberMe == nil {
			expires = -1
		} else {
			expires = time.Now().Add(*waechter.SessionDurationRememberMe).Unix()
		}

	} else {

		if waechter.SessionDurationDefault == nil {
			expires = time.Now().Add(time.Hour * 2).Unix()
		} else {
			expires = time.Now().Add(*waechter.SessionDurationDefault).Unix()
		}

	}

	token, tokenError := waechter.GenerateRefreshToken(user.ID, expires)

	if tokenError != nil {
		return nil, tokenError
	}

	return token, nil

}
