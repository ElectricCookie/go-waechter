package transportGin

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"

	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//Connector is used to setup gin with go-waechter
type Connector struct {
	Waechter              *waechter.Waechter
	Debug                 bool
	AuthPath              string
	Domain                string
	Writer                io.Writer
	ForceHTTPS            bool
	RedirectVerifySuccess string
	RedirectVerifyFailed  string
}

//JSONResponse describes the format in which data is returned from the connector
type JSONResponse struct {
	Status bool                `json:"status"`
	Err    *waechter.AuthError `json:"err"`
	Data   interface{}         `json:"data"`
}

// BindJSON is a shortcut for c.BindWith(obj, binding.JSON)
func BindJSON(obj interface{}, c *gin.Context) error {
	return BindWith(obj, binding.JSON, c)
}

// BindWith binds the passed struct pointer using the specified binding engine.
// See the binding package.
func BindWith(obj interface{}, b binding.Binding, c *gin.Context) error {
	if err := b.Bind(c.Request, obj); err != nil {
		//c.AbortWithError(400, err).SetType(ErrorTypeBind)
		return err
	}
	return nil
}

func (connector *Connector) bindParameters(c *gin.Context, parameters interface{}) bool {
	err := BindJSON(parameters, c)

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
	engine.POST("/auth/forgot-password", connector.ForgotPassword)
	engine.POST("/auth/reset-password", connector.ResetPassword)
	engine.GET("/auth/verify-email", connector.VerifyEmail)
	engine.POST("/auth/request-verification", connector.RequestVerification)

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

		context.JSON(200, JSONResponse{Status: false, Err: err, Data: nil})
	} else {
		context.JSON(200, JSONResponse{Status: true, Err: nil, Data: data})
	}
}

type VerifyFunc func(claims *jwt.StandardClaims, user *waechter.User) (bool, time.Duration)
type HandleFunc func(claims *jwt.StandardClaims, userId string)

//Enfore enforces a permission on a route
func (connector *Connector) Enforce(target string, path string, c *gin.Context, verify VerifyFunc, handle HandleFunc) {

	// Check if theres an access token
	claims, err := connector.CheckAccessToken(target, c)

	if err == nil {
		handle(claims, claims.Subject)
	} else {
		user, claims, err := connector.GetRefreshToken(c)
		if err != nil {
			connector.respondHTTPDefault(nil, waechter.NotLoggedInError(), c)
		} else {
			if ok, expires := verify(claims, user); !ok {
				connector.respondHTTPDefault(nil, &waechter.AuthError{
					ErrorCode:   "notEnoughRights",
					Description: "No Access token for this action present.",
				}, c)
			} else {
				accessToken, authErr := connector.Waechter.UserGenerateAccessToken(claims, target, expires)
				if authErr != nil {
					fmt.Println(authErr)
					connector.respondHTTPDefault(nil, waechter.NotLoggedInError(), c)
					return
				}
				c.SetCookie(
					"Waechter-Access-"+target,
					accessToken,
					int(expires.Seconds()),
					path,
					"",
					connector.ForceHTTPS,
					true,
				)
				handle(claims, user.ID)
			}
		}
	}

}

func (connector *Connector) GetRefreshToken(context *gin.Context) (*waechter.User, *jwt.StandardClaims, error) {

	refreshToken, err := context.Cookie("Waechter-RefreshToken")

	if err != nil {
		return nil, nil, err
	}

	return connector.Waechter.UserCheckRefreshToken(refreshToken)

}

func (connector *Connector) CreateAccessToken(target string, path string, claims *jwt.StandardClaims, expires time.Duration, context *gin.Context) *waechter.AuthError {

	// Retrieve refresh token

	accessToken, authErr := connector.Waechter.UserGenerateAccessToken(claims, target, expires)

	if authErr != nil {
		return authErr
	}

	context.SetCookie("Waechter-Access-"+target, accessToken, int(expires.Seconds()), path, "", connector.ForceHTTPS, true)

	return nil
}

func (connector *Connector) GetAccessToken(target string, context *gin.Context) (string, *waechter.AuthError) {

	accessToken, err := context.Cookie("Waechter-Access-" + target)

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
	return connector.Waechter.UserCheckAccessToken(target, accessToken)

}
