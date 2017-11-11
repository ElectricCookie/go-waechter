package transportGin

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
)

//ResetPassword resets the password of a user
func (connector *Connector) ResetPassword(context *gin.Context) {

	parameters := waechter.UserResetPasswordParams{}

	if !connector.bindParameters(context, &parameters) {
		return //Error occurred
	}

	err := connector.Waechter.UserResetPassword(parameters)

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}
