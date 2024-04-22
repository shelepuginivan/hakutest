package i18n

// AvailableLanguages is a slice of available languages.
var AvailableLanguages = []string{
	LanguageEn,
	LanguageRu,
}

// LanguageCodeMap is a mapping between language codes and language names.
// E.g. LanguageCodeMap["en"] is "English".
var LanguageCodeMap = map[string]string{
	LanguageEn: "English",
	LanguageRu: "Русский",
}
