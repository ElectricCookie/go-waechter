package waechter

//Waechter wraps a waechter instance
type Waechter struct {
	// JWT Information
	JwtSecret string
	JwtIssuer string
	// Email Address of administrator
	AdminEmail string
	// Notifications
	NotificationLogin           bool
	NotificationRegisterd       bool
	NotificationAdminRegistered bool
	// Registration related
	RequireInvite     bool
	RequireActivation bool
	// Login related
	SessionDurationDefault    int64
	SessionDurationRememberMe int64
	// Adapters
	DbAdapter    DBAdapter
	Locales      TranslationAdapter
	EmailAdapter EmailAdapter
}

func (waechter *Waechter) getDBAdapter() DBAdapter {
	return waechter.DbAdapter
}

func (waechter *Waechter) getEmailAdapter() EmailAdapter {
	return waechter.EmailAdapter
}

//New creates a new waechter
func New(jwtSecret string, jwtIssuer string, dbAdapter DBAdapter, emailAdapter EmailAdapter, translations TranslationAdapter) *Waechter {

	w := Waechter{
		JwtSecret:    jwtSecret,
		JwtIssuer:    jwtIssuer,
		DbAdapter:    dbAdapter,
		EmailAdapter: emailAdapter,
		Locales:      translations,
	}

	return &w
}
