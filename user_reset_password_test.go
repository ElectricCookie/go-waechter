package waechter_test

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/ElectricCookie/go-waechter/dbMemory"
	"github.com/ElectricCookie/go-waechter/localeDefault"
	"github.com/ElectricCookie/go-waechter/testEmail"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User:ResetPassword", func() {

	var w *waechter.Waechter
	var token string
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

		token, _ = w.UserSendVerificationEmail(waechter.UserSendVerficationParameters{
			Email: user.Email,
		})

		w.UserVerifyEmailAddress(waechter.UserVerifyEmailParameters{UserID: user.ID, Token: token})

	})

	Context("Forgot password was called", func() {

		BeforeEach(func() {
			token, _ = w.ForgotPassword(waechter.ForgotPasswordParams{Email: user.Email})
		})

		It("should reset the password if all paremeters are correct", func() {

			err := w.ResetPassword(waechter.ResetPasswordParams{
				UserID:      user.ID,
				Token:       token,
				NewPassword: "newPassword",
			})

			Expect(err).To(BeNil())

			// Do a login

			_, errLogin := w.UserLoginWithUsernameOrEmail(waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: user.Username,
				Password:        "newPassword",
				RememberMe:      false,
			})

			Expect(errLogin).To(BeNil())

		})

	})

})
