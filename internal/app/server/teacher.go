package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/network"
	"github.com/shelepuginivan/hakutest/internal/pkg/uptime"
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
}
