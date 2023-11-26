package editor

import (
	"strings"
	"testing"

	"github.com/fatih/color"
	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	cases := []string{"message", "1", "2", " ", "a b c"}

	for _, c := range cases {
		assert.Equal(t, color.New(color.Bold, color.FgYellow).Sprint(c), message(c))
	}
}

func TestSecondaryMessage(t *testing.T) {
	cases := []string{"secondary", " ", "1", "2", "a b c"}

	for _, c := range cases {
		assert.Equal(t, color.New(color.FgMagenta, color.Bold).Sprintf("- %s", c), secondaryMessage(c))
	}
}

func TestNestedMessage(t *testing.T) {
	cases := []struct {
		message string
		level   int
	}{
		{message: "string", level: 0},
		{message: "some", level: 1},
		{message: " ", level: 2},
		{message: "a b c d e", level: 3},
		{message: "message", level: 4},
		{message: "Enter filename", level: 5},
	}

	for _, c := range cases {
		assert.Equal(
			t,
			color.New(color.Bold).Sprintf("%s %s", strings.Repeat("-", c.level), c.message),
			nestedMessage(c.message, c.level),
		)
	}
}

func TestGetAttachmentSrc(t *testing.T) {
	cases := []struct {
		src   string
		isErr bool
	}{
		{src: "https://example.com/image.png", isErr: false},
		{src: "", isErr: true},
		{src: "invalid URL", isErr: true},
	}

	for _, c := range cases {
		_, err := getAttachmentSrc(c.src)

		assert.Equal(t, err != nil, c.isErr)
	}
}
