package editor

import (
	"fmt"

	"github.com/Songmu/prompter"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

func promptAddAttachment() test.Attachment {
	attachment := test.Attachment{}

	attachment.Name = prompter.Prompt(nestedMessage("Name", 2), "")
	attachment.Type = prompter.Choose(
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

	attachment.Src = attachmentSrc

	return attachment
}

func promptEditAttachment(attachment test.Attachment) test.Attachment {
	attachment.Name = prompter.Prompt(nestedMessage("Name", 2), attachment.Name)
	attachment.Type = prompter.Choose(
		nestedMessage("Type", 2),
		[]string{"image", "video", "audio", "file"},
		attachment.Type,
	)

	if prompter.YN(secondaryMessage("Edit attachment source"), false) {
		src := prompter.Prompt(nestedMessage("Source (URL or local file)", 2), "")
		attachmentSrc, err := getAttachmentSrc(src)

		for err != nil && prompter.YN(secondaryMessage("Failed to add attachment! Try again?"), false) {
			src = prompter.Prompt(nestedMessage("Source (URL or local file)", 2), "")
			attachmentSrc, err = getAttachmentSrc(src)
		}

		attachment.Src = attachmentSrc
	}

	return attachment
}

func promptNewTask() test.Task {
	task := test.Task{}
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
		task.Attachment = promptAddAttachment()
	}

	return task
}

func promptEditTask(task test.Task) test.Task {
	taskTypeName := map[string]string{
		"single":   "Single answer",
		"multiple": "Multiple answers",
		"open":     "Open question",
	}

	taskType := map[string]string{
		"Single answer":    "single",
		"Multiple answers": "multiple",
		"Open question":    "open",
	}

	task.Type = taskType[prompter.Choose(
		secondaryMessage("Type of the task"),
		[]string{"Single answer", "Multiple answers", "Open question"},
		taskTypeName[task.Type],
	)]

	task.Text = prompter.Prompt(secondaryMessage("Task text"), task.Text)

	fmt.Println(secondaryMessage("Answer options:"))

	for i := 0; i < len(task.Options); {
		option := task.Options[i]

		switch prompter.Choose(
			nestedMessage(fmt.Sprintf("%d) %s", i+1, option), 2),
			[]string{"leave unchanged", "remove", "edit"},
			"leave unchanged",
		) {
		case "remove":
			task.Options = append(task.Options[:i], task.Options[i+1:]...)
		case "edit":
			task.Options[i] = prompter.Prompt(nestedMessage("Answer option", 3), option)
			i++
		case "leave unchanged":
			i++
		}
	}

	option := prompter.Prompt(nestedMessage("Add option (leave blank to stop)", 2), "")

	for option != "" {
		task.Options = append(task.Options, option)
		option = prompter.Prompt(nestedMessage("Add option (leave blank to stop)", 2), "")
	}

	task.Answer = prompter.Prompt(secondaryMessage("Correct answer"), task.Answer)

	switch prompter.Choose(
		secondaryMessage("Task attachment"),
		[]string{"leave unchanged", "edit", "remove"},
		"leave unchanged",
	) {
	case "edit":
		task.Attachment = promptEditAttachment(task.Attachment)
	case "remove":
		task.Attachment = test.Attachment{}
	}

	return task
}
