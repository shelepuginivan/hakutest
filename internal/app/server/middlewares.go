package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/pkg/test"
)

// Logger is a custom logging middleware using [github.com/rs/zerolog].
// It logs method, path, status code, latency, and errors (if any).
func Logger(c *gin.Context) {
	t := time.Now()
	c.Next()

	var level zerolog.Level
	switch {
	case c.Writer.Status() >= 500:
		level = zerolog.ErrorLevel
	case c.Writer.Status() >= 400:
		level = zerolog.WarnLevel
	default:
		level = zerolog.InfoLevel
	}

	msg := c.Errors.String()
	if msg == "" {
		msg = "No errors"
	}

	log.WithLevel(level).
		Str("method", c.Request.Method).
		Str("path", c.Request.RequestURI).
		Int("status", c.Writer.Status()).
		Dur("latency", time.Since(t)).
		Msg(msg)
}

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
		c.HTML(http.StatusNotFound, "info.gohtml", gin.H{
			"Lang":  i18n.Lang(),
			"I18n":  i18n.Get,
			"Title": i18n.Get("test_not_found.title"),
			"Text":  i18n.Get("test_not_found.text"),
		})
		c.Abort()
		return
	}

	if t.IsExpired() {
		c.HTML(http.StatusGone, "info.gohtml", gin.H{
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
