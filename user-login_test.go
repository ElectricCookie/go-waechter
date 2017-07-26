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

		dbAdapter := &waechter.MemoryAdapter{}

		dbAdapter.Reset()

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

	Context("email is verified", func() {

		BeforeEach(func() {
			w.VerifyEmailAddress(user.ID, *token)
		})

		It("should return a new token if the login succeeds", func() {

			refreshToken, err := w.LoginWithUsernameOrEmail(waechter.LoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookie",
				Password:        "test123",
				RememberMe:      false,
			})

			Expect(err).To(BeNil())

			Expect(*refreshToken).ToNot(BeEmpty())

			_, verifyErr := w.CheckRefreshToken(*refreshToken)

			Expect(verifyErr).To(BeNil())

		})

		It("should fail if the password is wrong", func() {
			_, err := w.LoginWithUsernameOrEmail(waechter.LoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookie",
				Password:        "test1234",
				RememberMe:      false,
			})

			Expect(err).ToNot(BeNil())
			Expect(err.ErrorCode).To(Equal("wrongPassword"))

		})

		It("should fail if the username is not found", func() {
			_, err := w.LoginWithUsernameOrEmail(waechter.LoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookiee",
				Password:        "test123",
				RememberMe:      true,
			})

			Expect(err).ToNot(BeNil())
			Expect(err.ErrorCode).To(Equal("unknownUser"))

		})

	})

	It("should fail if the email address is not verified", func() {

		_, err := w.LoginWithUsernameOrEmail(waechter.LoginEmailOrUsernameData{
			UsernameOrEmail: "ElectricCookie",
			Password:        "test123",
			RememberMe:      false,
		})

		Expect(err).ToNot(BeNil())
		Expect(err.ErrorCode).To(Equal("emailNotVerified"))

	})

})
