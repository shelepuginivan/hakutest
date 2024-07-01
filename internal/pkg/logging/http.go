package logging

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func httpFormatter(param gin.LogFormatterParams) string {
	return fmt.Sprintf("HTTP | %s %s %s - %d (%s) | [%s]\n",
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.TimeStamp.Format(time.RFC1123),
	)
}

// RegisterHttp registers logger for Gin HTTP Engine.
func RegisterHttp(e *gin.Engine) {
	cfg := gin.LoggerConfig{
		Output:    Output(),
		Formatter: httpFormatter,
	}

	e.Use(gin.LoggerWithConfig(cfg))
}
