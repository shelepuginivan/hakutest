package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/web"
)

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
