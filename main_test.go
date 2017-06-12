package waechter

import (
	"log"
	"testing"

	"github.com/fatih/color"
)

func TestRegister(t *testing.T) {

	// Create mongo db connection

	dbAdapter := NewMongoAdapter("localhost:27017", "waechter-test")

	color.Green("Setting up test database waechter-test")

	dbAdapter.Db.DropDatabase()

	emailAdapter := NewTestEmailAdapter()

	w := New("somesecret", "go-waechter", dbAdapter, emailAdapter)

	err := w.Register(RegisterParams{
		Username:  "ElectricCookie",
		Email:     "yannick.stolle@gmail.com",
		Password:  "test123",
		Language:  "de",
		FirstName: "Electric",
		LastName:  "Cookie",
	})

	if err != nil {
		log.Fatal(err)
	}

}
