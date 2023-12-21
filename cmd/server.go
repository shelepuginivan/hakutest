package cmd

import (
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [port]",
	Short: "Start Hakutest server",
	Long:  "Start Hakutest server",
	Args:  cobra.RangeArgs(0, 1),
	RunE:  server.Init,
}
