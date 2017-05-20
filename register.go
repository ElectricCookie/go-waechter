package waechter

import (
	"time"

	"golang.org/x/crypto/scrypt"
)

// RegisterParams are the parameters used to register a new user.
type RegisterParams struct {
	Username   string `validate:"required"`
	Password   string `validate:"required"`
	Email      string `validate:"required,email"`
	FirstName  string `validate:"required"`
	LastName   string `validate:"required"`
	Language   string `validate:"required"`
	Registered *time.Time
	LastLogin  *time.Time
}

//Register a new user
func Register(params RegisterParams) *AuthError {

	// Check if user exists

	_, err := getAdapter().GetUserByUsername(params.Username)

	if err == nil {
		return &AuthError{
			ErrorCode:   "usernameUsed",
			Description: "The desired username is already in use",
			IsInternal:  false,
		}
	}

	// Check if email exists

	_, err = getAdapter().GetUserByEmail(params.Email)

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
		VerificationToken: verificationToken,
		LastLogin:         time.Now(),
		Registered:        time.Now(),
	}

	saveErr := getAdapter().CreateUser(newUser)

	if saveErr != nil {
		return dbWriteErr(saveErr)
	}

	emailErr := getEmailAdapter().SendRegistrationEmail(newUser)

	if emailErr != nil {
		return emailSendErr(emailErr)
	}

	return nil
}
