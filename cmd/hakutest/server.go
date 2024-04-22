package main

import (
	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/spf13/cobra"
)

func init() {
	serverCmd.Flags().IntVarP(
		&app.Config.Server.Port,
		"port",
		"p",
		app.Config.Server.Port,
		"The port on which the server will be started",
	)
}

func serverCommand(cmd *cobra.Command, args []string) error {
	srv := server.New(app)
	return srv.ListenAndServe()
}

var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "Start Hakutest server",
	Long:    "Start Hakutest server",
	Args:    cobra.NoArgs,
	RunE:    serverCommand,
	Aliases: []string{"srv"},
}
