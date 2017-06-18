package waechter

//EmailAdapter is used to send registration emails or request a new password
type EmailAdapter interface {
	SendEmail(email *Email) error
}

type Email struct {
	From    string
	To      string
	Subject string
	Content string
}
