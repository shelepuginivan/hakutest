package i18n

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	i18nEn := New(LanguageEn)

	// Should fallback to English by default.
	assert.EqualValues(t, i18nEn, New(""))

	for _, l := range AvailableLanguages {
		if l != LanguageEn {
			assert.NotEqualValues(t, i18nEn, New(l))
		}
	}
}
