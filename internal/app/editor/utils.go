package editor

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gabriel-vasile/mimetype"
)

func getAttachmentSrc(src string) (string, error) {
	if bytes, err := os.ReadFile(src); err == nil {
		mimeType := mimetype.Detect(bytes)
		base64Endoding := base64.StdEncoding.EncodeToString(bytes)

		return fmt.Sprintf("data:%s;base64,%s", mimeType, base64Endoding), nil
	}

	_, err := url.ParseRequestURI(src)

	if err == nil {
		return src, nil
	}

	return "", err
}

func message(s string) string {
	return color.New(color.Bold, color.FgYellow).Sprint(s)
}

func secondaryMessage(s string) string {
	return color.New(color.FgMagenta, color.Bold).Sprintf("- %s", s)
}

func nestedMessage(s string, level int) string {
	return color.New(color.Bold).Sprintf("%s %s", strings.Repeat("-", level), s)
}
