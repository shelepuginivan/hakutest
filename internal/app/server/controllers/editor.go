package controllers

import (
	"encoding/json"
	"io"
	"net/http"

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

func (e EditorController) UploadTest(c *gin.Context) {
	test := parser.Test{}
	err := c.Request.ParseForm()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init(),
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	file, err := c.FormFile("file")

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init(),
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse form",
			"Error":  err.Error(),
		})

		return
	}

	uploadedFile, err := file.Open()

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init(),
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
			"Config": config.Init(),
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to read uploaded file",
			"Error":  err.Error(),
		})

		return
	}

	err = json.Unmarshal(data, &test)

	if err != nil {
		c.HTML(http.StatusUnprocessableEntity, "error.tmpl", gin.H{
			"Config": config.Init(),
			"Code":   http.StatusUnprocessableEntity,
			"Detail": "failed to parse test",
			"Error":  err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, test)
}
