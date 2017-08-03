package localeDefault

var de = Translation{
	"activation.sender":  InlineTemplate("accounts@{{ .CompanyWebsite }}"),
	"activation.subject": InlineTemplate("Willkommen bei {{ .CompanyName }}! Aktivieren Sie jetzt Ihren Account"),
	"activation.content": LoadTemplate("localeDefault/templates/confirm.html", "localeDefault/templates/confirm-de.html"),

	"forgotPassword.sender":  InlineTemplate("accounts@{{ .CompanyWebsite }}"),
	"forgotPassword.subject": InlineTemplate("Ihr Passwort bei {{ .CompanyName }} zurücksetzen"),
	"forgotPassword.content": LoadTemplate("localeDefault/templates/forgot-password-de.html", "localeDefault/templates/forgot-password.html"),

	"resetPassword.sender":  InlineTemplate("accounts@{{ .CompanyWebsite }}"),
	"resetPassword.subject": InlineTemplate("Ihr Passwort bei {{.CompanyName}} wurde geändert."),
	"resetPassword.content": LoadTemplate("localeDefault/templates/password-changed-de.html", "localeDefault/templates/password-changed.html"),
}
