package emailSmtp

import (
	waechter "github.com/ElectricCookie/go-waechter"
	gomail "gopkg.in/gomail.v2"
)

//New creates a new SMTP adapter
func New(url string, port int, user string, password string) *SMTPAdapter {
	return &SMTPAdapter{
		SMTPURL:      url,
		SMTPPort:     port,
		SMTPUser:     user,
		SMTPPassword: password,
	}
}

//SMTPAdapter is the default email Adapter. Using basic SMTP
type SMTPAdapter struct {
	SMTPURL      string
	SMTPPort     int
	SMTPUser     string
	SMTPPassword string
}

//SendEmail is used to implement other messages and notifications to a user.
func (adapter *SMTPAdapter) SendEmail(email *waechter.Email) error {

	message := gomail.NewMessage()
	message.SetHeader("From", email.From)
	message.SetHeader("To", email.To)
	message.SetHeader("Subject", email.Subject)
	message.SetBody("text/html", email.Content)

	d := gomail.NewDialer(
		adapter.SMTPURL,
		adapter.SMTPPort,
		adapter.SMTPUser,
		adapter.SMTPPassword,
	)

	if err := d.DialAndSend(message); err != nil {
		return err
	}

	return nil

}
