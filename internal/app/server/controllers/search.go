package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// SearchController is a controller that handles `/` requests.
type SearchController struct {
	BaseController
	s *test.TestService
}

// NewSearchController returns a new instance of SearchController.
func NewSearchController(app *application.App, s *test.TestService) *SearchController {
	return &SearchController{
		s:              s,
		BaseController: BaseController{app: app},
	}
}

// SearchPage handles `GET /` requests.
// It renders the test search page.
func (co SearchController) SearchPage(c *gin.Context) {
	c.HTML(http.StatusOK, "search.tmpl", gin.H{
		"Language": co.app.I18n.Language,
		"I18n":     co.app.I18n.Web.Search,
		"TestList": co.s.GetTestList(),
	})
}
