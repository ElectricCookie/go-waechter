package waechter_test

import (
	waechter "github.com/ElectricCookie/go-waechter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User:ForgotPassword", func() {

	var w *waechter.Waechter
	var token *string
	var user *waechter.User

	BeforeEach(func() {

		dbAdapter := waechter.NewMongoAdapter("localhost:27017", "waechter-test")

		dbAdapter.Db.DropDatabase()

		emailAdapter := NewTestEmailAdapter()

		translations := &waechter.DefaultTranslations{
			CompanyName:         "test-company",
			CompanyWebsite:      "test-website.com",
			LogoURL:             "https://codyhouse.co/demo/advanced-search-form/img/cd-logo.svg", //Shoutout to codyhouse.co for this awesome placeholder
			DefaultLanguage:     "en",
			Locales:             waechter.GetDefaultLocales(),
			ConfirmEmailAddress: "test-website.com/confirm/",
		}

		w = waechter.New("somesecret", "go-waechter", dbAdapter, emailAdapter, translations)

		w.Register(waechter.RegisterParams{
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

	Context("Email was verified", func() {

		BeforeEach(func() {
			w.VerifyEmailAddress(user.ID, *token)
		})

		It("should send an email to the targeted user", func() {

		})

		It("should fail if the email is unknown", func() {

		})

		It("should fail if the parameter passed is not an email", func() {

		})

		It("should always create a new token in the DB", func() {

		})

	})

	It("should fail if the email is not verified", func() {

		_, err := w.LoginWithUsernameOrEmail(waechter.LoginEmailOrUsernameData{
			UsernameOrEmail: "ElectricCookie",
			Password:        "test123",
			RememberMe:      false,
		})

		Expect(err).ToNot(BeNil())
		Expect(err.ErrorCode).To(Equal("emailNotVerified"))

	})

})
