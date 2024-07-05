// Package server provides HTTP server for the app.
package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/embedded"
	"github.com/shelepuginivan/hakutest/internal/pkg/logging"
	"github.com/shelepuginivan/hakutest/pkg/security"
)

// setMode sets mode of the gin engine.
// By default, mode is set to `"release"`.
func setMode(cfg *config.Config) {
	if cfg.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
}

// New returns a new instance of HTTP server.
func New(cfg *config.Config) *http.Server {
	setMode(cfg)

	engine := gin.New()

	engine.Use(gin.LoggerWithConfig(logging.HttpConfig()))
	engine.Use(RequestTimestamp)
	engine.Use(gin.Recovery())
	engine.Use(serveFavicon(embedded.Icon))

	registerStatic(engine)
	registerTemplates(engine)
	security.Register(engine, cfg.Security.Student, cfg.Security.Teacher)
	registerStudentInterface(engine, cfg)
	registerTeacherInterface(engine, cfg)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: engine,
	}
}
