// Package server provides HTTP server for the app.
package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/logging"
)

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

	engine.Use(gin.Recovery())
	logging.RegisterHttp(engine)
	registerStatic(engine)
	registerTemplates(engine)

	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Lang": cfg.Lang,
			"I18n": i18n.Get,
		})
	})

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: engine,
	}
}
