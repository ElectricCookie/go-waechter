package localeDefault

import (
	"bytes"
	"html/template"
	"strings"

	waechter "github.com/ElectricCookie/go-waechter"
)

var requiredTemplates = []string{
	"activation.sender",
	"activation.subject",
	"activation.content",

	"forgotPassword.sender",
	"forgotPassword.subject",
	"forgotPassword.content",

	"resetPassword.sender",
	"resetPassword.subject",
	"resetPassword.content",
}

//Translation is a map of locale identifiers and templates
type Translation map[string]*template.Template

//TranslationParameters contain parameters passed to a template
type TranslationParameters map[string]interface{}

//DefaultTranslations uses some predefined translations and templates for waechter
type DefaultTranslations struct {
	CompanyName          string
	CompanyWebsite       string
	LogoURL              string
	Locales              map[string]Translation
	DefaultLanguage      string
	UserVerifyEmailAddress   string
	ResetPasswordAddress string
}

//NewDefaultTranslations creates a new DefaultTranslation and loads en and de locales
func NewDefaultTranslations() *DefaultTranslations {
	translation := &DefaultTranslations{}

	translation.DefaultLanguage = "en"
	translation.Locales = make(map[string]Translation)
	translation.LoadLocale("en", en)
	translation.LoadLocale("de", de)
	return translation
}

//LoadLocale registers a new translation
func (def *DefaultTranslations) LoadLocale(langCode string, translation Translation) {

	for _, item := range requiredTemplates {

		if template, ok := translation[item]; !ok || template == nil {
			panic("Could not load locale. Missing " + item + " in translation: " + langCode)
		}

	}

	def.Locales[langCode] = translation

}

//GetTranslation returns a *Translation according to the language code passed. Returns the default translation if its unknown
func (def *DefaultTranslations) GetTranslation(lang string) Translation {
	lang = strings.ToLower(lang)

	if item, ok := def.Locales[lang]; ok {
		return item
	}

	return def.GetTranslation(def.DefaultLanguage)

}

func (def *DefaultTranslations) getString(template *template.Template, user *waechter.User, params *TranslationParameters) string {

	if template == nil {
		panic("Missing translation.")
	}

	var buf bytes.Buffer

	defaultParams := TranslationParameters{
		"Username":       user.Username,
		"Email":          user.Email,
		"LogoURL":        def.LogoURL,
		"CompanyName":    def.CompanyName,
		"CompanyWebsite": def.CompanyWebsite,
	}

	finalParams := joinMaps(&defaultParams, params)

	template.Execute(&buf, finalParams)

	return buf.String()
}

//GetRegistrationEmail returns an email struct for a new user
func (def *DefaultTranslations) GetRegistrationEmail(user *waechter.User, verificationToken string) (*waechter.Email, error) {

	trans := def.GetTranslation(user.Language)

	return &waechter.Email{
		Subject: def.getString(trans["activation.subject"], user, nil),
		From:    def.getString(trans["activation.sender"], user, nil),
		To:      user.Email,
		Content: def.getString(trans["activation.content"], user, &TranslationParameters{
			"ConfirmAddress": def.UserVerifyEmailAddress + "?id=" + user.ID + "&token=" + verificationToken,
		}),
	}, nil

}

//GetForgotPasswordEmail sends an email to the user with a link to reset their password
func (def *DefaultTranslations) GetForgotPasswordEmail(user *waechter.User, passwordToken string) (*waechter.Email, error) {

	trans := def.GetTranslation(user.Language)

	return &waechter.Email{
		Subject: def.getString(trans["forgotPassword.subject"], user, nil),
		From:    def.getString(trans["forgotPassword.sender"], user, nil),
		To:      user.Email,
		Content: def.getString(trans["forgotPassword.content"], user, &TranslationParameters{
			"ResetAddress": def.ResetPasswordAddress + "?id=" + user.ID + "&token=" + passwordToken,
		}),
	}, nil

}

//GetPasswordResetEmail return an email object for the notification of a reset password
func (def *DefaultTranslations) GetPasswordResetEmail(user *waechter.User) (*waechter.Email, error) {

	trans := def.GetTranslation(user.Language)

	return &waechter.Email{
		Subject: def.getString(trans["resetPassword.subject"], user, nil),
		From:    def.getString(trans["resetPassword.sender"], user, nil),
		To:      user.Email,
		Content: def.getString(trans["resetPassword.content"], user, &TranslationParameters{}),
	}, nil

}
