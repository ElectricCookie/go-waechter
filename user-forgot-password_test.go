package waechter_test

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/ElectricCookie/go-waechter/localeDefault"
	"github.com/ElectricCookie/go-waechter/testEmail"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User:ForgotPassword", func() {

	var w *waechter.Waechter
	var token *string
	var user *waechter.User

	BeforeEach(func() {

		dbAdapter := &waechter.MemoryAdapter{}

		dbAdapter.Reset()

		emailAdapter := testEmail.NewAdapter()

		translations := localeDefault.NewDefaultTranslations()

		translations.CompanyName = "test-company"
		translations.CompanyWebsite = "test-website.com"
		translations.LogoURL = "https://codyhouse.co/demo/advanced-search-form/img/cd-logo.svg" //Shoutout to codyhouse.co for this awesome placeholder
		translations.VerifyEmailAddress = "test-website.com/confirm/"

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
			w.VerifyEmailAddress(waechter.VerifyEmailParameters{
				UserID: user.ID,
				Token:  *token,
			})
		})

		It("should fail if the email is unknown", func() {

			_, err := w.ForgotPassword(waechter.ForgotPasswordParams{Email: "unknownEmail@something.com"})

			Expect(err).ToNot(BeNil())
			Expect(err.ErrorCode).To(Equal("userNotFound"))

		})

		It("should fail if the parameter passed is not an email", func() {

			_, err := w.ForgotPassword(waechter.ForgotPasswordParams{Email: "123coaks"})
			Expect(err).ToNot(BeNil())
			Expect(err.ErrorCode).To(Equal("invalidParameters"))

		})

		It("should always create a new token in the DB", func() {

			w.ForgotPassword(waechter.ForgotPasswordParams{Email: "somebody@something.com"})

			userA, _ := w.DbAdapter.GetUserByUsername("ElectricCookie")

			token := userA.ForgotPasswordToken

			Expect(user.ForgotPasswordToken).NotTo(Equal("deactivated"))

			w.ForgotPassword(waechter.ForgotPasswordParams{Email: "somebody@something.com"})

			userB, _ := w.DbAdapter.GetUserByUsername("ElectricCookie")

			Expect(token).ToNot(Equal(userB.ForgotPasswordToken))

		})

	})

	It("should fail if the email is not verified", func() {

		_, err := w.ForgotPassword(waechter.ForgotPasswordParams{Email: "somebody@something.com"})

		Expect(err).ToNot(BeNil())
		Expect(err.ErrorCode).To(Equal("emailNotVerified"))

	})

})
