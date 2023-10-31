package cmd

import (
	"fmt"
	"os"

	"github.com/shelepuginivan/hakutest/internal/app/server"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "hakutest [port]",
	RunE: server.Init,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
