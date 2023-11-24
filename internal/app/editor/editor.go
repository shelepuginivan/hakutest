package editor

import (
	"fmt"
	"time"

	"github.com/Songmu/prompter"
	"github.com/shelepuginivan/hakutest/internal/pkg/core"
	"github.com/spf13/cobra"
)

func Cmd(cmd *cobra.Command, args []string) error {
	var (
		test         = core.Test{}
		timeLayout   = "2006-01-02 15:04:05"
		tasksDeleted = 0
		name         string
	)

	if len(args) == 1 {
		name = args[0]

		parsedTest, err := core.Get(name)

		if err == nil {
			test = parsedTest
		}
	} else {
		name = prompter.Prompt(message("Test filename"), "test.json")
	}

	test.Title = prompter.Prompt(message("Title of the test"), test.Title)
	test.Description = prompter.Prompt(message("Description"), test.Description)
	test.Target = prompter.Prompt(message("Target audience"), test.Target)
	test.Subject = prompter.Prompt(message("Subject of the test"), test.Subject)
	test.Author = prompter.Prompt(message("Author"), test.Author)
	test.Institution = prompter.Prompt(message("Educational institution"), test.Institution)

	expiresInString := prompter.Prompt(message("Expires in"), test.ExpiresIn.Format(timeLayout))

	if expiresInString != "" {
		expiresIn, err := time.Parse(timeLayout, expiresInString)

		if err == nil {
			test.ExpiresIn = expiresIn
		}
	}

	for i := 0; i < len(test.Tasks); {
		action := prompter.Choose(
			message(fmt.Sprintf("Task %d:", i+tasksDeleted+1)),
			[]string{"leave unchanged", "edit", "replace", "remove"},
			"leave unchanged",
		)

		switch action {
		case "leave unchanged":
			i++
			continue
		case "edit":
			test.Tasks[i] = promptEditTask(test.Tasks[i])
			i++
		case "replace":
			test.Tasks[i] = promptNewTask()
			i++
		case "remove":
			test.Tasks = append(test.Tasks[:i], test.Tasks[i+1:]...)
			tasksDeleted++
		}
	}

	for prompter.YN(message("Add new task"), false) {
		test.Tasks = append(test.Tasks, promptNewTask())
	}

	return test.Save(name)
}
