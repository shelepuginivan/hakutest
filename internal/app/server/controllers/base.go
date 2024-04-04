package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

type BaseController struct{}

func (co BaseController) SendErrorResponse(c *gin.Context, code int, err error, detail string) {
	c.HTML(code, "error.tmpl", gin.H{
		"Code":   code,
		"I18n":   i18n.New().Web.Error,
		"Detail": detail,
		"Error":  err.Error(),
	})
}
