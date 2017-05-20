package waechter

//EmailAdapter is used to send registration emails or request a new password
type EmailAdapter interface {
	SendRegistrationEmail(user *User) error
	SendCustomEmail(receiver *User, subject string, from string, content string) error
}
