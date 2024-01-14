package server

import (
	"fmt"

	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func Cmd(t test.TestService, r results.ResultsService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		router := NewRouter(t, r)
		port, err := cmd.Flags().GetInt("port")

		if err != nil {
			return err
		}

		addr := fmt.Sprintf(":%d", port)

		return router.Run(addr)
	}
}
