package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/network"
	"github.com/shelepuginivan/hakutest/internal/pkg/uptime"
	"github.com/shelepuginivan/hakutest/pkg/result"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"github.com/shelepuginivan/hakutest/pkg/version"
)

// registerTeacherInterface adds endpoints for the teacher interface.
func registerTeacherInterface(e *gin.Engine, cfg *config.Config) {
	teacher := e.Group("/teacher")

	teacher.Use(security.Middleware(
		cfg.Security.Teacher,
		security.RoleTeacher,
	))

	teacher.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/teacher/dashboard")
	})

	teacher.GET("/dashboard", func(c *gin.Context) {
		localIP, err := network.GetLocalIP()
		if err == nil {
			localIP = fmt.Sprintf("http://%s:%d", localIP, cfg.Port)
		}

		u := uptime.Uptime()

		c.HTML(http.StatusOK, "dashboard.gohtml", gin.H{
			"AddressAvailable": err == nil,
			"Address":          localIP,
			"Version":          version.Version,
			"Uptime": gin.H{
				"Hours":   int(u.Hours()),
				"Minutes": int(u.Minutes()) % 60,
				"Seconds": int(u.Seconds()) % 60,
			},
		})
	})

	teacher.GET("/statistics", func(c *gin.Context) {
		resultName, ok := c.GetQuery("q")
		if ok {
			c.String(http.StatusOK, resultName)
			return
		}

		c.HTML(http.StatusOK, "statistics_menu.gohtml", gin.H{
			"AvailableResults": result.AvailableResults(),
		})
	})

	teacher.GET("/settings", func(c *gin.Context) {
		c.HTML(http.StatusOK, "settings.gohtml", gin.H{
			"Config":         cfg,
			"SupportedLangs": i18n.SupportedLangs(),
		})
	})

	teacher.POST("/settings", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
				"Title":   i18n.Get("settings.unprocessable.title"),
				"Text":    i18n.Get("settings.unprocessable.text"),
				"Code":    http.StatusUnprocessableEntity,
				"Message": "failed to parse form",
				"Error":   err.Error(),
			})
			return
		}

		fields := cfg.Fields

		fields.Debug = c.PostForm("debug") == "on"
		fields.DisableTray = c.PostForm("disable_tray") == "on"
		fields.Lang = c.PostForm("lang")
		fields.OverwriteResults = c.PostForm("overwrite_results") == "on"
		fields.ShowResults = c.PostForm("show_results") == "on"
		fields.TestsDirectory = c.PostForm("tests_directory")
		fields.ResultsDirectory = c.PostForm("results_directory")

		port, err := strconv.Atoi(c.PostForm("port"))
		if err == nil {
			fields.Port = port
		}

		err = cfg.Update(fields)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "error.gohtml", gin.H{
				"Title":   i18n.Get("settings.save_failed.title"),
				"Text":    i18n.Get("settings.save_failed.text"),
				"Code":    http.StatusUnprocessableEntity,
				"Message": "failed to write config file",
				"Error":   err.Error(),
			})
			return
		}

		c.HTML(http.StatusCreated, "settings.gohtml", gin.H{
			"Config":         cfg,
			"SupportedLangs": i18n.SupportedLangs(),
		})
	})
}
