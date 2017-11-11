package waechter

//UserResetPasswordParams describes the data passed to ResetPassword
type UserResetPasswordParams struct {
	UserID      string `json:"userId" bind:"required"`
	Token       string `json:"token" bind:"required"`
	NewPassword string `json:"newPassword" bind:"required"`
}

//UserResetPassword changes the password using  a reset token. It also sends an email to notify the user of the changes made to the account
func (waechter *Waechter) UserResetPassword(params UserResetPasswordParams) *AuthError {

	// Get the user
	user, userErr := waechter.getDBAdapter().GetUserByID(params.UserID)

	if userErr != nil || user == nil {
		return userNotFoundError(userErr)
	}

	// Hash the token

	tokenHash := hash(params.Token, user.Salt)

	if user.ForgotPasswordToken != tokenHash || user.ForgotPasswordToken == "dectivated" {

		return &AuthError{
			ErrorCode:   "invalidToken",
			Description: "The token passed is invalid",
			IsInternal:  false,
		}

	}

	// Set the new password

	newPasswordHash := hash(params.NewPassword, user.Salt)

	if user.PasswordHash == newPasswordHash {
		return &AuthError{
			ErrorCode:   "passwordsMatch",
			Description: "The old and new password matched",
			IsInternal:  false,
		}
	}

	dbErr := waechter.getDBAdapter().SetForgotPasswordToken(user.ID, "deactivated")

	if dbErr != nil {
		return internalError(dbErr)
	}

	dbErr = waechter.getDBAdapter().SetPassword(user.ID, newPasswordHash)

	if dbErr != nil {
		return internalError(dbErr)
	}

	// Send email

	email, emailErr := waechter.getLocales().GetPasswordResetEmail(user)

	if emailErr != nil {
		return internalError(emailErr)
	}

	waechter.getEmailAdapter().SendEmail(email)

	return nil

}
