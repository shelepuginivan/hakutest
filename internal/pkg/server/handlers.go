package server

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
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

// registerStudentInterface adds endpoints for the student interface.
func registerStudentInterface(e *gin.Engine, cfg *config.Config) {
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Lang":  cfg.Lang,
			"I18n":  i18n.Get,
			"Tests": test.GetList(),
		})
	})

	e.GET("/:test", func(c *gin.Context) {
		name := c.Param("test")
		t, err := test.GetByName(name)

		if err != nil {
			c.String(http.StatusNotFound, "not found")
			return
		}

		c.HTML(http.StatusOK, "test.html", gin.H{
			"Lang":     cfg.Lang,
			"I18n":     i18n.Get,
			"Test":     t,
			"TestName": name,
		})
	})

	e.POST("/:test", func(c *gin.Context) {
		s := &test.Solution{
			SubmittedAt: time.Now(),
		}

		name := c.Param("test")

		t, err := test.GetByName(name)
		if err != nil {
			c.String(http.StatusNotFound, "not found")
			return
		}

		if err = c.Request.ParseForm(); err != nil {
			c.String(http.StatusUnprocessableEntity, "unprocessable entity")
			return
		}

		s.Student = c.PostForm("student")

		for i := range len(t.Tasks) {
			answer := c.PostFormArray(strconv.Itoa(i))
			answerString := strings.Join(answer, ",")

			s.Answers = append(s.Answers, answerString)
		}

		r := results.New(t, s)

		if err = results.Save(r, name); err != nil {
			c.String(http.StatusInternalServerError, "failed to save results")
			return
		}

		if cfg.ShowResults {
			c.HTML(http.StatusCreated, "result.html", gin.H{
				"Lang":   cfg.Lang,
				"I18n":   i18n.Get,
				"Result": r,
				"Show":   cfg.ShowResults,
			})
			return
		}

		c.HTML(http.StatusCreated, "info.html", gin.H{
			"Lang":  cfg.Lang,
			"I18n":  i18n.Get,
			"Title": i18n.Get("result.title"),
			"Text":  i18n.Get("result.disabled"),
		})
	})
}
