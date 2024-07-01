package i18n

import _ "embed"

//go:embed translations/en.json
var en string

//go:embed translations/ru.json
var ru string

var translations = map[string]string{
	"en": en,
	"ru": ru,
}
