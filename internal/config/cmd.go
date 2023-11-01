package config

import (
	"fmt"
	"log"

	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, args []string) {
	userConfig := Init()

	if len(args) == 1 {
		switch args[0] {
		case "port":
			fmt.Println(userConfig.Port)
		case "tests_directory":
			fmt.Println(userConfig.TestsDirectory)
		case "results_directory":
			fmt.Println(userConfig.ResultsDirectory)
		case "student_name_label":
			fmt.Println(userConfig.StudentNameLabel)
		case "open_answer_label":
			fmt.Println(userConfig.OpenAnswerLabel)
		case "submit_button_label":
			fmt.Println(userConfig.SubmitButtonLabel)
		default:
			log.Fatal("Invalid field")
		}
	} else if len(args) == 2 {
		switch args[0] {
		case "port":
			userConfig.Port = args[1]
		case "tests_directory":
			userConfig.TestsDirectory = args[1]
		case "results_directory":
			userConfig.ResultsDirectory = args[1]
		case "student_name_label":
			userConfig.StudentNameLabel = args[1]
		case "open_answer_label":
			userConfig.OpenAnswerLabel = args[1]
		case "submit_button_label":
			userConfig.SubmitButtonLabel = args[1]
		default:
			log.Fatal("Invalid field")
		}

		err := userConfig.Save()

		if err != nil {
			log.Fatal(err)
		}
	} else {
		tbl := table.New("Key", "Value")

		tbl.AddRow("port", userConfig.Port)
		tbl.AddRow("tests_directory", userConfig.TestsDirectory)
		tbl.AddRow("results_directory", userConfig.ResultsDirectory)
		tbl.AddRow("student_name_label", userConfig.StudentNameLabel)
		tbl.AddRow("open_answer_label", userConfig.OpenAnswerLabel)
		tbl.AddRow("submit_button_label", userConfig.SubmitButtonLabel)

		tbl.Print()
	}
}
