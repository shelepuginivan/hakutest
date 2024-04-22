// Package controllers provides Gin controllers for the server.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
)

// BaseController is a base controller struct.
// It is used to promote common controller methods.
type BaseController struct {
	app *application.App
}

// SendErrorResponse renders error.tmpl template.
// It shows status code of the response, the occurred error, and details about this error.
func (co BaseController) SendErrorResponse(c *gin.Context, code int, err error, detail string) {
	c.HTML(code, "error.tmpl", gin.H{
		"Language": co.app.I18n.Language,
		"Code":     code,
		"I18n":     co.app.I18n.Web.Error,
		"Detail":   detail,
		"Error":    err.Error(),
	})
}
