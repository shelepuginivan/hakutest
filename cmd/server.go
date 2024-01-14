package cmd

import (
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/shelepuginivan/hakutest/internal/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func init() {
	port := config.New().Server.Port
	serverCmd.Flags().IntP("port", "p", port, "The port on which the server will be started")
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server [port]",
	Short: "Start Hakutest server",
	Long:  "Start Hakutest server",
	Args:  cobra.NoArgs,
	RunE:  server.Cmd(test.NewService(), results.NewService()),
}
