package waechter

import (
	"gopkg.in/asaskevich/govalidator.v4"
)

//ForgotPassword sends an email to recover the password
func (w *Waechter) ForgotPassword(emailAddress string) (*string, *AuthError) {

	if !govalidator.IsEmail(emailAddress) {
		return nil, invalidParameters(nil)
	}

	user, err := w.DbAdapter.GetUserByEmail(emailAddress)

	if err != nil {
		return nil, userNotFoundError(err)
	}

	if !user.EmailVerfied {

		return nil, &AuthError{
			ErrorCode:   "emailNotVerified",
			Description: "The email address is not verified",
			IsInternal:  false,
		}

	}

	token := generateRandomString(64)

	// Hash dat

	hash := hash(token, user.Salt)

	// Send email

	email, err := w.Locales.GetForgotPasswordEmail(user, token)

	if err != nil {
		return nil, internalError(err)
	}

	w.EmailAdapter.SendEmail(email)

	err = w.DbAdapter.SetForgotPasswordToken(user.ID, hash)

	if err != nil {
		return nil, dbWriteErr(err)
	}

	return &token, nil
}
