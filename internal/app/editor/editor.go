package editor

import (
	"fmt"
	"time"

	"github.com/Songmu/prompter"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
	"github.com/spf13/cobra"
)

func Cmd(s test.TestService) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var (
			t            = test.Test{}
			timeLayout   = "2006-01-02 15:04:05"
			tasksDeleted = 0
			name         string
		)

		if len(args) == 1 {
			name = args[0]

			parsedTest, err := s.GetByName(name)

			if err == nil {
				t = parsedTest
			}
		} else {
			name = prompter.Prompt(message("Test filename"), "test.json")
		}

		t.Title = prompter.Prompt(message("Title of the test"), t.Title)
		t.Description = prompter.Prompt(message("Description"), t.Description)
		t.Target = prompter.Prompt(message("Target audience"), t.Target)
		t.Subject = prompter.Prompt(message("Subject of the test"), t.Subject)
		t.Author = prompter.Prompt(message("Author"), t.Author)
		t.Institution = prompter.Prompt(message("Educational institution"), t.Institution)

		expiresInString := prompter.Prompt(message("Expires in"), t.ExpiresIn.Format(timeLayout))

		if expiresInString != "" {
			expiresIn, err := time.Parse(timeLayout, expiresInString)

			if err == nil {
				t.ExpiresIn = expiresIn
			}
		}

		for i := 0; i < len(t.Tasks); {
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
				t.Tasks[i] = promptEditTask(t.Tasks[i])
				i++
			case "replace":
				t.Tasks[i] = promptNewTask()
				i++
			case "remove":
				t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
				tasksDeleted++
			}
		}

		for prompter.YN(message("Add new task"), false) {
			t.Tasks = append(t.Tasks, promptNewTask())
		}

		return s.SaveToTestsDirectory(t, name)
	}
}
