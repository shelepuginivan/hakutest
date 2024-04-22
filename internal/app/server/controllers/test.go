package controllers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// TestController is a controller that handles `/:test` requests.
type TestController struct {
	BaseController
	s *test.TestService
	r *results.ResultsService
}

// NewTestController returns a new instance of TestController.
func NewTestController(app *application.App, s *test.TestService, r *results.ResultsService) *TestController {
	return &TestController{
		s:              s,
		r:              r,
		BaseController: BaseController{app: app},
	}
}

// GetTest handles `GET /:test` requests.
// It renders a page of the respective test.
func (co TestController) GetTest(c *gin.Context) {
	name := c.Param("test")
	t, err := co.s.GetTestByName(name)

	if err != nil {
		code := http.StatusBadRequest
		detail := "failed to parse test file"

		if os.IsNotExist(err) {
			code = http.StatusNotFound
			detail = "test file does not exist"
		}

		co.SendErrorResponse(c, code, err, detail)
		return
	}

	if t.IsExpired() {
		c.HTML(http.StatusGone, "expired.tmpl", gin.H{
			"Language": co.app.I18n.Language,
			"I18n":     co.app.I18n.Web.Expired,
		})

		return
	}

	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"Language": co.app.I18n.Language,
		"I18n":     co.app.I18n.Web.Test,
		"Title":    t.Title,
		"Tasks":    t.Tasks,
		"url": func(s string) template.URL {
			return template.URL(s)
		},
		"incr": func(n int) int {
			return n + 1
		},
	})
}

// SubmitTest handles `POST /:test` requests.
// It performs a check of the submitted solution.
// Depending on the configuration, it renders either the result or the
// submission page.
func (co TestController) SubmitTest(c *gin.Context) {
	// Limit the size of the submitted test form.
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, co.app.Config.Server.MaxUploadSize)

	if err := c.Request.ParseMultipartForm(co.app.Config.Server.MaxUploadSize); err != nil {
		co.SendErrorResponse(c, http.StatusUnprocessableEntity, err, "failed to parse form")
		return
	}

	name := c.Param("test")
	t, err := co.s.GetTestByName(name)

	if err != nil {
		code := http.StatusBadRequest
		detail := "failed to parse test file"

		if os.IsNotExist(err) {
			code = http.StatusNotFound
			detail = "test file does not exist"
		}

		co.SendErrorResponse(c, code, err, detail)
		return
	}

	if t.IsExpired() {
		c.HTML(http.StatusGone, "expired.tmpl", gin.H{
			"Language": co.app.I18n.Language,
			"I18n":     co.app.I18n.Web.Expired,
		})

		return
	}

	results := co.r.CheckAnswersWithFiles(name, t, c.Request.MultipartForm.Value, c.Request.MultipartForm.File)

	if err := co.r.Save(results, name); err != nil {
		co.SendErrorResponse(c, http.StatusBadRequest, err, "failed to save test results")
		return
	}

	if co.app.Config.General.ShowResults {
		c.HTML(http.StatusCreated, "results.tmpl", gin.H{
			"Language": co.app.I18n.Language,
			"Student":  results.Student,
			"Results":  results.Results,
		})
		return
	}

	c.HTML(http.StatusCreated, "submitted.tmpl", gin.H{
		"Language": co.app.I18n.Language,
		"I18n":     co.app.I18n.Web.Submitted,
	})
}
