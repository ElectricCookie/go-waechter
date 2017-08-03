package waechter_test

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/ElectricCookie/go-waechter/dbMemory"
	"github.com/ElectricCookie/go-waechter/localeDefault"
	"github.com/ElectricCookie/go-waechter/testEmail"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User:UserVerifyEmailAddress", func() {

	var w *waechter.Waechter
	var token *string
	var user *waechter.User

	BeforeEach(func() {

		dbAdapter := &dbMemory.MemoryAdapter{}

		dbAdapter.Reset()

		emailAdapter := testEmail.NewAdapter()

		translations := localeDefault.NewDefaultTranslations()

		translations.CompanyName = "test-company"
		translations.CompanyWebsite = "test-website.com"
		translations.LogoURL = "https://codyhouse.co/demo/advanced-search-form/img/cd-logo.svg" //Shoutout to codyhouse.co for this awesome placeholder
		translations.UserVerifyEmailAddress = "test-website.com/confirm/"

		w = waechter.New("somesecret", "go-waechter", dbAdapter, emailAdapter, translations)

		w.UserRegister(waechter.UserRegisterParams{
			Username:  "ElectricCookie",
			Email:     "somebody@something.com",
			Password:  "test123",
			Language:  "de",
			FirstName: "Electric",
			LastName:  "Cookie",
		})

		user, _ = w.DbAdapter.GetUserByUsername("ElectricCookie")

		token, _ = w.SendVerificationEmail(user.Email)

	})

	It("should verify the email address if the token is correct", func() {

		err := w.UserVerifyEmailAddress(waechter.UserVerifyEmailParameters{
			UserID: user.ID,
			Token:  *token,
		})

		Expect(err).To(BeNil())

		user, _ = w.DbAdapter.GetUserByUsername("ElectricCookie")

		Expect(user.EmailVerfied).To(BeTrue())

	})

	It("should not verify the email address if the token is incorrect", func() {

		err := w.UserVerifyEmailAddress(waechter.UserVerifyEmailParameters{
			UserID: user.ID,
			Token:  "Invalid token",
		})

		Expect(err).ToNot(BeNil())
		Expect(err.ErrorCode).To(Equal("invalidVerificationToken"))

		user, _ = w.DbAdapter.GetUserByUsername("ElectricCookie")

		Expect(user.EmailVerfied).ToNot(BeTrue())
	})

})
