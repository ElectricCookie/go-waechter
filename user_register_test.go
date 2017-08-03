package waechter_test

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/ElectricCookie/go-waechter/dbMemory"
	"github.com/ElectricCookie/go-waechter/localeDefault"
	"github.com/ElectricCookie/go-waechter/testEmail"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User:Register", func() {

	var w *waechter.Waechter

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
	})

	Describe("register", func() {

		Context("Proper registration", func() {
			BeforeEach(func() {
				w.UserRegister(waechter.UserRegisterParams{
					Username:  "ElectricCookie",
					Email:     "somebody@something.com",
					Password:  "test123",
					Language:  "de",
					FirstName: "Electric",
					LastName:  "Cookie",
				})
			})

			It("Should create a record that can be retrieved", func() {
				user, err := w.DbAdapter.GetUserByUsername("ElectricCookie")
				Expect(user).ToNot(BeNil())
				Expect(err).To(BeNil())
			})

			It("Should create a non verified account", func() {
				user, err := w.DbAdapter.GetUserByUsername("ElectricCookie")
				Expect(user.EmailVerfied).To(BeFalse())
				Expect(err).To(BeNil())
			})

		})

		It("should fail if the username is taken", func() {

			errOne := w.UserRegister(waechter.UserRegisterParams{
				Username:  "ElectricCookie",
				Email:     "somebody@something.com",
				Password:  "test123",
				Language:  "de",
				FirstName: "Electric",
				LastName:  "Cookie",
			})

			errTwo := w.UserRegister(waechter.UserRegisterParams{
				Username:  "ElectricCookie",
				Email:     "somebody@something123.com",
				Password:  "test123",
				Language:  "de",
				FirstName: "Electric",
				LastName:  "Cookie",
			})

			Expect(errOne).To(BeNil())
			Expect(errTwo.ErrorCode).To(Equal("usernameUsed"))
			u, _ := w.DbAdapter.GetUserByUsername("ElectricCookie")
			Expect(u.Email).ToNot(Equal("somebody@something123.com"))

		})

		It("should fail if the email is taken", func() {

			errOne := w.UserRegister(waechter.UserRegisterParams{
				Username:  "ElectricCookie",
				Email:     "somebody@something.com",
				Password:  "test123",
				Language:  "de",
				FirstName: "Electric",
				LastName:  "Cookie",
			})

			errTwo := w.UserRegister(waechter.UserRegisterParams{
				Username:  "ElectricCookie2",
				Email:     "somebody@something.com",
				Password:  "test123",
				Language:  "de",
				FirstName: "Electric",
				LastName:  "Cookie",
			})

			Expect(errOne).To(BeNil())
			Expect(errTwo.ErrorCode).To(Equal("emailUsed"))
			u, _ := w.DbAdapter.GetUserByUsername("ElectricCookie")
			Expect(u.Username).ToNot(Equal("ElectricCookie2"))

		})

		Context("check parameters of registration", func() {

			It("should fail if something is missing", func() {
				err := w.UserRegister(waechter.UserRegisterParams{
					// Username is missing here
					Email:     "test@something.com",
					Password:  "test123",
					Language:  "de",
					FirstName: "Electric",
					LastName:  "Cookie",
				})
				Expect(err).ToNot(BeNil())
				Expect(err.ErrorCode).To(Equal("invalidParameters"))
			})

			It("should fail if the email is invalid", func() {
				err := w.UserRegister(waechter.UserRegisterParams{
					Username:  "ElectricCookie",
					Email:     "foo", // Invalid email
					Password:  "test123",
					Language:  "de",
					FirstName: "Electric",
					LastName:  "Cookie",
				})

				Expect(err).ToNot(BeNil())
				Expect(err.ErrorCode).To(Equal("invalidParameters"))
			})

		})

	})

})
