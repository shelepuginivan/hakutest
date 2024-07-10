package server

import (
	"html/template"

	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/markdown"
	"github.com/shelepuginivan/hakutest/web"
)

// templates provides HTML templates defined in `/web/templates` directory.
// It parses the embedded filesystem and adds custom functions.
func templates() *template.Template {
	tmpl := template.New("embedded")

	tmpl = tmpl.Funcs(template.FuncMap{
		"i": i18n.Get,
		"incr": func(i int) int {
			return i + 1
		},
		"markdown": markdown.ToGoHTML,
		"iter": func(i int) (stream chan int) {
			stream = make(chan int)
			go func() {
				for k := range i {
					stream <- k
				}
				close(stream)
			}()
			return
		},
	})

	return template.Must(tmpl.ParseFS(
		web.Templates,
		"templates/*.gohtml",
		"partials/*.gohtml",
	))
}
