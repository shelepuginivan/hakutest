// Package server provides HTTP server for the app.
package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"github.com/shelepuginivan/hakutest/web"
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

	engine.Use(gin.Recovery())
	engine.Use(Logger)
	engine.Use(RequestTimestamp)
	engine.SetHTMLTemplate(templates())
	engine.StaticFS("/static", http.FS(web.Static))

	security.Register(engine, cfg.Security.Student, cfg.Security.Teacher)

	// Student interface.
	student := engine.Group("/", security.Middleware(
		cfg.Security.Student,
		security.RoleStudent,
	))
	registerStudentInterface(student, cfg)

	// Teacher interface.
	teacher := engine.Group("/teacher", security.Middleware(
		cfg.Security.Teacher,
		security.RoleTeacher,
	))
	registerTeacherInterface(teacher, cfg)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: engine,
	}
}
