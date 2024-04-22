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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:   "hakutest",
	Short: "Reliable and efficient educational testing platform",
	Long:  "Reliable and efficient educational testing platform",
}
