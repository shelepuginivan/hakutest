package desktop

import (
	"slices"
	"strconv"
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// TaskInput is a GTK component based on Gtk.Frame.
// It is used to edit test task.
type TaskInput struct {
	*gtk.Frame

	typeComboBox             *gtk.ComboBoxText
	textView                 *gtk.TextView
	answerEntry              *gtk.Entry
	answerInput              *components.Input
	answerOptionsSingle      *components.EntryListSingle
	answerOptionsMultiple    *components.EntryListMultiple
	hasAttachmentCheckButton *gtk.CheckButton
	attachmentInput          *AttachmentInput
}

// NewTaskInput returns a new instance of TaskInput.
func (b Builder) NewTaskInput() *TaskInput {
	ti := TaskInput{
		Frame: Must(gtk.FrameNew("Task")),

		typeComboBox:             Must(gtk.ComboBoxTextNew()),
		textView:                 Must(gtk.TextViewNew()),
		answerEntry:              Must(gtk.EntryNew()),
		answerOptionsSingle:      Must(components.NewEntryListSingle("Options", "+ Add answer", "-")),
		answerOptionsMultiple:    Must(components.NewEntryListMultiple("Options", "+ Add answer", "-")),
		hasAttachmentCheckButton: Must(gtk.CheckButtonNewWithLabel("Include attachment")),
		attachmentInput:          b.NewAttachmentInput(),
	}

	ti.answerInput = Must(components.NewInput("Answer", ti.answerEntry))

	ti.typeComboBox.Append(test.TaskSingle, "Single answer")
	ti.typeComboBox.Append(test.TaskMultiple, "Multiple answers")
	ti.typeComboBox.Append(test.TaskOpen, "Open question")
	ti.typeComboBox.Append(test.TaskFile, "Answer with file(s)")
	ti.typeComboBox.SetActiveID(test.TaskSingle)

	box := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6))

	box.SetMarginStart(12)
	box.SetMarginEnd(12)

	typeInput := Must(components.NewInput("Type", ti.typeComboBox))
	textInput := Must(components.NewInput("Text", ti.textView))

	ti.typeComboBox.Connect("changed", func() {
		ti.SetType(ti.typeComboBox.GetActiveID())
	})

	ti.hasAttachmentCheckButton.Connect("toggled", func() {
		ti.attachmentInput.SetVisible(ti.hasAttachmentCheckButton.GetActive())
	})

	ti.Connect("show", func() {
		ti.SetType(ti.typeComboBox.GetActiveID())
		ti.attachmentInput.SetVisible(ti.hasAttachmentCheckButton.GetActive())
	})

	box.PackStart(typeInput, false, false, 4)
	box.PackStart(textInput, false, false, 4)
	box.PackStart(ti.answerOptionsSingle, false, false, 4)
	box.PackStart(ti.answerOptionsMultiple, false, false, 4)
	box.PackStart(ti.answerInput, false, false, 4)
	box.PackStart(ti.hasAttachmentCheckButton, false, false, 4)
	box.PackStart(ti.attachmentInput, false, false, 4)

	ti.Add(box)

	return &ti
}

// SetType sets the type of the task.
// It shows and hides respective fields.
func (ti *TaskInput) SetType(t string) {
	switch t {
	case test.TaskSingle:
		ti.answerInput.Hide()
		ti.answerOptionsMultiple.Hide()
		ti.answerOptionsSingle.Show()
	case test.TaskMultiple:
		ti.answerInput.Hide()
		ti.answerOptionsMultiple.Show()
		ti.answerOptionsSingle.Hide()
	case test.TaskOpen:
		ti.answerInput.Show()
		ti.answerOptionsMultiple.Hide()
		ti.answerOptionsSingle.Hide()
	default:
		ti.answerInput.Hide()
		ti.answerOptionsMultiple.Hide()
		ti.answerOptionsSingle.Hide()
	}
}

// GetTask returns the task.
func (ti TaskInput) GetTask() (*test.Task, error) {
	task := &test.Task{}

	task.Type = ti.typeComboBox.GetActiveID()
	b, err := ti.textView.GetBuffer()
	if err != nil {
		return task, err
	}

	task.Text, err = b.GetText(b.GetStartIter(), b.GetEndIter(), true)
	if err != nil {
		return task, err
	}

	task.Attachment, err = ti.attachmentInput.GetAttachment()
	if err != nil {
		return task, err
	}

	switch task.Type {
	case test.TaskSingle:
		for i, v := range ti.answerOptionsSingle.GetValues() {
			task.Options = append(task.Options, v.Text)

			if v.Selected {
				task.Answer = strconv.Itoa(i + 1)
			}
		}
	case test.TaskMultiple:
		var ans []string

		for idx, v := range ti.answerOptionsMultiple.GetValues() {
			task.Options = append(task.Options, v.Text)

			if v.Selected {
				ans = append(ans, strconv.Itoa(idx+1))
			}
		}
		task.Answer = strings.Join(ans, ",")
	case test.TaskOpen:
		task.Answer, err = ti.answerEntry.GetText()
	}

	return task, err
}

// SetTask sets the task.
func (ti *TaskInput) SetTask(task *test.Task) error {
	ti.typeComboBox.SetActiveID(task.Type)

	b, err := ti.textView.GetBuffer()
	if err != nil {
		return err
	}
	b.SetText(task.Text)

	if task.Attachment != nil {
		ti.hasAttachmentCheckButton.SetActive(true)
		ti.attachmentInput.SetAttachment(task.Attachment)
	}

	switch task.Type {
	case test.TaskSingle:
		var options []*components.EntryListSingleValue

		for i, o := range task.Options {
			options = append(options, &components.EntryListSingleValue{
				Text:     o,
				Selected: strconv.Itoa(i+1) == task.Answer,
			})
		}

		ti.answerOptionsSingle.SetValues(options)
	case test.TaskMultiple:
		var options []*components.EntryListMultipleValue
		ans := strings.Split(task.Answer, ",")

		for i, o := range task.Options {
			options = append(options, &components.EntryListMultipleValue{
				Text:     o,
				Selected: slices.Contains(ans, strconv.Itoa(i+1)),
			})
		}

		ti.answerOptionsMultiple.SetValues(options)
	case test.TaskOpen:
		ti.answerEntry.SetText(task.Answer)
	}

	ti.SetType(task.Type)

	return nil
}
