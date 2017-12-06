package waechter

import (
	"time"
)

// performLogin checks if the password passed works for the user
func (waechter *Waechter) performLogin(password string, user *User, rememberMe bool) (string, *AuthError) {

	identifier := user.Username + ":Login"

	if try(identifier) != nil {
		return "", resourceBlockedError()
	}

	defer time.Sleep(time.Second * 3)
	defer release(identifier)

	if dk := hash(password, user.Salt); string(dk) != user.PasswordHash {
		return "", invalidPasswordError()
	}
	if !user.EmailVerfied && waechter.RequireActivation {
		return "", emailNotVerifiedError()
	}

	// Generate a refresh token
	var expires int64

	if rememberMe {

		if waechter.SessionDurationRememberMe == -1 {
			expires = -1
		} else {
			expires = time.Now().Add(waechter.SessionDurationRememberMe).Unix()
		}

	} else {

		expires = time.Now().Add(waechter.SessionDurationDefault).Unix()

	}

	token, tokenError := waechter.UserGenerateRefreshToken(user.ID, expires)

	if tokenError != nil {
		return "", tokenError
	}

	return token, nil

}
