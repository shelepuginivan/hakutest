// Package i18n provides internationalization for the application.
package i18n

import (
	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

var translation gjson.Result

// Init initializes internationalization with the language.
// It must be called at least once before `Get` is called.
func Init(lang string) {
	translationJson, ok := translations[lang]

	if !ok {
		translationJson = translations[language.English.String()]
	}

	translation = gjson.Parse(translationJson)
}

// Get returns translated string by key for the initial language.
// If the key does not exist, it returns the key.
// It must be called after `Init` was called at least once.
func Get(key string) string {
	result := translation.Get(key)

	if !result.Exists() || result.IsObject() {
		return key
	}

	return result.String()
}
