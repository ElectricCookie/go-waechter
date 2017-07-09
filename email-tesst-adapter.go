package waechter

import (
	"log"
	"os"
	"text/template"
)

//TestEmailAdapter is used in tests. It writes emails to files instead of sending them.
type TestEmailAdapter struct {
	template *template.Template
}

//NewTestEmailAdapter creates a new test-email adapter
func NewTestEmailAdapter() *TestEmailAdapter {

	t := template.Must(template.ParseFiles("./templates/test-email-template.html"))

	return &TestEmailAdapter{
		template: t,
	}
}

//SendEmail executes the test template and saves the test result
func (adapter *TestEmailAdapter) SendEmail(email *Email) error {

	os.MkdirAll("./test_results/emails", os.ModePerm)

	f, err := os.Create("./test_results/emails/" + email.From + "-" + email.Subject + ".html")
	if err != nil {
		log.Println("create file: ", err)
		return err
	}

	err = adapter.template.Execute(f, email)
	if err != nil {
		log.Print("execute: ", err)
		return err
	}

	return nil

}
