package server

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Init(cmd *cobra.Command, args []string) error {
	r := NewRouter()
	port, err := cmd.Flags().GetInt("port")

	if err != nil {
		return err
	}

	addr := fmt.Sprintf(":%d", port)

	return r.Run(addr)
}
