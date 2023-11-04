package config

import (
	"log"

	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, args []string) {
	userConfig := Init()

	switch len(args) {
	case 0:
		userConfig.Print()
	case 1:
		userConfig.PrintField(args[0])
	case 2:
		err := userConfig.SetField(args[0], args[1])

		if err != nil {
			log.Fatal(err)
		}

		userConfig.Save()
	}
}
