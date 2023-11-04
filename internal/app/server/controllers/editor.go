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
	parser "github.com/shelepuginivan/hakutest/internal/pkg/test_parser"
)

type EditorController struct{}

func (e EditorController) ChooseTest(c *gin.Context) {
	c.HTML(http.StatusOK, "editor_upload.tmpl", gin.H{
		"Config": config.Init(),
	})
}

func (e EditorController) NewTest(c *gin.Context) {
	c.HTML(http.StatusOK, "editor.tmpl", gin.H{
		"Config": config.Init(),
		"Test":   parser.Test{},
		"incr": func(n int) int {
			return n + 1
		},
	})
}

func (e EditorController) UploadTest(c *gin.Context) {
	test := parser.Test{}
	err := c.Request.ParseForm()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	uploadedFile, err := file.Open()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init().Ui.Error,
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
			"Config": config.Init().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to read uploaded file",
			"Error":  err.Error(),
		})

		return
	}

	err = json.Unmarshal(data, &test)

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse test",
			"Error":  err.Error(),
		})

		return
	}

	c.HTML(http.StatusOK, "editor.tmpl", gin.H{
		"Config": config.Init(),
		"Test":   test,
		"incr": func(n int) int {
			return n + 1
		},
	})
}

func (e EditorController) CreateTest(c *gin.Context) {
	test := parser.Test{}
	err := c.Request.ParseMultipartForm(1000)

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init().Ui.Error,
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	expiresIn, err := time.Parse("2006-01-02T15:04:05", c.Request.Form.Get("expiresIn"))

	if err == nil {
		test.ExpiresIn = expiresIn
	}

	test.Title = c.Request.Form.Get("title")
	test.Description = c.Request.Form.Get("description")
	test.Subject = c.Request.Form.Get("subject")
	test.Target = c.Request.Form.Get("target")
	test.Title = c.Request.Form.Get("title")
	test.Author = c.Request.Form.Get("author")
	test.Institution = c.Request.Form.Get("institution")

	numberOfTasks, err := strconv.Atoi(c.Request.Form.Get("number-of-tasks"))

	if err != nil {
		numberOfTasks = 0
	}

	for i := 0; i < numberOfTasks; i++ {
		task := parser.Task{}

		task.Type = c.Request.Form.Get(fmt.Sprintf("%d-type", i))
		task.Text = c.Request.Form.Get(fmt.Sprintf("%d-text", i))
		task.Answer = c.Request.Form.Get(fmt.Sprintf("%d-answer", i))
		task.Options = c.PostFormArray(fmt.Sprintf("%d-options", i))

		if c.Request.Form.Get(fmt.Sprintf("%d-has-attachment", i)) == "on" {
			attachment := parser.Attachment{}

			attachment.Type = c.Request.Form.Get(fmt.Sprintf("%d-attachment-type", i))
			attachment.Name = c.Request.Form.Get(fmt.Sprintf("%d-attachment-name", i))
			attachment.Src = c.Request.Form.Get(fmt.Sprintf("%d-attachment-src", i))

			task.Attachment = attachment
		}

		test.Tasks = append(test.Tasks, task)
	}

	data, err := json.Marshal(test)

	if err != nil {
		c.HTML(http.StatusBadRequest, "error.tmpl", gin.H{
			"Config": config.Init().Ui.Error,
			"Code":   http.StatusBadRequest,
			"Detail": "failed to create a test file",
			"Error":  err.Error(),
		})

		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s.json", test.Title))
	c.Header("Content-Type", "application/json")
	c.Status(http.StatusCreated)
	c.Writer.Write(data)
}
