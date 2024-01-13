package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

type SearchController struct {
	s test.TestService
}

func NewSearchController(s test.TestService) SearchController {
	return SearchController{s: s}
}

func (co SearchController) SearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"Config":   config.New().Ui.Search,
		"TestList": co.s.GetTestList(),
	})
}
