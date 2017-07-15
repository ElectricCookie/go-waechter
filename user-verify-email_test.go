package waechter_test

import (
	waechter "github.com/ElectricCookie/go-waechter"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User:VerifyEmailAddress", func() {

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

	It("should verify the email address if the token is correct", func() {

		err := w.VerifyEmailAddress(user.ID, *token)

		Expect(err).To(BeNil())

		user, _ = w.DbAdapter.GetUserByUsername("ElectricCookie")

		Expect(user.EmailVerfied).To(BeTrue())

	})

	It("should not verify the email address if the token is incorrect", func() {
		err := w.VerifyEmailAddress(user.ID, "Invalid token")

		Expect(err).ToNot(BeNil())
		Expect(err.ErrorCode).To(Equal("invalidToken"))

		user, _ = w.DbAdapter.GetUserByUsername("ElectricCookie")

		Expect(user.EmailVerfied).ToNot(BeTrue())
	})

})
