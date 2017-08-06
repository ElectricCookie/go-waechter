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

	Context("email is verified", func() {

		BeforeEach(func() {

			w.UserVerifyEmailAddress(waechter.UserVerifyEmailParameters{
				UserID: user.ID,
				Token:  *token,
			})
		})

		It("should return a new token if the login succeeds", func() {

			refreshToken, err := w.UserLoginWithUsernameOrEmail(waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookie",
				Password:        "test123",
				RememberMe:      false,
			})

			Expect(err).To(BeNil())

			Expect(*refreshToken).ToNot(BeEmpty())

			_, _, verifyErr := w.CheckRefreshToken(*refreshToken)

			Expect(verifyErr).To(BeNil())

		})

		It("should fail if the password is wrong", func() {
			_, err := w.UserLoginWithUsernameOrEmail(waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookie",
				Password:        "test1234",
				RememberMe:      false,
			})

			Expect(err).ToNot(BeNil())
			Expect(err.ErrorCode).To(Equal("wrongPassword"))

		})

		It("should fail if the username is not found", func() {
			_, err := w.UserLoginWithUsernameOrEmail(waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookiee",
				Password:        "test123",
				RememberMe:      true,
			})

			Expect(err).ToNot(BeNil())
			Expect(err.ErrorCode).To(Equal("userNotFound"))

		})

	})

	It("should fail if the email address is not verified", func() {

		_, err := w.UserLoginWithUsernameOrEmail(waechter.UserLoginEmailOrUsernameData{
			UsernameOrEmail: "ElectricCookie",
			Password:        "test123",
			RememberMe:      false,
		})

		Expect(err).ToNot(BeNil())
		Expect(err.ErrorCode).To(Equal("emailNotVerified"))

	})

})
