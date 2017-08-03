package waechter

import (
	"time"
)

//UserLoginEmailData is the required information for loggin in using the email address
type UserLoginEmailData struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" binding:"required"`
}

//UserLoginEmail logs in using email
func (waechter *Waechter) UserLoginEmail(parameters UserLoginEmailData) (*string, *AuthError) {

	var user *User
	var err error

	// Retrieve the desired user
	if user, err = waechter.DbAdapter.GetUserByEmail(parameters.Email); err != nil {
		return nil, userNotFoundError(err)
	}
	return waechter.performLogin(parameters.Password, user, parameters.RememberMe)

}

//UserLoginUsernameData is the required infromation for logging in using the username
type UserLoginUsernameData struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RememberMe bool   `json:"rememberMe" binding:"required"`
}

//UserLoginUsername logs in using email
func (waechter *Waechter) UserLoginUsername(parameters UserLoginUsernameData) (*string, *AuthError) {

	var user *User
	var err error
	// Retrieve the desired user
	if user, err = waechter.DbAdapter.GetUserByUsername(parameters.Username); err != nil {
		return nil, userNotFoundError(err)
	}

	return waechter.performLogin(parameters.Password, user, parameters.RememberMe)
}

// UserLoginEmailOrUsernameData is the required information for logging in
type UserLoginEmailOrUsernameData struct {
	UsernameOrEmail string `json:"id" binding:"required"`
	Password        string `json:"password" binding:"required"`
	RememberMe      bool   `json:"rememberMe"`
}

// UserLoginWithUsernameOrEmail logs a user in using email or username and password. Returns a new refresh token. Possible AuthErrors: unknownUser, cryptError, wrongPassword
func (waechter *Waechter) UserLoginWithUsernameOrEmail(parameters UserLoginEmailOrUsernameData) (*string, *AuthError) {
	var user *User
	var err error
	// Retrieve the desired user
	if user, err = waechter.getDBAdapter().GetUserByUsernameOrEmail(parameters.UsernameOrEmail); err != nil {
		return nil, userNotFoundError(err)
	}

	return waechter.performLogin(parameters.Password, user, parameters.RememberMe)
}

func (waechter *Waechter) performLogin(password string, user *User, rememberMe bool) (*string, *AuthError) {

	if dk := hash(password, user.Salt); string(dk) != user.PasswordHash {
		return nil, invalidPasswordError()
	}
	if !user.EmailVerfied {
		return nil, emailNotVerifiedError()
	}

	// Generate a refresh token
	var expires int64

	if rememberMe {

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
