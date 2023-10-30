package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	r := gin.Default()
	config := config.Init()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return r.Run(":" + config.Port)
}
