package controllers

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

type TestController struct {
	s test.TestService
	r results.ResultsService
}

func NewTestController(s test.TestService, r results.ResultsService) TestController {
	return TestController{s: s, r: r}
}

func (co TestController) GetTest(c *gin.Context) {
	name := c.Param("test")
	t, err := co.s.GetByName(name)

	if err != nil {
		code := http.StatusBadRequest
		detail := "failed to parse test file"

		if os.IsNotExist(err) {
			code = http.StatusNotFound
			detail = "test file does not exist"
		}

		c.HTML(code, "error.tmpl", gin.H{
			"Code":   code,
			"Config": config.New().Ui.Error,
			"Detail": detail,
			"Error":  err.Error(),
		})

		return
	}

	if !t.ExpiresIn.IsZero() && t.ExpiresIn.Before(time.Now()) {
		c.HTML(http.StatusGone, "expired.tmpl", gin.H{
			"Config": config.New().Ui.Expired,
		})

		return
	}

	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"Config": config.New().Ui.Test,
		"Title":  t.Title,
		"Tasks":  t.Tasks,
		"url": func(s string) template.URL {
			return template.URL(s)
		},
		"incr": func(n int) int {
			return n + 1
		},
	})
}

func (co TestController) SubmitTest(c *gin.Context) {
	err := c.Request.ParseForm()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Code":   http.StatusUnprocessableEntity,
			"Config": config.New().Ui.Error,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	name := c.Param("test")
	t, err := co.s.GetByName(name)

	if err != nil {
		code := http.StatusBadRequest
		detail := "failed to parse test file"

		if os.IsNotExist(err) {
			code = http.StatusNotFound
			detail = "test file does not exist"
		}

		c.HTML(code, "error.tmpl", gin.H{
			"Code":   code,
			"Config": config.New().Ui.Error,
			"Detail": detail,
			"Error":  err.Error(),
		})

		return
	}

	if !t.ExpiresIn.IsZero() && t.ExpiresIn.Before(time.Now()) {
		c.HTML(http.StatusGone, "expired.tmpl", gin.H{
			"Config": config.New().Ui.Expired,
		})

		return
	}

	results := co.r.GetResults(t, c.Request.PostForm)

	if err := co.r.Save(results, name); err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Code":   http.StatusBadRequest,
			"Config": config.New().Ui.Error,
			"Detail": "failed to save test results",
			"Error":  err.Error(),
		})

		return
	}

	c.HTML(http.StatusCreated, "results.tmpl", gin.H{
		"Student": results.Student,
		"Results": results.Results,
	})
}
