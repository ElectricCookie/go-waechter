package waechter

//UserVerifyEmailParams describes the paramteres passed to UserVerifyEmailAddress
type UserVerifyEmailParams struct {
	UserID string `json:"userId" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

//UserVerifyEmailAddress verifies the email address
func (w *Waechter) UserVerifyEmailAddress(parameters UserVerifyEmailParams) *AuthError {

	// Find user to retrieve token and salt
	var user *User

	var err error
	if user, err = w.DbAdapter.GetUserByID(parameters.UserID); err != nil || user == nil {
		return userNotFoundError(err)
	}

	if user.EmailVerfied {
		return nil
	}

	// Hash the token

	tokenHash := hash(parameters.Token, user.Salt)

	// Check if tokens match
	if tokenHash == user.VerificationToken {
		if writeErr := w.DbAdapter.VerifyEmail(parameters.UserID); writeErr != nil {
			return dbWriteError(writeErr)
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
