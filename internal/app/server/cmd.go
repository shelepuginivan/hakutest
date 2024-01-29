package server

import (
	"net/http"

	"github.com/spf13/cobra"
)

func Cmd(handler http.Handler) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		port, err := cmd.Flags().GetInt("port")

		if err != nil {
			return err
		}

		srv := NewServer(handler, port)
		return srv.ListenAndServe()
	}
}
