// Package logging provides internal methods for logging.
package logging

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func httpFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("HTTP | %s %s %s - %d (%s) | %s\n",
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.TimeStamp.Format(time.RFC1123),
	)
}

// HttpConfig returns configuration for gin.Logger.
func HttpConfig() gin.LoggerConfig {
	return gin.LoggerConfig{
		Output:    Output,
		Formatter: httpFormatter,
	}
}
