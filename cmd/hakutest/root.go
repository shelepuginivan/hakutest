package main

import (
	"fmt"
	"os"

	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/spf13/cobra"
)

var app *application.App

func init() {
	app = application.New()

	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(resultsCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(statisticsCmd)
	rootCmd.AddCommand(testsCmd)
}

var rootCmd = &cobra.Command{
	Use:   "hakutest",
	Short: "Reliable and efficient educational testing platform",
	Long:  "Reliable and efficient educational testing platform",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
