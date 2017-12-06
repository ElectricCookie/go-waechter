package waechter

//LocalesAdapter describes all localization options in go-waechter
type LocalesAdapter interface {
	GetRegistrationEmail(user *User, verificationToken string) (*Email, error)
	GetForgotPasswordEmail(user *User, forgotPasswordToken string) (*Email, error)
	GetPasswordResetEmail(user *User) (*Email, error)
	GetLanguages() []string
	GetDefaultLanguage() string
}
