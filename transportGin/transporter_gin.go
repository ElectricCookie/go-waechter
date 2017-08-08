package transportGin

import (
	"encoding/json"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"

	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
)

//Connector is used to setup gin with go-waechter
type Connector struct {
	Waechter           *waechter.Waechter
	Debug              bool
	AuthPath           string
	Domain             string
	Writer             io.Writer
	ForceHTTPS         bool
	AccesTokenLifetime time.Duration
}

//JSONResponse describes the format in which data is returned from the connector
type JSONResponse struct {
	Status bool                `json:"status"`
	Err    *waechter.AuthError `json:"err"`
	Data   interface{}         `json:"data"`
}

func (connector *Connector) bindParamters(c *gin.Context, parameters interface{}) bool {
	err := c.BindJSON(parameters)
	if err != nil {
		connector.respondHTTPDefault(nil, waechter.InvalidParametersError(err), c)
		return false
	}
	return true
}

//DefaultRoutes mounts routes under the /auth/ path
func (connector *Connector) DefaultRoutes(engine *gin.Engine) {

	engine.POST("/auth/login/username-or-email", connector.LoginUsernameOrEmail)
	engine.POST("/auth/login/username", connector.LoginUsername)
	engine.POST("/auth/login/email", connector.LoginEmail)
	engine.POST("/auth/register", connector.Register)
	engine.POST("/auth/forgotPassword", connector.ForgotPassword)
	engine.POST("/auth/resetPassword", connector.ResetPassword)
	engine.POST("/auth/verifyEmail", connector.VerifyEmail)

}

func (connector *Connector) respondHTTPDefault(data interface{}, err *waechter.AuthError, context *gin.Context) {
	if err != nil {
		if connector.Debug {
			s, _ := json.Marshal(err)
			connector.Writer.Write(s)
		}

		// Make sure internal errors are not passed to the outer world
		if err.IsInternal {
			err.ErrorCode = "internalError"
			err.Description = "Internal Error occurred. "
			err.Err = nil
		}

		context.JSON(500, JSONResponse{Status: false, Err: err, Data: nil})
	} else {
		context.JSON(200, JSONResponse{Status: true, Err: nil, Data: data})
	}
}

//LoginUsername ...
func (connector *Connector) LoginUsername(context *gin.Context) {
	// Retrieve data

	parameters := waechter.UserLoginUsernameData{}

	if !connector.bindParamters(context, &parameters) {
		return // Error occurred
	}
	// Use waechter to log in
	refreshToken, err := connector.Waechter.UserLoginUsername(parameters)

	if err == nil {
		context.SetCookie("Waechter-RefreshToken", *refreshToken, 2629743, connector.AuthPath, connector.Domain, connector.ForceHTTPS, true)
	}

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}

//LoginEmail ...
func (connector *Connector) LoginEmail(context *gin.Context) {
	// Retrieve data

	parameters := waechter.UserLoginEmailData{}

	if !connector.bindParamters(context, &parameters) {
		return // Error occurred
	}
	// Use waechter to log in
	refreshToken, err := connector.Waechter.UserLoginEmail(parameters)

	if err == nil {
		context.SetCookie("Waechter-RefreshToken", *refreshToken, 2629743, connector.AuthPath, connector.Domain, connector.ForceHTTPS, true)
	}

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}

//LoginUsernameOrEmail ...
func (connector *Connector) LoginUsernameOrEmail(context *gin.Context) {
	// Retrieve data

	parameters := waechter.UserLoginEmailOrUsernameData{}

	if !connector.bindParamters(context, &parameters) {
		return // Error occurred
	}
	// Use waechter to log in
	refreshToken, err := connector.Waechter.UserLoginWithUsernameOrEmail(parameters)

	if err == nil {
		context.SetCookie("Waechter-RefreshToken", *refreshToken, 2629743, connector.AuthPath, connector.Domain, connector.ForceHTTPS, true)
	}

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}

