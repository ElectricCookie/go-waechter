package waechter

import govalidator "github.com/asaskevich/govalidator"

//UserSendVerficationParameters describes parameters  that are required to request a verification
type UserSendVerficationParameters struct {
	Email string `json:"email" validate:"email" binding:"required"`
}

//UserSendVerificationEmail sends an email to the user with a link to verify their email address
func (waechter *Waechter) UserSendVerificationEmail(parameters UserSendVerficationParameters) (string, *AuthError) {

	if v, vErr := govalidator.ValidateStruct(parameters); !v {
		return "", InvalidParametersError(vErr)
	}

	user, err := waechter.DbAdapter.GetUserByEmail(parameters.Email)

	if err != nil {
		return "", userNotFoundError(err)
	}

	if user.EmailVerfied {

		return "", &AuthError{
			ErrorCode:   "alreadyVerified",
			Description: "The email address is already verified",
			IsInternal:  false,
		}

	}

	verificationToken := generateRandomString(32)

	tokenHash := hash(verificationToken, user.Salt)

	waechter.getDBAdapter().SetVerificationToken(user.ID, string(tokenHash))

	email, err := waechter.getLocales().GetRegistrationEmail(user, verificationToken)

	if err != nil {
		return "", internalError(err)
	}

	emailErr := waechter.getEmailAdapter().SendEmail(email)

	if emailErr != nil {
		return "", emailSendError(emailErr)
	}

	return verificationToken, nil

}
