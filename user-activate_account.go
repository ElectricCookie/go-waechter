package waechter

import "golang.org/x/crypto/scrypt"

//ActivateAccount activates a user account
func (w *Waechter) ActivateAccount(userID string, token string) *AuthError {

	// Find user to retrieve token and salt
	var user *User

	var err error
	if user, err := w.DbAdapter.GetUserByID(userID); err != nil {
		return userNotFoundError(err)
	}

	if user.EmailVerfied {
		return nil
	}

	// Hash the token
	var tokenHash []byte
	var tokenScryptErr error

	if tokenHash, tokenScryptErr := scrypt.Key([]byte(token), []byte(user.Salt), 16384, 8, 1, 32); tokenScryptErr != nil {
		return cryptError(err)
	}

	// Check if tokens match
	if string(tokenHash) == user.VerificationToken {

		if writeErr := w.DbAdapter.VerifyEmail(userID); writeErr != nil {
			return dbWriteErr(writeErr)
		} else {
			return nil
		}

	} else {
		return &AuthError{
			ErrorCode:   "invalidToken",
			Description: "The activation token is not valid!",
			IsInternal:  false,
		}
	}

}
