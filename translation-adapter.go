package waechter

type TranslationAdapter interface {
	GetRegistrationEmail(language string, user *User) (*Email, error)
}
