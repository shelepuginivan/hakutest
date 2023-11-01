package editor

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/Songmu/prompter"
	"github.com/gabriel-vasile/mimetype"
	parser "github.com/shelepuginivan/hakutest/internal/pkg/test_parser"
	"github.com/spf13/cobra"
)

func getAttachmentSrc(src string) (string, error) {
	if bytes, err := os.ReadFile(src); err == nil {
		mimeType := mimetype.Detect(bytes)
		base64Endoding := base64.StdEncoding.EncodeToString(bytes)

		return fmt.Sprintf("data:%s;base64,%s", mimeType, base64Endoding), nil
	}

	_, err := url.ParseRequestURI(src)

	if err == nil {
		return src, nil
	}

	return "", err
}

func promptTask() parser.Task {
	task := parser.Task{}
	taskType := map[string]string{
		"Single answer":    "single",
		"Multiple answers": "multiple",
		"Open question":    "open",
	}

	task.Type = taskType[prompter.Choose(
		"Type of the task",
		[]string{"Single answer", "Multiple answers", "Open question"},
		"Open question",
	)]

	task.Text = prompter.Prompt("Task text", "")

	option := prompter.Prompt("Answer option (leave blank to stop)", "")

	for option != "" {
		task.Options = append(task.Options, option)
		option = prompter.Prompt("Answer option (leave blank to stop)", "")
	}

	task.Answer = prompter.Prompt("Correct answer", "")

	if prompter.YN("Add attachment", false) {
		task.Attachment.Name = prompter.Prompt("Name", "")
		task.Attachment.Type = prompter.Choose(
			"Type",
			[]string{"image", "video", "audio", "file"},
			"file",
		)

		src := prompter.Prompt("Source (URL or local file)", "")
		attachmentSrc, err := getAttachmentSrc(src)

		for err != nil && prompter.YN("Failed to add attachment! Try again?", false) {
			src = prompter.Prompt("Source (URL or local file)", "")
			attachmentSrc, err = getAttachmentSrc(src)
		}

		task.Attachment.Src = attachmentSrc
	}

	return task
}

func EditorCmd(cmd *cobra.Command, args []string) error {
	var (
		test         = parser.Test{}
		timeLayout   = "2006-01-02 15:04:05"
		tasksDeleted = 0
		name         string
	)

	if len(args) == 1 {
		name = args[0]

		parsedTest, err := parser.ParseTest(name)

		if err == nil {
			test = parsedTest
		}
	} else {
		name = prompter.Prompt("Test filename", "test.json")
	}

	test.Title = prompter.Prompt("Title of the test", test.Title)
	test.Target = prompter.Prompt("Target audience", test.Target)
	test.Subject = prompter.Prompt("Subject of the test", test.Subject)
	test.Institution = prompter.Prompt("Educational institution", test.Institution)

	expiresInString := prompter.Prompt("Expires in", test.ExpiresIn.Format(timeLayout))

	if expiresInString != "" {
		expiresIn, err := time.Parse(timeLayout, expiresInString)

		if err == nil {
			test.ExpiresIn = expiresIn
		}
	}

	for i := 0; i < len(test.Tasks); {
		action := prompter.Choose(
			fmt.Sprintf("Task %d:", i+tasksDeleted+1),
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

	for prompter.YN("Add new task", false) {
		test.Tasks = append(test.Tasks, promptTask())
	}

	return test.Save(name)
}
