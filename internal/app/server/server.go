package server

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	r := NewRouter()
	port := config.New().Server.Port
	addr := fmt.Sprintf(":%d", port)

	return r.Run(addr)
}
