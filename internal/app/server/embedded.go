package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
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
	templatesFS := template.Must(template.ParseFS(web.Templates, "templates/*.html"))
	e.SetHTMLTemplate(templatesFS)
}
