package waechter

//VerifyEmailParameters describes the paramteres passed to VerifyEmailAddress
type VerifyEmailParameters struct {
	UserID string `json:"userId" binding:"required"`
	Token  string `json:"token" binding:"required"`
}

//VerifyEmailAddress verifies the email address
func (w *Waechter) VerifyEmailAddress(parameters VerifyEmailParameters) *AuthError {

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
