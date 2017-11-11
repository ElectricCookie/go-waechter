package transportGin

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
)

//RequestVerification ...
func (connector *Connector) RequestVerification(context *gin.Context) {

	parameters := waechter.UserSendVerficationParams{}

	if !connector.bindParameters(context, &parameters) {
		return // Error occurred
	}

	_, err := connector.Waechter.UserSendVerificationEmail(parameters)

	connector.respondHTTPDefault(struct{ status bool }{true}, err, context)

}

//VerifyEmail of a new account
func (connector *Connector) VerifyEmail(context *gin.Context) {

	err := connector.Waechter.UserVerifyEmailAddress(waechter.UserVerifyEmailParams{
		UserID: context.Query("id"),
		Token:  context.Query("token"),
	})

	if err != nil {
		context.Redirect(303, connector.RedirectVerifyFailed)
	} else {
		context.Redirect(303, connector.RedirectVerifySuccess)
	}

}
