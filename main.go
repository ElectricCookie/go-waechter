package waechter

import "time"

var dbAdapter DBAdapter
var parameters Parameters
var emailAdapter EmailAdapter

//Parameters define some global configuration options of the package
type Parameters struct {
	jwtSecret              string
	jwtIssuer              string
	jwtAccessTokenLifetime time.Duration
	requireInvite          bool
	requireActivation      bool

	sessionDuration int64
}

//Config configures the package
func Config(newDBAdapter DBAdapter, newEmailAdapter EmailAdapter, params Parameters) {

	dbAdapter = newDBAdapter
	emailAdapter = newEmailAdapter
	parameters = params
}

func getAdapter() DBAdapter {

	return dbAdapter

}

func getEmailAdapter() EmailAdapter {
	return emailAdapter
}

func getParameters() *Parameters {
	return &parameters
}

func main() {

}