// Register a new account
func (connector *Connector) Register(context *gin.Context) {
	// Retrieve data

	params := waechter.UserRegisterParams{}

	if !connector.bindParamters(context, &params) {
		return // Error occurred
	}

	err := connector.Waechter.UserRegister(params)

	if err != nil {
		connector.respondHTTPDefault(nil, err, context)
		return
	}

	token, err := connector.Waechter.SendVerificationEmail(params.Email)

	if connector.Debug {
		connector.respondHTTPDefault(struct {
			Token string `json:"token"`
		}{Token: *token}, err, context)
	} else {
		connector.respondHTTPDefault(struct{ status bool }{true}, err, context)
	}

}

//VerifyEmail of a new account
func (connector *Connector) VerifyEmail(context *gin.Context) {

	parameters := waechter.UserVerifyEmailParameters{}

	if !connector.bindParamters(context, &parameters) {
		return //Error occurred
	}

	err := connector.Waechter.UserVerifyEmailAddress(parameters)

	connector.respondHTTPDefault(true, err, context)

}

//ForgotPassword requests a reset password email
func (connector *Connector) ForgotPassword(context *gin.Context) {

	parameters := waechter.ForgotPasswordParams{}

	if !connector.bindParamters(context, &parameters) {
		return //Error occurred
	}

	token, err := connector.Waechter.ForgotPassword(parameters)

	if connector.Debug {
		connector.respondHTTPDefault(struct {
			Token string `json:"token"`
		}{Token: *token}, err, context)
	} else {
		connector.respondHTTPDefault(struct{ status bool }{true}, err, context)
	}

}

//ResetPassword resets the password of a user
func (connector *Connector) ResetPassword(context *gin.Context) {

	parameters := waechter.ResetPasswordParams{}

	if !connector.bindParamters(context, &parameters) {
		return //Error occurred
	}

	err := connector.Waechter.ResetPassword(parameters)

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}

func (connector *Connector) GetRefreshToken(context *gin.Context) (string, error) {

	refreshToken, err := context.Cookie("Waechter-RefreshToken")

	if err != nil {
		return "", err
	}

	return refreshToken, nil

}

func (connector *Connector) CreateAccessToken(target string, context *gin.Context, verify func(*jwt.StandardClaims, *waechter.User) *waechter.AuthError) {

	// Retrieve refresh token

	refreshToken, err := connector.GetRefreshToken(context)

	if err != nil {
		connector.respondHTTPDefault(nil, waechter.NotLoggedInError(), context)
		return
	}

	user, claims, err := connector.Waechter.CheckRefreshToken(refreshToken)

	if err != nil {
		connector.respondHTTPDefault(nil, waechter.NotLoggedInError(), context)
		return
	}

	// Verify

	authErr := verify(claims, user)

	if authErr != nil {
		connector.respondHTTPDefault(nil, authErr, context)
		return
	}

	accessToken, authErr := connector.Waechter.GenerateAccessToken(refreshToken, target)

	if authErr != nil {
		connector.respondHTTPDefault(nil, waechter.NotLoggedInError(), context)
		return
	}

	context.SetCookie("Waechter-Access:"+target, *accessToken, int(connector.AccesTokenLifetime.Seconds()), "", "", connector.ForceHTTPS, true)

}

func (connector *Connector) GetAccessToken(target string, context *gin.Context) (string, *waechter.AuthError) {

	accessToken, err := context.Cookie("Waechter-Access:" + target)

	if err != nil {

		authErr := &waechter.AuthError{
			ErrorCode:   "notEnoughRights",
			Description: "No Access token for this action present.",
		}
		return "", authErr
	}

	return accessToken, nil

}

func (connector *Connector) CheckAccessToken(target string, context *gin.Context) (*jwt.StandardClaims, *waechter.AuthError) {
	var accessToken string
	accessToken, err := connector.GetAccessToken(target, context)
	if err != nil {
		return nil, err
	}
	return connector.Waechter.CheckAccessToken(target, accessToken)

}
