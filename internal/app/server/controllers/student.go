package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/test"
)

// StudentController is a [github.com/gin-gonic/gin] controller that provides
// handlers for student routes.
type StudentController struct {
	cfg *config.Config
}

// NewStudent returns a new instance of [StudentController].
func NewStudent(cfg *config.Config) *StudentController {
	return &StudentController{cfg: cfg}
}

// SearchPage is a [github.com/gin-gonic/gin] handler for the `GET /` route.
// It renders HTML page with test search.
func (co *StudentController) SearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.gohtml", gin.H{
		"Tests": test.GetList(),
	})
}

// TestIsAvailable is a [github.com/gin-gonic/gin] middleware. It checks
// whether requested test exists and is not expired. If conditions are
// satisfied, it sets context parameter `test` containing a pointer to the
// resolved test.

// TestIsAvailable is intended to be used with handlers for routes that contain
// `:test` parameter.
func (co *StudentController) TestIsAvailable(c *gin.Context) {
	testName := c.Param("test")

	t, err := test.GetByName(testName)
	if err != nil {
		c.HTML(http.StatusNotFound, "info.gohtml", gin.H{
			"Title": i18n.Get("test_not_found.title"),
			"Text":  i18n.Get("test_not_found.text"),
		})
		c.Abort()
		return
	}

	if t.IsExpired() {
		c.HTML(http.StatusGone, "info.gohtml", gin.H{
			"Title": i18n.Get("expired.title"),
			"Text":  i18n.Get("expired.text"),
		})
		c.Abort()
		return
	}

	c.Set("test", t)
	c.Next()
}

// SearchPage is a [github.com/gin-gonic/gin] handler for the `GET /:test`
// route. It renders HTML page for the specified test.
//
// TestPage must be used after [StudentController.TestIsAvailable] middleware.
func (co *StudentController) TestPage(c *gin.Context) {
	t, _ := c.Get("test")

	c.HTML(http.StatusOK, "test.gohtml", gin.H{
		"Test": t,
	})
}

// TestSubmission is a [github.com/gin-gonic/gin] handler for the `POST /:test`
// route. It checks answers given by the student via form and saves the result.
//
// TestSubmission must be used after [StudentController.TestIsAvailable]
// middleware.
func (co *StudentController) TestSubmission(c *gin.Context) {
	testName := c.Param("test")
	te, _ := c.Get("test")
	t := te.(*test.Test)

	if err := c.Request.ParseForm(); err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
			"Title":   i18n.Get("err.unprocessable.title"),
			"Text":    i18n.Get("err.unprocessable.text"),
			"Code":    http.StatusUnprocessableEntity,
			"Message": "failed to parse answers",
			"Error":   err.Error(),
		})
		return
	}

	s := &test.Solution{
		Student:     c.PostForm("student"),
		SubmittedAt: c.GetTime("timestamp"),
	}

	for i := range t.Tasks {
		answer := c.PostFormArray(strconv.Itoa(i))
		answerString := strings.Join(answer, ",")

		s.Answers = append(s.Answers, answerString)
	}

	r := result.New(t, s)

	if err := result.Save(r, testName); err != nil {
		c.HTML(http.StatusConflict, "error.gohtml", gin.H{
			"Title":   i18n.Get("err.write.title"),
			"Text":    i18n.Get("err.write.text"),
			"Code":    http.StatusConflict,
			"Message": "failed to save answers",
			"Error":   err.Error(),
		})
		return
	}

	if co.cfg.Result.Show {
		c.HTML(http.StatusCreated, "result.gohtml", gin.H{
			"Result": r,
		})
		return
	}

	c.HTML(http.StatusCreated, "info.gohtml", gin.H{
		"Title": i18n.Get("result.title"),
		"Text":  i18n.Get("result.disabled"),
	})
}
