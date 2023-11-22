package editor

import (
	"fmt"
	"time"

	"github.com/Songmu/prompter"
	parser "github.com/shelepuginivan/hakutest/internal/pkg/test_parser"
	"github.com/spf13/cobra"
)

func promptTask() parser.Task {
	task := parser.Task{}
	taskType := map[string]string{
		"Single answer":    "single",
		"Multiple answers": "multiple",
		"Open question":    "open",
	}

	task.Type = taskType[prompter.Choose(
		secondaryMessage("Type of the task"),
		[]string{"Single answer", "Multiple answers", "Open question"},
		"Open question",
	)]

	task.Text = prompter.Prompt(secondaryMessage("Task text"), "")

	fmt.Println(secondaryMessage("Answer options:"))

	option := prompter.Prompt(nestedMessage("Add option (leave blank to stop)", 2), "")

	for option != "" {
		task.Options = append(task.Options, option)
		option = prompter.Prompt(nestedMessage("Add option (leave blank to stop)", 2), "")
	}

	task.Answer = prompter.Prompt(secondaryMessage("Correct answer"), "")

	if prompter.YN(secondaryMessage("Add attachment"), false) {
		task.Attachment.Name = prompter.Prompt(nestedMessage("Name", 2), "")
		task.Attachment.Type = prompter.Choose(
			nestedMessage("Type", 2),
			[]string{"image", "video", "audio", "file"},
			"file",
		)

		src := prompter.Prompt(nestedMessage("Source (URL or local file)", 2), "")
		attachmentSrc, err := getAttachmentSrc(src)

		for err != nil && prompter.YN(secondaryMessage("Failed to add attachment! Try again?"), false) {
			src = prompter.Prompt(nestedMessage("Source (URL or local file)", 2), "")
			attachmentSrc, err = getAttachmentSrc(src)
		}

		task.Attachment.Src = attachmentSrc
	}

	return task
}

func Cmd(cmd *cobra.Command, args []string) error {
	var (
		test         = parser.Test{}
		timeLayout   = "2006-01-02 15:04:05"
		tasksDeleted = 0
		name         string
	)

	if len(args) == 1 {
		name = args[0]

		parsedTest, err := parser.Get(name)

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
			[]string{"leave unchanged", "replace", "delete"},
			"leave unchanged",
		)

		switch action {
		case "leave unchanged":
			i++
			continue
		case "replace":
			test.Tasks[i] = promptTask()
			i++
			continue
		case "delete":
			test.Tasks = append(test.Tasks[:i], test.Tasks[i+1:]...)
			tasksDeleted++
			continue
		}
	}

	for prompter.YN(message("Add new task"), false) {
		test.Tasks = append(test.Tasks, promptTask())
	}

	return test.Save(name)
}
