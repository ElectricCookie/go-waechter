package waechter

import (
	"time"
)

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
	SessionDurationDefault    *time.Duration
	SessionDurationRememberMe *time.Duration
	// Adapters
	DbAdapter    DBAdapter
	Locales      LocalesAdapter
	EmailAdapter EmailAdapter
}

func (waechter *Waechter) getDBAdapter() DBAdapter {
	return waechter.DbAdapter
}

func (waechter *Waechter) getEmailAdapter() EmailAdapter {
	return waechter.EmailAdapter
}

func (waechter *Waechter) getLocales() LocalesAdapter {
	return waechter.Locales
}

//New creates a new waechter
func New(jwtSecret string, jwtIssuer string, dbAdapter DBAdapter, emailAdapter EmailAdapter, translations LocalesAdapter) *Waechter {

	w := Waechter{
		JwtSecret:                 jwtSecret,
		JwtIssuer:                 jwtIssuer,
		DbAdapter:                 dbAdapter,
		SessionDurationDefault:    nil,
		SessionDurationRememberMe: nil,
		EmailAdapter:              emailAdapter,
		Locales:                   translations,
	}

	return &w
}
