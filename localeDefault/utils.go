package localeDefault

import (
	"fmt"
	"html/template"
)

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

//LoadTemplate concatenates templates from bindata
func LoadTemplate(names ...string) *template.Template {
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

//InlineTemplate creates a template from the string passed to it
func InlineTemplate(templ string) *template.Template {

	return template.Must(template.New("inline-template").Parse(templ))

}
