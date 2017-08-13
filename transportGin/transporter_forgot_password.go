package transportGin

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
)

//ForgotPassword requests a reset password email
func (connector *Connector) ForgotPassword(context *gin.Context) {

	parameters := waechter.ForgotPasswordParams{}

	if !connector.bindParameters(context, &parameters) {
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
