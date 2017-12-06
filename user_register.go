package waechter

import (
	"time"

	validator "github.com/asaskevich/govalidator"
)

//UserRegisterParams are the parameters used to register a new user.
type UserRegisterParams struct {
	Username  string `valid:"required" json:"username" binding:"required"`
	Password  string `valid:"required" json:"password" binding:"required"`
	Email     string `valid:"required,email" json:"email" binding:"required"`
	FirstName string `valid:"required" json:"firstName" binding:"required"`
	LastName  string `valid:"required" json:"lastName" binding:"required"`
	Language  string `valid:"required" json:"language" binding:"required"`
}

func (waechter *Waechter) languageExists(lang string) bool {

	languages := waechter.Locales.GetLanguages()

	for _, langI := range languages {
		if langI == lang {
			return true
		}
	}
	return false
}

//UserRegister a new user
func (waechter *Waechter) UserRegister(params UserRegisterParams) *AuthError {
	// Check if user exists
	valid, validationErrs := validator.ValidateStruct(params)

	if !valid {
		return InvalidParametersError(validationErrs)
	}

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
			ErrorCode:   "emailUsed",
			Description: "The desired email is already in use",
		}
	}

	// Check if the langauge is valid

	if !waechter.languageExists(params.Language) {
		params.Language = waechter.Locales.GetDefaultLanguage()
	}

	// Generate salt

	salt := generateRandomString(32)

	// Generate activation/verification token

	// TODO Validate language

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
		VerificationToken: "deactivated",
		LastLogin:         time.Now(),
		Registered:        time.Now(),
	}

	saveErr := waechter.getDBAdapter().CreateUser(newUser)

	if saveErr != nil {
		return dbWriteError(saveErr)
	}

	return nil
}
