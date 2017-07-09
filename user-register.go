package waechter

import (
	"time"

	validator "gopkg.in/asaskevich/govalidator.v4"
)

// RegisterParams are the parameters used to register a new user.
type RegisterParams struct {
	Username  string `valid:"required"`
	Password  string `valid:"required"`
	Email     string `valid:"required,email"`
	FirstName string `valid:"required"`
	LastName  string `valid:"required"`
	Language  string `valid:"required"`
}

//Register a new user
func (waechter *Waechter) Register(params RegisterParams) (*string, *AuthError) {
	// Check if user exists
	valid, validationErrs := validator.ValidateStruct(params)

	if !valid {
		return nil, invalidParameters(validationErrs)
	}

	_, err := waechter.getDBAdapter().GetUserByUsername(params.Username)

	if err == nil {
		return nil, &AuthError{
			ErrorCode:   "usernameUsed",
			Description: "The desired username is already in use",
			IsInternal:  false,
		}
	}

	// Check if email exists

	_, err = waechter.getDBAdapter().GetUserByEmail(params.Email)

	if err == nil {
		return nil, &AuthError{
			ErrorCode:   "emailUsed",
			Description: "The desired email is already in use",
		}
	}

	// Generate salt

	salt := generateRandomString(32)

	// Generate activation/verfication token

	verificationToken := generateRandomString(32)

	tokenHash := hash(verificationToken, salt)

	passwordHash := hash(params.Password, salt)

	newUser := &User{
		ID:                generateRandomString(32),
		Username:          params.Username,
		Email:             params.Email,
		FirstName:         params.FirstName,
		LastName:          params.LastName,
		PasswordHash:      string(passwordHash),
		Salt:              salt,
		Language:          params.Language,
		EmailVerfied:      false,
		VerificationToken: string(tokenHash),
		LastLogin:         time.Now(),
		Registered:        time.Now(),
	}

	saveErr := waechter.getDBAdapter().CreateUser(newUser)

	if saveErr != nil {
		return nil, dbWriteErr(saveErr)
	}

	email, err := waechter.getLocales().GetRegistrationEmail(newUser.Language, newUser, verificationToken)

	emailErr := waechter.getEmailAdapter().SendEmail(email)

	if emailErr != nil {
		return nil, emailSendErr(emailErr)
	}

	return &verificationToken, nil
}
