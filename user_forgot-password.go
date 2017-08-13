package waechter

import (
	validator "github.com/asaskevich/govalidator"
)

//ForgotPasswordParams describes parameters passed to ForgotPassword
type ForgotPasswordParams struct {
	Email string `json:"email" binding:"required" valid:"required,email"`
}

//ForgotPassword sends an email to recover the password
func (w *Waechter) ForgotPassword(parameters ForgotPasswordParams) (string, *AuthError) {

	valid, validationErrs := validator.ValidateStruct(parameters)

	if !valid {
		return "", InvalidParametersError(validationErrs)
	}

	user, err := w.DbAdapter.GetUserByEmail(parameters.Email)

	if err != nil || user == nil {
		return "", userNotFoundError(err)
	}

	if !user.EmailVerfied {

		return "", &AuthError{
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
		return "", internalError(err)
	}

	w.EmailAdapter.SendEmail(email)

	err = w.DbAdapter.SetForgotPasswordToken(user.ID, hash)

	if err != nil {
		return "", dbWriteError(err)
	}

	return token, nil
}
