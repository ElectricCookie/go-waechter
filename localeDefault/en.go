package localeDefault

var en = Translation{
	"activation.sender":  InlineTemplate("accounts@{{ .CompanyWebsite }}"),
	"activation.subject": InlineTemplate("Welcome at {{ .CompanyName }}! Verify your email address"),
	"activation.content": LoadTemplate("localeDefault/templates/confirm-en.html", "localeDefault/templates/confirm.html"),

	"forgotPassword.sender":  InlineTemplate("accounts@{{ .CompanyWebsite }}"),
	"forgotPassword.subject": InlineTemplate("Restore your password at {{ .CompanyName }}"),
	"forgotPassword.content": LoadTemplate("localeDefault/templates/forgot-password-en.html", "localeDefault/templates/forgot-password.html"),

	"resetPassword.sender":  InlineTemplate("accounts@{{ .CompanyWebsite }}"),
	"resetPassword.subject": InlineTemplate("Your password has been changed"),
	"resetPassword.content": LoadTemplate("localeDefault/templates/password-changed-en.html", "localeDefault/templates/password-changed.html"),
}
