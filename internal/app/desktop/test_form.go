package desktop

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// TestForm is a GTK component based on Gtk.Box.
// It is a form of the test editor.
type TestForm struct {
	*gtk.Box

	titleEntry          *gtk.Entry
	descriptionTextView *gtk.TextView
	subjectEntry        *gtk.Entry
	authorEntry         *gtk.Entry
	targetAudienceEntry *gtk.Entry
	institutionEntry    *gtk.Entry
	expiresAtCheck      *gtk.CheckButton
	expiresAtInput      *components.DatetimePicker
	taskList            *TaskList
}

// NewTestForm returns a new instance of TestForm.
func (b Builder) NewTestForm(
	onSubmit func(t *test.Test) error,
	onCancel func(),
) *TestForm {
	form := &TestForm{
		Box: b.NewBaseForm(),

		titleEntry:          Must(gtk.EntryNew()),
		descriptionTextView: Must(gtk.TextViewNew()),
		subjectEntry:        Must(gtk.EntryNew()),
		authorEntry:         Must(gtk.EntryNew()),
		targetAudienceEntry: Must(gtk.EntryNew()),
		institutionEntry:    Must(gtk.EntryNew()),
		expiresAtCheck:      Must(gtk.CheckButtonNewWithLabel("Expires at")),
		expiresAtInput:      Must(components.NewDatetimePicker()),
		taskList:            b.NewTaskList(),
	}

	form.Connect("show", form.expiresAtInput.Hide)

	cancelButton := Must(gtk.ButtonNewWithLabel("Close"))
	cancelButton.Connect("clicked", onCancel)

	cancelButtonBox := Must(gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0))
	cancelButtonBox.SetHAlign(gtk.ALIGN_END)
	cancelButtonBox.PackStart(cancelButton, false, false, 0)

	titleInput := Must(components.NewInput("Title", form.titleEntry))
	descriptionInput := Must(components.NewInput("Description", form.descriptionTextView))
	subjectInput := Must(components.NewInput("Subject", form.subjectEntry))
	authorInput := Must(components.NewInput("Author", form.authorEntry))
	targetAudienceInput := Must(components.NewInput("Target Audience", form.targetAudienceEntry))
	institutionInput := Must(components.NewInput("Institution", form.institutionEntry))

	form.expiresAtCheck.Connect("clicked", func(w *gtk.CheckButton) {
		form.expiresAtInput.SetVisible(w.GetActive())
	})

	submitButton := Must(gtk.ButtonNewWithLabel("Save"))
	submitResultLabel := Must(gtk.LabelNew(""))

	submitButton.Connect("clicked", func() {
		defer time.AfterFunc(time.Second*4, func() {
			submitResultLabel.SetText("")
		})

		test, err := form.GetTest()
		if err != nil {
			submitResultLabel.SetText("An error occurred!")
			return
		}

		if err = onSubmit(test); err != nil {
			submitResultLabel.SetText("An error occurred!")
			return
		}

		submitResultLabel.SetText("Test saved to the tests directory")
	})

	form.PackStart(cancelButtonBox, false, false, 24)
	form.PackStart(titleInput, false, false, 8)
	form.PackStart(descriptionInput, false, false, 8)
	form.PackStart(subjectInput, false, false, 8)
	form.PackStart(authorInput, false, false, 8)
	form.PackStart(targetAudienceInput, false, false, 8)
	form.PackStart(institutionInput, false, false, 8)
	form.PackStart(form.expiresAtCheck, false, false, 8)
	form.PackStart(form.expiresAtInput, false, false, 8)
	form.PackStart(Must(gtk.SeparatorNew(gtk.ORIENTATION_HORIZONTAL)), false, false, 16)
	form.PackStart(form.taskList, false, false, 8)
	form.PackStart(submitButton, false, false, 24)
	form.PackStart(submitResultLabel, false, false, 8)

	return form
}

// GetTest returns values of the TestForm.
func (form TestForm) GetTest() (*test.Test, error) {
	var err error
	t := &test.Test{}

	t.Title, err = form.titleEntry.GetText()
	if err != nil {
		return t, err
	}

	b, err := form.descriptionTextView.GetBuffer()
	if err != nil {
		return t, err
	}

	t.Description, err = b.GetText(b.GetStartIter(), b.GetEndIter(), true)
	if err != nil {
		return t, err
	}

	t.Subject, err = form.subjectEntry.GetText()
	if err != nil {
		return t, err
	}

	t.Author, err = form.authorEntry.GetText()
	if err != nil {
		return t, err
	}

	t.Target, err = form.targetAudienceEntry.GetText()
	if err != nil {
		return t, err
	}

	t.Institution, err = form.institutionEntry.GetText()
	if err != nil {
		return t, err
	}

	if form.expiresAtCheck.GetActive() {
		t.ExpiresIn = form.expiresAtInput.GetDate()
	}

	t.Tasks = form.taskList.GetTasks()

	return t, nil
}

// SetTest sets values of the TestForm.
func (form *TestForm) SetTest(t *test.Test) {
	form.titleEntry.SetText(t.Title)

	b, err := form.descriptionTextView.GetBuffer()
	if err == nil {
		b.SetText(t.Description)
	}

	form.subjectEntry.SetText(t.Subject)
	form.authorEntry.SetText(t.Author)
	form.targetAudienceEntry.SetText(t.Target)
	form.institutionEntry.SetText(t.Institution)

	if !t.ExpiresIn.IsZero() {
		form.expiresAtCheck.SetActive(true)
		form.expiresAtInput.SetDate(t.ExpiresIn)
	}

	form.taskList.SetTasks(t.Tasks)
}
