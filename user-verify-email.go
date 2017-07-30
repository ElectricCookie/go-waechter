package waechter

//VerifyEmailAddress verifies the email address
func (w *Waechter) VerifyEmailAddress(userID string, token string) *AuthError {

	// Find user to retrieve token and salt
	var user *User

	var err error
	if user, err = w.DbAdapter.GetUserByID(userID); err != nil || user == nil {
		return userNotFoundError(err)
	}

	if user.EmailVerfied {
		return nil
	}

	// Hash the token

	tokenHash := hash(token, user.Salt)

	// Check if tokens match
	if tokenHash == user.VerificationToken {
		if writeErr := w.DbAdapter.VerifyEmail(userID); writeErr != nil {
			return dbWriteErr(writeErr)
		}

	} else {
		return &AuthError{
			ErrorCode:   "invalidVerificationToken",
			Description: "The verification token is not valid!",
			IsInternal:  false,
		}
	}

	return nil
}
