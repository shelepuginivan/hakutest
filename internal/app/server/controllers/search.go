package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/core"
)

type SearchController struct{}

func (s SearchController) SearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"Config":   config.Init().Ui.Search,
		"TestList": core.GetTestList(),
	})
}