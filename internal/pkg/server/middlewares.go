package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
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
		c.String(http.StatusNotFound, "not found")
		c.Abort()
		return
	}

	if t.IsExpired() {
		c.String(http.StatusGone, "expired")
		c.Abort()
		return
	}

	c.Next()
}
