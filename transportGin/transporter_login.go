package transportGin

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
)

//LoginUsername ...
func (connector *Connector) LoginUsername(context *gin.Context) {
	// Retrieve data

	parameters := waechter.UserLoginUsernameParams{}

	if !connector.bindParameters(context, &parameters) {
		return // Error occurred
	}
	// Use waechter to log in
	refreshToken, err := connector.Waechter.UserLoginUsername(parameters)

	if err == nil {
		context.SetCookie("Waechter-RefreshToken", refreshToken, 2629743, connector.AuthPath, connector.Domain, connector.ForceHTTPS, true)
	}

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}

//LoginEmail ...
func (connector *Connector) LoginEmail(context *gin.Context) {
	// Retrieve data

	parameters := waechter.UserLoginEmailParams{}

	if !connector.bindParameters(context, &parameters) {
		return // Error occurred
	}
	// Use waechter to log in
	refreshToken, err := connector.Waechter.UserLoginEmail(parameters)

	if err == nil {
		context.SetCookie("Waechter-RefreshToken", refreshToken, 2629743, connector.AuthPath, connector.Domain, connector.ForceHTTPS, true)
	}

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}

//LoginUsernameOrEmail ...
func (connector *Connector) LoginUsernameOrEmail(context *gin.Context) {
	// Retrieve data

	parameters := waechter.UserLoginEmailOrUsernameParams{}

	if !connector.bindParameters(context, &parameters) {
		return // Error occurred
	}
	// Use waechter to log in
	refreshToken, err := connector.Waechter.UserLoginWithUsernameOrEmail(parameters)

	if err == nil {
		context.SetCookie("Waechter-RefreshToken", refreshToken, 2629743, connector.AuthPath, connector.Domain, connector.ForceHTTPS, true)
	}

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}
