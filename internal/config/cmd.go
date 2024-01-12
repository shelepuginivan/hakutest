package config

import (
	"log"

	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, args []string) {
	New()

	var err error = nil

	switch len(args) {
	case 0:
		err = Print()
	case 1:
		err = PrintField(args[0])
	case 2:
		err = SetField(args[0], args[1])
	}

	if err != nil {
		log.Fatal(err)
	}
}
