package waechter

import (
	"time"

	"golang.org/x/crypto/scrypt"
)

// RegisterParams are the parameters used to register a new user.
type RegisterParams struct {
	Username  string `validate:"required"`
	Password  string `validate:"required"`
	Email     string `validate:"required,email"`
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Language  string `validate:"required"`
}

//Register a new user
func (waechter *Waechter) Register(params RegisterParams) *AuthError {

	// Check if user exists

	_, err := waechter.getDBAdapter().GetUserByUsername(params.Username)

	if err == nil {
		return &AuthError{
			ErrorCode:   "usernameUsed",
			Description: "The desired username is already in use",
			IsInternal:  false,
		}
	}

	// Check if email exists

	_, err = waechter.getDBAdapter().GetUserByEmail(params.Email)

	if err == nil {
		return &AuthError{
			ErrorCode:   "emailTaken",
			Description: "The desired email is already in use",
		}
	}

	// Generate salt

	salt, saltGenerationErr := generateRandomString(64)

	if saltGenerationErr != nil {
		return randomError(saltGenerationErr)
	}

	// Generate activation/verfication token

	verificationToken, verificationTokenGenerationErr := generateRandomString(64)

	if verificationTokenGenerationErr != nil {
		return randomError(verificationTokenGenerationErr)
	}

	tokenHash, tokenScryptErr := scrypt.Key([]byte(verificationToken), []byte(salt), 16384, 8, 1, 32)

	if tokenScryptErr != nil {
		return cryptError(tokenScryptErr)
	}

	passwordHash, scryptErr := scrypt.Key([]byte(params.Password), []byte(salt), 16384, 8, 1, 32)

	if scryptErr != nil {
		return cryptError(scryptErr)
	}

	newUser := &User{
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
		return dbWriteErr(saveErr)
	}

	email, err := waechter.getLocales().GetRegistrationEmail(newUser.Language, newUser)

	emailErr := waechter.getEmailAdapter().SendEmail(email)

	if emailErr != nil {
		return emailSendErr(emailErr)
	}

	return nil
}
