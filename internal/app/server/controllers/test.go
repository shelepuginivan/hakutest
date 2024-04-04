package controllers

import (
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

type TestController struct {
	BaseController
	s *test.TestService
	r *results.ResultsService
}

func NewTestController(s *test.TestService, r *results.ResultsService) *TestController {
	return &TestController{s: s, r: r}
}

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
			"Language": co.I18n().Language,
			"I18n":     co.I18n().Web.Expired,
		})

		return
	}

	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"Language": co.I18n().Language,
		"I18n":     co.I18n().Web.Test,
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

func (co TestController) SubmitTest(c *gin.Context) {
	maxUploadSize := config.New().Server.MaxUploadSize

	// Limit the size of the submitted test form.
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxUploadSize)

	if err := c.Request.ParseMultipartForm(maxUploadSize); err != nil {
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
			"Language": co.I18n().Language,
			"I18n":     co.I18n().Web.Expired,
		})

		return
	}

	results := co.r.CheckAnswersWithFiles(name, t, c.Request.MultipartForm.Value, c.Request.MultipartForm.File)

	if err := co.r.Save(results, name); err != nil {
		co.SendErrorResponse(c, http.StatusBadRequest, err, "failed to save test results")
		return
	}

	if config.New().General.ShowResults {
		c.HTML(http.StatusCreated, "results.tmpl", gin.H{
			"Language": co.I18n().Language,
			"Student":  results.Student,
			"Results":  results.Results,
		})
		return
	}

	c.HTML(http.StatusCreated, "submitted.tmpl", gin.H{
		"Language": co.I18n().Language,
		"I18n":     co.I18n().Web.Submitted,
	})
}
