// Package controllers provides Gin controllers for the server.
package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

// BaseController is a base controller struct.
// It is used to promote common controller methods.
type BaseController struct {
	internationalization *i18n.I18n
}

// I18n returns global app internationalization.
// It saves i18n.I18n struct, so the file is only read once.
func (co BaseController) I18n() *i18n.I18n {
	if co.internationalization == nil {
		co.internationalization = i18n.New()
	}

	return co.internationalization
}

// SendErrorResponse renders error.tmpl template.
// It shows status code of the response, the occurred error, and details about this error.
func (co BaseController) SendErrorResponse(c *gin.Context, code int, err error, detail string) {
	c.HTML(code, "error.tmpl", gin.H{
		"Language": co.I18n().Language,
		"Code":     code,
		"I18n":     co.I18n().Web.Error,
		"Detail":   detail,
		"Error":    err.Error(),
	})
}
