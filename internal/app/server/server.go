package server

import (
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	r := NewRouter()
	port := config.New().Server.Port

	if len(args) == 1 {
		port = args[0]
	}

	return r.Run(":" + port)
}
