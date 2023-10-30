package server

import (
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	r := NewRouter()
	config := config.Init()
	return r.Run(":" + config.Port)
}
