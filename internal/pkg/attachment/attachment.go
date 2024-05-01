package attachment

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/url"
	"os"
	"path/filepath"
)

// Attachment represents an attachment of the Task with a name, type, and source.
type Attachment struct {
	Name string `json:"name"` // Name of the attachment.
	Type string `json:"type"` // Type of the attachment.
	Src  string `json:"src"`  // Source of the attachment, URL or base64 URL.
}

// IsURL reports whether src is a valid URL and not a filepath.
func IsURL(src string) bool {
	u, err := url.Parse(src)
	return err == nil && u.Scheme != "" && u.Host != ""
}

// New returns a new instance of Attachment.
// If src is an URL, it sets it as the source of the Attachment.
// If src is a path to a file, it encodes the file to base64 URL.
func New(n, t, src string) *Attachment {
	bytes, err := os.ReadFile(src)
	if err != nil {
		return &Attachment{
			Name: n,
			Type: t,
			Src:  src,
		}
	}

	mimeType := mime.TypeByExtension(filepath.Ext(src))
	base64Endoding := base64.StdEncoding.EncodeToString(bytes)
	base64URL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Endoding)

	return &Attachment{
		Name: n,
		Type: t,
		Src:  base64URL,
	}
}
