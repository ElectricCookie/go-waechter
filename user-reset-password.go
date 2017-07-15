package waechter

//ResetPassword changes the password using  a reset token. It also sends an email to notify the user of the changes made to the account
func (waechter *Waechter) ResetPassword(userID string, token string, newPassword string) *AuthError {

	// Get the user
	user, userErr := waechter.getDBAdapter().GetUserByID(userID)

	if userErr != nil || user == nil {
		return userNotFoundError(userErr)
	}

	// Hash the token

	tokenHash := hash(token, user.Salt)

	if user.ForgotPasswordToken != tokenHash || user.ForgotPasswordToken == "dectivated" {

		return &AuthError{
			ErrorCode:   "invalidToken",
			Description: "The token passed is invalid",
			IsInternal:  false,
		}

	}

	// Set the new password

	newPasswordHash := hash(newPassword, user.Salt)

	if user.PasswordHash == newPasswordHash {
		return &AuthError{
			ErrorCode:   "passwordsMatch",
			Description: "The old and new password matched",
			IsInternal:  false,
		}
	}

	waechter.getDBAdapter().SetForgotPasswordToken(user.ID, "deactivated")
	waechter.getDBAdapter().SetPassword(user.ID, newPasswordHash)

	// Send email

	email, emailErr := waechter.getLocales().GetPasswordResetEmail(user)

	if emailErr != nil {
		return internalError(emailErr)
	}

	waechter.getEmailAdapter().SendEmail(email)

	return nil

}
