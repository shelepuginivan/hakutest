package controllers

import (
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	parser "github.com/shelepuginivan/hakutest/internal/pkg/test_parser"
)

type TestController struct{}

func (t TestController) GetTest(c *gin.Context) {
	testName := c.Param("test")
	test, err := parser.ParseTest(testName)

	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"detail": "test file does not exist"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"detail": "failed to parse test file"})
		}

		return
	}

	if !test.ExpiresIn.IsZero() && test.ExpiresIn.Before(time.Now()) {
		c.JSON(http.StatusGone, gin.H{"detail": "Test expired"})
		return
	}

	c.HTML(http.StatusOK, "test.tmpl", gin.H{
		"Title": test.Title,
		"Tasks": test.Tasks,
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
		c.JSON(http.StatusUnprocessableEntity, gin.H{"detail": "failed to parse form"})
	}

	results, err := parser.GetTestResults(c.Param("test"), c.Request.PostForm)

	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"detail": "test file does not exist"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"detail": "failed to parse test file"})
		}

		return
	}

	c.JSON(http.StatusOK, results)
}
