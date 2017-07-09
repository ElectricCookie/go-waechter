package waechter

//LocalesAdapter describes all localization options in go-waechter
type LocalesAdapter interface {
	GetRegistrationEmail(language string, user *User, verificationToken string) (*Email, error)
}
