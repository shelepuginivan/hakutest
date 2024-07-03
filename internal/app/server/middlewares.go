package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/pkg/test"
)

// RequestTimestamp is a middleware that sets timestamp of the request to the
// context.
func RequestTimestamp(c *gin.Context) {
	c.Set("timestamp", time.Now())
	c.Next()
}

// TestIsAvailable is a middleware function that performs various checks:
// - Test exists;
// - Test is not expired.
func TestIsAvailable(c *gin.Context) {
	testName := c.Param("test")

	t, err := test.GetByName(testName)
	if err != nil {
		c.HTML(http.StatusNotFound, "info.html", gin.H{
			"Lang":  i18n.Lang(),
			"I18n":  i18n.Get,
			"Title": i18n.Get("test_not_found.title"),
			"Text":  i18n.Get("test_not_found.text"),
		})
		c.Abort()
		return
	}

	if t.IsExpired() {
		c.HTML(http.StatusGone, "info.html", gin.H{
			"Lang":  i18n.Lang(),
			"I18n":  i18n.Get,
			"Title": i18n.Get("expired.title"),
			"Text":  i18n.Get("expired.text"),
		})
		c.Abort()
		return
	}

	c.Next()
}
