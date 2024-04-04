package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

type BaseController struct {
	internationalization *i18n.I18n
}

func (co BaseController) I18n() *i18n.I18n {
	if co.internationalization == nil {
		co.internationalization = i18n.New()
	}

	return co.internationalization
}

func (co BaseController) SendErrorResponse(c *gin.Context, code int, err error, detail string) {
	c.HTML(code, "error.tmpl", gin.H{
		"Language": co.I18n().Language,
		"Code":     code,
		"I18n":     co.I18n().Web.Error,
		"Detail":   detail,
		"Error":    err.Error(),
	})
}
