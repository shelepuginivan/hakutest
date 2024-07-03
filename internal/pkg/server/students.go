package server

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// registerStudentInterface adds endpoints for the student interface.
func registerStudentInterface(e *gin.Engine, cfg *config.Config) {
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Lang":  i18n.Lang(),
			"I18n":  i18n.Get,
			"Tests": test.GetList(),
		})
	})

	e.GET("/:test", TestIsAvailable, func(c *gin.Context) {
		name := c.Param("test")
		t, _ := test.GetByName(name)

		c.HTML(http.StatusOK, "test.html", gin.H{
			"Lang":     i18n.Lang(),
			"I18n":     i18n.Get,
			"Test":     t,
			"TestName": name,
		})
	})

	e.POST("/:test", TestIsAvailable, func(c *gin.Context) {
		name := c.Param("test")

		t, _ := test.GetByName(name)

		if err := c.Request.ParseForm(); err != nil {
			c.HTML(http.StatusUnprocessableEntity, "error.html", gin.H{
				"Lang":    i18n.Lang(),
				"I18n":    i18n.Get,
				"Title":   i18n.Get("submission.unprocessable.title"),
				"Text":    i18n.Get("submission.unprocessable.text"),
				"Code":    http.StatusUnprocessableEntity,
				"Message": "failed to parse form",
				"Error":   err.Error(),
			})
			return
		}

		s := &test.Solution{
			Student:     c.PostForm("student"),
			SubmittedAt: c.GetTime("timestamp"),
		}

		for i := range len(t.Tasks) {
			answer := c.PostFormArray(strconv.Itoa(i))
			answerString := strings.Join(answer, ",")

			s.Answers = append(s.Answers, answerString)
		}

		r := results.New(t, s)

		if err := results.Save(r, name); err != nil {
			c.HTML(http.StatusConflict, "error.html", gin.H{
				"Lang":    i18n.Lang(),
				"I18n":    i18n.Get,
				"Title":   i18n.Get("submission.save_failed.title"),
				"Text":    i18n.Get("submission.save_failed.text"),
				"Code":    http.StatusConflict,
				"Message": "failed to save answers",
				"Error":   err.Error(),
			})
			return
		}

		if cfg.ShowResults {
			c.HTML(http.StatusCreated, "result.html", gin.H{
				"Lang":   i18n.Lang(),
				"I18n":   i18n.Get,
				"Result": r,
				"Incr": func(i int) int {
					return i + 1
				},
			})
			return
		}

		c.HTML(http.StatusCreated, "info.html", gin.H{
			"Lang":  i18n.Lang(),
			"I18n":  i18n.Get,
			"Title": i18n.Get("result.title"),
			"Text":  i18n.Get("result.disabled"),
		})
	})
}
