package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/markdown"
	"github.com/shelepuginivan/hakutest/web"
)

// serveFavicon adds routes for serving favicon.
func serveFavicon(favicon []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.RequestURI != "/favicon.ico" {
			return
		}

		if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodHead {
			status := http.StatusOK
			if c.Request.Method != http.MethodOptions {
				status = http.StatusMethodNotAllowed
			}

			c.Header("Allow", "GET,HEAD,OPTIONS")
			c.AbortWithStatus(status)
			return
		}

		c.Data(http.StatusOK, "image/x-icon", favicon)
	}
}

// registerStatic sets engine static file system.
func registerStatic(e *gin.Engine) {
	staticFS := http.FS(web.Static)
	e.StaticFS("/static", staticFS)
}

// registerTemplates sets engine template file system.
func registerTemplates(e *gin.Engine) {
	var err error
	tmpl := template.New("embedded")

	tmpl = tmpl.Funcs(template.FuncMap{
		"i": i18n.Get,
		"incr": func(i int) int {
			return i + 1
		},
		"lang":     i18n.Lang,
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

	tmpl, err = tmpl.ParseFS(
		web.Templates,
		"templates/*.gohtml",
		"partials/*.gohtml",
	)
	if err != nil {
		panic(err)
	}

	e.SetHTMLTemplate(tmpl)
}
