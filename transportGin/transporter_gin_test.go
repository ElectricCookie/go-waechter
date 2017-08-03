package transportGin_test

import (
	"encoding/json"
	"net/http"
	"testing"

	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/ElectricCookie/go-waechter/dbMemory"
	"github.com/ElectricCookie/go-waechter/localeDefault"
	"github.com/ElectricCookie/go-waechter/testEmail"
	"github.com/ElectricCookie/go-waechter/transportGin"
	"github.com/braintree/manners"
	"github.com/franela/goreq"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func makeRequest(method string, endpoint string, parameters interface{}, result interface{}, cookies *http.Cookie) *waechter.AuthError {

	req := goreq.Request{
		Uri:    "http://127.0.0.1:3333" + endpoint,
		Method: method,
		Body:   parameters,
	}

	if cookies != nil {
		req = req.WithCookie(cookies)
	}
	res, err := req.Do()

	if err != nil {
		panic(err)
	}

	// Try to get the result
	response := transportGin.JSONResponse{}

	res.Body.FromJsonTo(&response)

	if response.Status {

		// Marshal the result

		b, _ := json.Marshal(response.Data)

		json.Unmarshal(b, result)

		return nil

	} else {
		return response.Err
	}

}

var _ = Describe("User:Register", func() {

	dbAdapter := &dbMemory.MemoryAdapter{}

	var verificationToken string
	var resetToken string

	dbAdapter.Reset()

	emailAdapter := testEmail.NewAdapter()

	translations := localeDefault.NewDefaultTranslations()

	translations.CompanyName = "test-company"
	translations.CompanyWebsite = "test-website.com"
	translations.LogoURL = "https://codyhouse.co/demo/advanced-search-form/img/cd-logo.svg" //Shoutout to codyhouse.co for this awesome placeholder
	translations.UserVerifyEmailAddress = "test-website.com/confirm/"

	w := waechter.New("somesecret", "go-waechter", dbAdapter, emailAdapter, translations)

	gin.DefaultWriter = GinkgoWriter
	ginC := transportGin.GinConnector{
		Waechter:   w,
		AuthPath:   "/auth",
		Domain:     "",
		Debug:      true,
		Writer:     GinkgoWriter,
		ForceHTTPS: false,
	}

	r := gin.Default()

	ginC.DefaultRoutes(r)

	go func() {
		manners.ListenAndServe("127.0.0.1:3333", r)
	}()

	Describe("Register", func() {

		It("Should register a new user", func() {
			res := struct {
				Token string `json:"token"`
			}{}
			authErr := makeRequest("POST", "/auth/register", waechter.UserRegisterParams{
				Username:  "ElectricCookie",
				Password:  "Password123",
				Email:     "test-email@foo.com",
				FirstName: "Sir Electric",
				LastName:  "Cookie",
				Language:  "de",
			}, &res, nil)

			verificationToken = res.Token

			Expect(authErr).To(BeNil())
			user, findErr := dbAdapter.GetUserByUsername("ElectricCookie")
			Expect(user).NotTo(BeNil())
			Expect(findErr).To(BeNil())

		})

	})

	Describe("VerifyEmail", func() {

		It("should verify the email of a user", func() {
			user, _ := dbAdapter.GetUserByUsername("ElectricCookie")
			authErr := makeRequest("POST", "/auth/verifyEmail", waechter.UserVerifyEmailParameters{
				Token:  verificationToken,
				UserID: user.ID,
			}, nil, nil)

			Expect(authErr).To(BeNil())

			user, _ = dbAdapter.GetUserByUsername("ElectricCookie")

			Expect(user.EmailVerfied).To(BeTrue())

		})

	})

	Describe("Login", func() {

		Describe("UserLoginUsername", func() {

		})

		Describe("UserLoginEmail", func() {

		})

		Describe("UserLoginUsernameOrEmail", func() {

		})

		It("should login somebody with the correct credentials in", func() {

			authErr := makeRequest("POST", "/auth/login", waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookie",
				Password:        "Password123",
				RememberMe:      false,
			}, nil, nil)

			Expect(authErr).To(BeNil())

		})

		It("should fail to log in somebody with the wrong credentials", func() {

			authErr := makeRequest("POST", "/auth/login", waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookie",
				Password:        "Password1234",
				RememberMe:      false,
			}, nil, nil)

			Expect(authErr).ToNot(BeNil())

			authErr = makeRequest("POST", "/auth/login", waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookiee",
				Password:        "Password123",
				RememberMe:      false,
			}, nil, nil)

			Expect(authErr).ToNot(BeNil())

		})
	})

	Describe("ForgotPassword", func() {

		It("should generate a token and send an email when called", func() {
			res := struct {
				Token string `json:"token"`
			}{}
			authErr := makeRequest("POST", "/auth/forgotPassword", waechter.ForgotPasswordParams{
				Email: "test-email@foo.com",
			}, &res, nil)

			Expect(authErr).To(BeNil())

			user, _ := dbAdapter.GetUserByEmail("test-email@foo.com")

			resetToken = res.Token

			Expect(user.ForgotPasswordToken).NotTo(Equal("deactivated"))

		})

	})

	Describe("ResetPassword", func() {

		It("Should set the password to a new one", func() {
			user, _ := dbAdapter.GetUserByUsername("ElectricCookie")

			authErr := makeRequest("POST", "/auth/resetPassword", waechter.ResetPasswordParams{
				UserID:      user.ID,
				NewPassword: "newPassword123",
				Token:       resetToken,
			}, nil, nil)

			Expect(authErr).To(BeNil())

			authErr = makeRequest("POST", "/auth/login", waechter.UserLoginEmailOrUsernameData{
				UsernameOrEmail: "ElectricCookie",
				Password:        "newPassword123",
				RememberMe:      false,
			}, nil, nil)

			Expect(authErr).To(BeNil())

		})

	})

	AfterSuite(func() {
		// Close the server to allow the program to shut down properly.
		manners.Close()

	})

})

func TestGinTransporter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoWaechter Suite")
}
