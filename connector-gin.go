package waechter

import "github.com/gin-gonic/gin"

//GinConnector is used to setup gin with go-waechter
type GinConnector struct {
	waechter   *Waechter
	authPath   string
	domain     string
	forceHTTPS true
}

func respondHTTPDefault(data interface{}, err *AuthError, context *gin.Context) {
	if err != nil {
		// Make sure internal errors are not passed to the outer world
		if err.IsInternal {
			err.ErrorCode = "internalError"
			err.Description = "Internal Error occured. "
			err.Err = nil
		}
		context.JSON(500, err)
	} else {
		context.JSON(200, data)
	}
}

//Login to an account
func (connector *GinConnector) Login(context *gin.Context) {
	// Retrieve data
	usernameOrEmail := context.PostForm("usernameOrEmail")
	password := context.PostForm("password")
	// Use waechter to log in
	refreshToken, err := connector.waechter.LoginWithUsernameOrEmail(LoginEmailOrUsernameData{
		UsernameOrEmail: usernameOrEmail,
		Password:        password,
	})

	if err != nil {
		context.SetCookie("Waechter-RefreshToken", *refreshToken, 2629743, connector.authPath, connector.domain, connector.forceHTTPS, true)
	}

	respondHTTPDefault(struct{ status bool }{true}, err, context)

}
