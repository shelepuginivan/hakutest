package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/config"
)

type EditorController struct{}

func (e EditorController) ChooseTest(c *gin.Context) {
	c.HTML(http.StatusOK, "editor_upload.tmpl", gin.H{
		"Config": config.Init(),
	})
}
