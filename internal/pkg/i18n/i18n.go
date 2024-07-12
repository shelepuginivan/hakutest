// Package i18n provides internationalization for the application.
package i18n

import (
	"sync"

	"github.com/tidwall/gjson"
	"golang.org/x/text/language"
)

var (
	translation gjson.Result

	mu sync.Mutex
)

// Init initializes internationalization with the language.
// It must be called at least once before `Get` is called.
// This method is concurrent safe.
func Init(lang string) {
	mu.Lock()
	defer mu.Unlock()

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

// SupportedLangs returns a map of currently supported languages.
// Keys of the returned map are the language codes, and values are
// human-readable names of the languages.
func SupportedLangs() map[string]string {
	return map[string]string{
		"en": "English",
		"ru": "Русский",
	}
}
