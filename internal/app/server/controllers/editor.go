package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

type EditorController struct{}

func NewEditorController() EditorController {
	return EditorController{}
}

func (co EditorController) ChooseTest(c *gin.Context) {
	c.HTML(http.StatusOK, "editor_upload.tmpl", gin.H{
		"Config": config.New().Ui.Editor,
	})
}

func (co EditorController) NewTest(c *gin.Context) {
	c.HTML(http.StatusOK, "editor.tmpl", gin.H{
		"Config": config.New().Ui.Editor,
		"Test":   test.Test{},
		"incr": func(n int) int {
			return n + 1
		},
	})
}

func (co EditorController) UploadTest(c *gin.Context) {
	t := test.Test{}
	err := c.Request.ParseForm()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	uploadedFile, err := file.Open()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to open uploaded file",
			"Error":  err.Error(),
		})

		return
	}

	defer uploadedFile.Close()

	data, err := io.ReadAll(uploadedFile)

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to read uploaded file",
			"Error":  err.Error(),
		})

		return
	}

	err = json.Unmarshal(data, &t)

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse test",
			"Error":  err.Error(),
		})

		return
	}

	c.HTML(http.StatusOK, "editor.tmpl", gin.H{
		"Config": config.New().Ui.Editor,
		"Test":   t,
		"incr": func(n int) int {
			return n + 1
		},
	})
}

func (co EditorController) CreateTest(c *gin.Context) {
	t := test.Test{}
	err := c.Request.ParseMultipartForm(1000)

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	expiresIn, err := time.Parse("2006-01-02T15:04:05", c.Request.Form.Get("expiresIn"))

	if err == nil {
		t.ExpiresIn = expiresIn
	}

	t.Title = c.Request.Form.Get("title")
	t.Description = c.Request.Form.Get("description")
	t.Subject = c.Request.Form.Get("subject")
	t.Target = c.Request.Form.Get("target")
	t.Title = c.Request.Form.Get("title")
	t.Author = c.Request.Form.Get("author")
	t.Institution = c.Request.Form.Get("institution")

	numberOfTasks, err := strconv.Atoi(c.Request.Form.Get("number-of-tasks"))

	if err != nil {
		numberOfTasks = 0
	}

	for i := 0; i < numberOfTasks; i++ {
		task := test.Task{}

		task.Type = c.Request.Form.Get(fmt.Sprintf("%d-type", i))
		task.Text = c.Request.Form.Get(fmt.Sprintf("%d-text", i))
		task.Answer = c.Request.Form.Get(fmt.Sprintf("%d-answer", i))
		task.Options = c.PostFormArray(fmt.Sprintf("%d-options", i))

		if c.Request.Form.Get(fmt.Sprintf("%d-has-attachment", i)) == "on" {
			attachment := test.Attachment{}

			attachment.Type = c.Request.Form.Get(fmt.Sprintf("%d-attachment-type", i))
			attachment.Name = c.Request.Form.Get(fmt.Sprintf("%d-attachment-name", i))
			attachment.Src = c.Request.Form.Get(fmt.Sprintf("%d-attachment-src", i))

			task.Attachment = attachment
		}

		t.Tasks = append(t.Tasks, task)
	}

	data, err := json.Marshal(t)

	if err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusBadRequest,
			"Detail": "failed to create a test file",
			"Error":  err.Error(),
		})

		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.json", t.Title))
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)

	if _, err := c.Writer.Write(data); err != nil {
		c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
			"Config": config.New().Ui.Error,
			"Code":   http.StatusInternalServerError,
			"Detail": "failed to write response data",
			"Error":  err.Error(),
		})
	}
}
