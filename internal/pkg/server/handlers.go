package server

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/shelepuginivan/hakutest/web"
)

func registerStatic(e *gin.Engine) {
	staticFS := http.FS(web.Static)
	e.StaticFS("/static", staticFS)
}

func registerTemplates(e *gin.Engine) {
	templatesFS := template.Must(template.ParseFS(web.Templates, "templates/*.html"))
	e.SetHTMLTemplate(templatesFS)
}

// registerStudentInterface adds endpoints for the student interface.
func registerStudentInterface(e *gin.Engine, cfg *config.Config) {
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Lang":  cfg.Lang,
			"I18n":  i18n.Get,
			"Tests": test.GetList(),
		})
	})
}
