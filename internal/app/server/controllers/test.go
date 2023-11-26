package controllers

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/core"
)

type TestController struct{}

func (t TestController) GetTest(c *gin.Context) {
	testName := c.Param("test")
	test, err := core.GetTest(testName)

	if err != nil {
		code := http.StatusBadRequest
		detail := "failed to parse test file"

		if os.IsNotExist(err) {
			code = http.StatusNotFound
			detail = "test file does not exist"
		}

		c.HTML(code, "error.tmpl", gin.H{
			"Code":   code,
			"Config": config.Init().Ui.Error,
			"Detail": detail,
			"Error":  err.Error(),
		})

		return
	}

	if !test.ExpiresIn.IsZero() && test.ExpiresIn.Before(time.Now()) {
		c.HTML(http.StatusGone, "expired.tmpl", gin.H{
			"Config": config.Init().Ui.Expired,
		})

		return
	}

	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"Config": config.Init().Ui.Test,
		"Title":  test.Title,
		"Tasks":  test.Tasks,
		"url": func(s string) template.URL {
			return template.URL(s)
		},
		"incr": func(n int) int {
			return n + 1
		},
	})
}

func (t TestController) SubmitTest(c *gin.Context) {
	err := c.Request.ParseForm()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Code":   http.StatusUnprocessableEntity,
			"Config": config.Init().Ui.Error,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	name := c.Param("test")
	test, err := core.GetTest(name)

	if err != nil {
		code := http.StatusBadRequest
		detail := "failed to parse test file"

		if os.IsNotExist(err) {
			code = http.StatusNotFound
			detail = "test file does not exist"
		}

		c.HTML(code, "error.tmpl", gin.H{
			"Code":   code,
			"Config": config.Init().Ui.Error,
			"Detail": detail,
			"Error":  err.Error(),
		})

		return
	}

	if !test.ExpiresIn.IsZero() && test.ExpiresIn.Before(time.Now()) {
		c.HTML(http.StatusGone, "expired.tmpl", gin.H{
			"Config": config.Init().Ui.Expired,
		})

		return
	}

	results := test.GetResults(c.Request.PostForm)

	if err := results.Save(name); err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Code":   http.StatusBadRequest,
			"Config": config.Init().Ui.Error,
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
