// Package markdown provides encapsulated methods to convert Markdown to HTML.
package markdown

import (
	"html/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
	"github.com/microcosm-cc/bluemonday"
)

const (
	parserExt = parser.CommonExtensions | parser.OrderedListStart | parser.SuperSubscript
	htmlFlags = html.CommonFlags | html.HrefTargetBlank
)

var (
	rendererOptions = html.RendererOptions{Flags: htmlFlags}
)

// ToHTML converts Markdown to HTML and sanitizes the resulting HTML.
func ToHTML(md string) string {
	p := parser.NewWithExtensions(parserExt)
	document := p.Parse([]byte(md))
	r := html.NewRenderer(rendererOptions)
	hypertext := markdown.Render(document, r)

	sanitized := bluemonday.UGCPolicy().SanitizeBytes(hypertext)

	return string(sanitized)
}

// ToGoHTML converts Markdown to Go template.HTML.
// It runs ToHTML internally.
func ToGoHTML(md string) template.HTML {
	return template.HTML(ToHTML(md))
}
