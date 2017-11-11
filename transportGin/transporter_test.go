package transportGin_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/ElectricCookie/go-waechter/dbMemory"
	"github.com/ElectricCookie/go-waechter/localeDefault"
	"github.com/ElectricCookie/go-waechter/testEmail"
	"github.com/ElectricCookie/go-waechter/transportGin"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var client = http.Client{}

func makeRequest(method string, endpoint string, parameters interface{}, result interface{}, cookies *http.Cookie) *waechter.AuthError {

	body, err := json.Marshal(parameters)

	req, err := http.NewRequest(method, endpoint, bytes.NewReader(body))
	if err != nil {
		panic(err)
	}

	if cookies != nil {
		req.AddCookie(cookies)
	}
	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	// Try to get the result
	response := transportGin.JSONResponse{}

	decoder := json.NewDecoder(res.Body)

	decoder.Decode(&response)

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
	ginC := transportGin.Connector{
		Waechter:   w,
		AuthPath:   "/auth",
		Domain:     "",
		Debug:      true,
		Writer:     GinkgoWriter,
		ForceHTTPS: false,
	}

	r := gin.Default()

	ginC.DefaultRoutes(r)

	s := httptest.NewServer(r)

	fmt.Println(fmt.Sprint(s.URL, "/auth/register"))

	Describe("Register", func() {

		It("Should register a new user", func() {
			res := struct {
				Token string `json:"token"`
			}{}
			authErr := makeRequest("POST", fmt.Sprint(s.URL, "/auth/register"), waechter.UserRegisterParams{
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
			authErr := makeRequest("GET", fmt.Sprint(s.URL, "/auth/verify-email?token="+verificationToken+"&id="+user.ID), nil, nil, nil)

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

			authErr := makeRequest("POST", fmt.Sprint(s.URL, "/auth/login/username-or-email"), waechter.UserLoginEmailOrUsernameParams{
				UsernameOrEmail: "ElectricCookie",
				Password:        "Password123",
				RememberMe:      false,
			}, nil, nil)

			Expect(authErr).To(BeNil())

		})

		It("should fail to log in somebody with the wrong credentials", func() {

			authErr := makeRequest("POST", fmt.Sprint(s.URL, "/auth/login/username-or-email"), waechter.UserLoginEmailOrUsernameParams{
				UsernameOrEmail: "ElectricCookie",
				Password:        "Password1234",
				RememberMe:      false,
			}, nil, nil)

			Expect(authErr).ToNot(BeNil())

			authErr = makeRequest("POST", fmt.Sprint(s.URL, "/auth/login/username-or-email"), waechter.UserLoginEmailOrUsernameParams{
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
			authErr := makeRequest("POST", fmt.Sprint(s.URL, "/auth/forgot-password"), waechter.ForgotPasswordParams{
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

			authErr := makeRequest("POST", fmt.Sprint(s.URL, "/auth/reset-password"), waechter.ResetPasswordParams{
				UserID:      user.ID,
				NewPassword: "newPassword123",
				Token:       resetToken,
			}, nil, nil)

			Expect(authErr).To(BeNil())

			authErr = makeRequest("POST", fmt.Sprint(s.URL, "/auth/login/username-or-email"), waechter.UserLoginEmailOrUsernameParams{
				UsernameOrEmail: "ElectricCookie",
				Password:        "newPassword123",
				RememberMe:      false,
			}, nil, nil)

			Expect(authErr).To(BeNil())

		})

	})

	AfterSuite(func() {
		// Close the server to allow the program to shut down properly.
		s.Close()

	})

})

func TestGinTransporter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoWaechter Suite")
}
