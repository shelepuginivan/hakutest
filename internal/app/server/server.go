// Package server provides HTTP server for the app.
package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/app/server/controllers"
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
	s := controllers.NewStudent(cfg)
	student := engine.Group("/", security.Middleware(
		cfg.Security.Student,
		security.RoleStudent,
	))
	student.GET("/", s.SearchPage)
	student.GET("/:test", s.TestIsAvailable, s.TestPage)
	student.POST("/:test", s.TestIsAvailable, s.TestSubmission)

	// Teacher interface.
	t := controllers.NewTeacher(cfg)
	teacher := engine.Group("/teacher", security.Middleware(
		cfg.Security.Teacher,
		security.RoleTeacher,
	))
	teacher.GET("/", t.Index)
	teacher.GET("/dashboard", t.Dashboard)
	teacher.GET("/tests", t.Tests)
	teacher.GET("/tests/selected", t.DownloadSelected)
	teacher.POST("/tests/selected", t.DeleteSelected)
	teacher.POST("/tests/import", t.ImportTests)
	teacher.GET("/tests/action/:test", t.DownloadTest)
	teacher.POST("/tests/action/:test", t.DeleteTest)
	teacher.GET("/statistics", t.Statistics)
	teacher.GET("/statistics/export", t.StatisticsExport)
	teacher.GET("/settings", t.SettingsPage)
	teacher.POST("/settings", t.SettingsUpdate)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: engine,
	}
}
