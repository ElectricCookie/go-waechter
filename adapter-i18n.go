package waechter

import (
	"html/template"
	"strings"
)

type DefaultTranslations struct {
	CompanyName     string
	CompanyWebsite  string
	LogoUrl         string
	Locales         []*Translation
	DefaultLanguage string
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

func (translations *DefaultTranslations) GetRegistrationEmail(lang string, user *User) (*Email, error) {

}

//Translation is used to translate the default emails and messages
type Translation struct {
	LanguageCode string

	ActiviationSubject   *template.Template
	ActivationContent    *template.Template
	ActivationButtonText *template.Template
	ActivationTitle      *template.Template
	ActivationSender     *template.Template

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

func loadTemplate(name string) *template.Template {

	return template.Must(template.ParseFiles("./templates/" + name))

}

func inlineTemplate(templ string) *template.Template {
	t := template.Template{}
	return template.Must(t.Parse(templ))

}

//DefaultTranslations loads default templates and strings for german and english
func DefaultTranslations() []*Translation {

	return []*Translation{
		&Translation{
			LanguageCode: "en",

			ActivationSender:     inlineTemplate("activation@{{ .CompanyWebsite }}"),
			ActiviationSubject:   inlineTemplate("Welcome at {{ .CompanyName }}! Activate your account"),
			ActivationTitle:      inlineTemplate("Welcome {{ .Username }},"),
			ActivationContent:    inlineTemplate("thanks for registering at {{ .CompanyName }}! To use your account please click the link below to verify your email address."),
			ActivationButtonText: inlineTemplate("Verify email address"),
		},
		&Translation{
			LanguageCode: "de",

			ActivationSender:     inlineTemplate("aktivierung@{{ .CompanyWebsite }}"),
			ActiviationSubject:   inlineTemplate("Willkommen bei {{ .CompanyName }}! Aktivieren Sie jetzt Ihren Account"),
			ActivationTitle:      inlineTemplate("Willkommen {{ .Username }},"),
			ActivationContent:    inlineTemplate("vielen Dank f&uuml;r Ihre Registrierung bei{{ .CompanyName }}! Um Ihren Account benutzen zu k&ouml;nnen, folgen Sie bitten dem Link."),
			ActivationButtonText: inlineTemplate("E-Mail Adresse Verifizieren"),
		},
	}

}
