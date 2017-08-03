package testEmail

import (
	"html/template"
	"log"
	"os"

	waechter "github.com/ElectricCookie/go-waechter"
)

//Adapter is used in tests. It writes emails to files instead of sending them.
type Adapter struct {
	template *template.Template
}

//NewAdapter creates a new test-email adapter
func NewAdapter() *Adapter {
	b, _ := Asset("testEmail/test-email-template.html")
	t := template.Must(template.New("test-email-template").Parse(string(b)))

	return &Adapter{
		template: t,
	}
}

//SendEmail executes the test template and saves the test result
func (adapter *Adapter) SendEmail(email *waechter.Email) error {

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
