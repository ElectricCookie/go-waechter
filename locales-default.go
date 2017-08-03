package waechter

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

//TranslationParameters contain parameters passed to a template
type TranslationParameters map[string]interface{}

//DefaultTranslations uses some predefined translations and templates for waechter
type DefaultTranslations struct {
	CompanyName          string
	CompanyWebsite       string
	LogoURL              string
	Locales              []*Translation
	DefaultLanguage      string
	VerifyEmailAddress   string
	ResetPasswordAddress string
}

//GetDefaultLocales loads default templates and strings for german and english
func GetDefaultLocales() []*Translation {

	return []*Translation{
		&Translation{
			LanguageCode: "en",

			ActivationSender:   inlineTemplate("accounts@{{ .CompanyWebsite }}"),
			ActiviationSubject: inlineTemplate("Welcome at {{ .CompanyName }}! Verify your email address"),
			ActivationContent:  loadTemplate("templates/confirm-en.html", "templates/confirm.html"),

			ForgotPasswordSender:  inlineTemplate("accounts@{{ .CompanyWebsite }}"),
			ForgotPasswordSubject: inlineTemplate("Restore your password at {{ .CompanyName }}"),
			ForgotPasswordContent: loadTemplate("templates/forgot-password-en.html", "templates/forgot-password.html"),

			PasswordResetSender:  inlineTemplate("accounts@{{ .CompanyWebsite }}"),
			PasswordResetSubject: inlineTemplate("Your password has been changed"),
			PasswordResetContent: loadTemplate("templates/password-changed-en.html", "templates/password-changed.html"),
		},
		&Translation{
			LanguageCode: "de",

			ActivationSender:   inlineTemplate("accounts@{{ .CompanyWebsite }}"),
			ActiviationSubject: inlineTemplate("Willkommen bei {{ .CompanyName }}! Aktivieren Sie jetzt Ihren Account"),
			ActivationContent:  loadTemplate("templates/confirm.html", "templates/confirm-de.html"),

			ForgotPasswordSender:  inlineTemplate("accounts@{{ .CompanyWebsite }}"),
			ForgotPasswordSubject: inlineTemplate("Ihr Passwort bei {{ .CompanyName }} zurücksetzen"),
			ForgotPasswordContent: loadTemplate("templates/forgot-password-de.html", "templates/forgot-password.html"),

			PasswordResetSender:  inlineTemplate("accounts@{{ .CompanyWebsite }}"),
			PasswordResetSubject: inlineTemplate("Ihr Passwort wurde geändert."),
			PasswordResetContent: loadTemplate("templates/password-changed-de.html", "templates/password-changed.html"),
		},
	}

}

//GetTranslation returns a *Translation according to the language code passed. Returns the default translation if its unknown
func (adapter *DefaultTranslations) GetTranslation(lang string) *Translation {
	lang = strings.ToLower(lang)
	for _, t := range adapter.Locales {
		if t.LanguageCode == lang {
			return t
		}
	}

	// No language found

	return adapter.GetTranslation(adapter.DefaultLanguage)

}

func joinMaps(maps ...*TranslationParameters) *TranslationParameters {

	result := TranslationParameters{}

	for _, item := range maps {
		if item != nil {
			for key, value := range *item {
				result[key] = value
			}
		}

	}

	return &result
}

func (adapter *DefaultTranslations) getString(template *template.Template, user *User, params *TranslationParameters) string {

	var buf bytes.Buffer

	defaultParams := TranslationParameters{
		"Username":       user.Username,
		"Email":          user.Email,
		"LogoURL":        adapter.LogoURL,
		"CompanyName":    adapter.CompanyName,
		"CompanyWebsite": adapter.CompanyWebsite,
	}

	finalParams := joinMaps(&defaultParams, params)

	template.Execute(&buf, finalParams)

	return buf.String()
}

//Translation is used to translate the default emails and messages
type Translation struct {
	LanguageCode string

	ActiviationSubject *template.Template
	ActivationContent  *template.Template
	ActivationSender   *template.Template

	ForgotPasswordSubject *template.Template
	ForgotPasswordContent *template.Template
	ForgotPasswordSender  *template.Template

	PasswordResetSubject *template.Template
	PasswordResetContent *template.Template
	PasswordResetSender  *template.Template

	LoginNotificationSubject string
	LoginNotificationContent string
	LoginNotificationSender  string

	UsernameChangedSubject string
	UsernameChangedContent string
	UsernameChangedSender  string
}

func loadTemplate(names ...string) *template.Template {
	finalString := ""
	for _, name := range names {
		if file, err := Asset(name); err != nil {
			fmt.Printf("Failed to load asset: " + name)
			panic(err)
		} else {
			finalString += string(file) + "\n"

		}
	}

	return template.Must(template.New("file-template").Parse(finalString))

}

func inlineTemplate(templ string) *template.Template {

	return template.Must(template.New("inline-template").Parse(templ))

}

//GetRegistrationEmail returns an email struct for a new user
func (adapter *DefaultTranslations) GetRegistrationEmail(user *User, verificationToken string) (*Email, error) {

	trans := adapter.GetTranslation(user.Language)

	return &Email{
		Subject: adapter.getString(trans.ActiviationSubject, user, nil),
		From:    adapter.getString(trans.ActivationSender, user, nil),
		To:      user.Email,
		Content: adapter.getString(trans.ActivationContent, user, &TranslationParameters{
			"ConfirmAddress": adapter.VerifyEmailAddress + "?id=" + user.ID + "&token=" + verificationToken,
		}),
	}, nil

}

//GetForgotPasswordEmail sends an email to the user with a link to reset their password
func (adapter *DefaultTranslations) GetForgotPasswordEmail(user *User, passwordToken string) (*Email, error) {

	trans := adapter.GetTranslation(user.Language)

	return &Email{
		Subject: adapter.getString(trans.ForgotPasswordSubject, user, nil),
		From:    adapter.getString(trans.ForgotPasswordSender, user, nil),
		To:      user.Email,
		Content: adapter.getString(trans.ForgotPasswordContent, user, &TranslationParameters{
			"ResetAddress": adapter.ResetPasswordAddress + "?id=" + user.ID + "&token=" + passwordToken,
		}),
	}, nil

}

//GetPasswordResetEmail return an email object for the notification of a reset password
func (adapter *DefaultTranslations) GetPasswordResetEmail(user *User) (*Email, error) {

	trans := adapter.GetTranslation(user.Language)

	return &Email{
		Subject: adapter.getString(trans.PasswordResetSubject, user, nil),
		From:    adapter.getString(trans.PasswordResetSender, user, nil),
		To:      user.Email,
		Content: adapter.getString(trans.PasswordResetContent, user, &TranslationParameters{}),
	}, nil

}
