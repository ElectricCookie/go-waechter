package waechter

import govalidator "gopkg.in/asaskevich/govalidator.v4"

//SendVerificationEmail sends an email to the user with a link to verify their email address
func (waechter *Waechter) SendVerificationEmail(emailAddress string) (*string, *AuthError) {

	if !govalidator.IsEmail(emailAddress) {
		return nil, InvalidParametersError(nil)
	}

	user, err := waechter.DbAdapter.GetUserByEmail(emailAddress)

	if err != nil {
		return nil, userNotFoundError(err)
	}

	if user.EmailVerfied {

		return nil, &AuthError{
			ErrorCode:   "alreadyVerified",
			Description: "The email address is already verified",
			IsInternal:  false,
		}

	}

	verificationToken := generateRandomString(32)

	tokenHash := hash(verificationToken, user.Salt)

	waechter.getDBAdapter().SetVerificationToken(user.ID, string(tokenHash))

	email, err := waechter.getLocales().GetRegistrationEmail(user, verificationToken)

	emailErr := waechter.getEmailAdapter().SendEmail(email)

	if emailErr != nil {
		return nil, emailSendError(emailErr)
	}

	return &verificationToken, nil

}
