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
	ConfirmAddress       string
	ResetPasswordAddress string
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

//GetRegistrationEmail returns an email struct for a new user
func (adapter *DefaultTranslations) GetRegistrationEmail(lang string, user *User, verificationToken string) (*Email, error) {

	trans := adapter.GetTranslation(lang)

	return &Email{
		Subject: adapter.getString(trans.ActiviationSubject, user, nil),
		From:    adapter.getString(trans.ActivationSender, user, nil),
		To:      user.Email,
		Content: adapter.getString(trans.ActivationContent, user, &TranslationParameters{
			"ConfirmAddress": adapter.ConfirmAddress + "?id=" + user.ID + "&token=" + verificationToken,
		}),
	}, nil

}

//Translation is used to translate the default emails and messages
type Translation struct {
	LanguageCode string

	ActiviationSubject *template.Template
	ActivationContent  *template.Template
	ActivationTitle    *template.Template
	ActivationSender   *template.Template

	ForgotPasswordSubject string
	ForgotPasswordContent string
	ForgotPasswordSender  string

	PasswordResetSubject string
	PasswordResetContent string
	PasswordResetSender  string

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

//GetDefaultLocales loads default templates and strings for german and english
func GetDefaultLocales() []*Translation {

	return []*Translation{
		&Translation{
			LanguageCode: "en",

			ActivationSender:   inlineTemplate("activation@{{ .CompanyWebsite }}.com"),
			ActiviationSubject: inlineTemplate("Welcome at {{ .CompanyName }}! Activate your account"),
			ActivationTitle:    inlineTemplate("Welcome {{ .Username }},"),
			ActivationContent:  loadTemplate("templates/confirm-en.html", "templates/confirm.html"),
		},
		&Translation{
			LanguageCode: "de",

			ActivationSender:   inlineTemplate("aktivierung@{{ .CompanyWebsite }}"),
			ActiviationSubject: inlineTemplate("Willkommen bei {{ .CompanyName }}! Aktivieren Sie jetzt Ihren Account"),
			ActivationTitle:    inlineTemplate("Willkommen {{ .Username }},"),
			ActivationContent:  loadTemplate("templates/confirm.html", "templates/confirm-de.html"),
		},
	}

}
