package editor

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Songmu/prompter"
	"github.com/fatih/color"
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

func message(s string) string {
	return color.New(color.Bold, color.FgYellow).Sprint(s)
}

func secondaryMessage(s string) string {
	return color.New(color.FgMagenta, color.Bold).Sprintf("- %s", s)
}

func nestedMessage(s string, level int) string {
	return color.New(color.Bold).Sprintf("%s %s", strings.Repeat("-", level), s)
}

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

		parsedTest, err := parser.ParseTest(name)

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
