package i18n

import (
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

var translationJson string

func Init(lang string) {
	var ok bool
	translationJson, ok = translations[lang]

	if !ok {
		translationJson = translations[language.English.String()]
	}
}

func Get(key string) string {
	result := gjson.Get(translationJson, key)

	if !result.Exists() || result.IsObject() {
		return key
	}

	return result.String()
}
