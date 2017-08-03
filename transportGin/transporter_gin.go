package transportGin

import (
	"encoding/json"
	"io"

	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
)

//GinConnector is used to setup gin with go-waechter
type GinConnector struct {
	Waechter   *waechter.Waechter
	Debug      bool
	AuthPath   string
	Domain     string
	Writer     io.Writer
	ForceHTTPS bool
}

//JSONResponse describes the format in which data is returned from the connector
type JSONResponse struct {
	Status bool                `json:"status"`
	Err    *waechter.AuthError `json:"err"`
	Data   interface{}         `json:"data"`
}

func (connector *GinConnector) bindParamters(c *gin.Context, parameters interface{}) bool {
	err := c.BindJSON(parameters)
	if err != nil {
		connector.respondHTTPDefault(nil, waechter.InvalidParametersError(err), c)
		return false
	}
	return true
}

//DefaultRoutes mounts routes under the /auth/ path
func (connector *GinConnector) DefaultRoutes(engine *gin.Engine) {

	engine.POST("/auth/login", connector.Login)
	engine.POST("/auth/register", connector.Register)
	engine.POST("/auth/forgotPassword", connector.ForgotPassword)
	engine.POST("/auth/resetPassword", connector.ResetPassword)
	engine.POST("/auth/verifyEmail", connector.VerifyEmail)

}

func (connector *GinConnector) respondHTTPDefault(data interface{}, err *waechter.AuthError, context *gin.Context) {
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

//Login to an account
func (connector *GinConnector) Login(context *gin.Context) {
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
func (connector *GinConnector) Register(context *gin.Context) {
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
func (connector *GinConnector) VerifyEmail(context *gin.Context) {

	parameters := waechter.UserVerifyEmailParameters{}

	if !connector.bindParamters(context, &parameters) {
		return //Error occurred
	}

	err := connector.Waechter.UserVerifyEmailAddress(parameters)

	connector.respondHTTPDefault(true, err, context)

}

//ForgotPassword requests a reset password email
func (connector *GinConnector) ForgotPassword(context *gin.Context) {

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
func (connector *GinConnector) ResetPassword(context *gin.Context) {

	parameters := waechter.ResetPasswordParams{}

	if !connector.bindParamters(context, &parameters) {
		return //Error occurred
	}

	err := connector.Waechter.ResetPassword(parameters)

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}
