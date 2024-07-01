package i18n

import (
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

var translation gjson.Result

func Init(lang string) {
	translationJson, ok := translations[lang]

	if !ok {
		translationJson = translations[language.English.String()]
	}

	translation = gjson.Parse(translationJson)
}

func Get(key string) string {
	result := translation.Get(key)

	if !result.Exists() || result.IsObject() {
		return key
	}

	return result.String()
}
