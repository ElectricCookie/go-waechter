package transportGin

import (
	waechter "github.com/ElectricCookie/go-waechter"
	"github.com/gin-gonic/gin"
)

// Register a new account
func (connector *Connector) Register(context *gin.Context) {
	// Retrieve data

	params := waechter.UserRegisterParams{}

	if !connector.bindParameters(context, &params) {
		return // Error occurred
	}

	err := connector.Waechter.UserRegister(params)

	if err != nil {
		connector.respondHTTPDefault(nil, err, context)
		return
	}

	token, err := connector.Waechter.UserSendVerificationEmail(waechter.UserSendVerficationParameters{
		Email: params.Email,
	})

	if connector.Debug {
		connector.respondHTTPDefault(struct {
			Token string `json:"token"`
		}{Token: *token}, err, context)
	} else {
		connector.respondHTTPDefault(struct{ status bool }{true}, err, context)
	}

}
