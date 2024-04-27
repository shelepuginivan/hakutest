package desktop

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// AttachmentInput is a GTK component based on Gtk.Frame.
// It allows to configure attachment for the test task.
type AttachmentInput struct {
	*gtk.Frame

	i18n             *i18n.GtkEditorAttachmentI18n
	nameEntry        *gtk.Entry
	typeComboBox     *gtk.ComboBoxText
	attachmentSource *components.AttachmentSource
}

// NewAttachmentInput returns a new instance of AttachmentInput.
func (b Builder) NewAttachmentInput() *AttachmentInput {
	ai := AttachmentInput{
		Frame: Must(gtk.FrameNew("")),

		i18n:         b.app.I18n.Gtk.Editor.Task.Attachment,
		nameEntry:    Must(gtk.EntryNew()),
		typeComboBox: Must(gtk.ComboBoxTextNew()),
	}

	ai.attachmentSource = Must(components.NewAttachmentSource(
		ai.i18n.LabelModeFile,
		ai.i18n.LabelModeUrl,
		ai.i18n.LabelModeLoaded,
		ai.i18n.FileDialogTitle,
		ai.i18n.FileDialogButtonOpen,
		ai.i18n.FileDialogButtonCancel,
	))

	box := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4))

	nameInput := Must(components.NewInput(ai.i18n.InputName, ai.nameEntry))

	typeMap := map[string]string{
		test.AttachmentFile:  ai.i18n.LabelTypeFile,
		test.AttachmentImage: ai.i18n.LabelTypeImage,
		test.AttachmentAudio: ai.i18n.LabelTypeAudio,
		test.AttachmentVideo: ai.i18n.LabelTypeVideo,
	}

	for id, typeText := range typeMap {
		ai.typeComboBox.Append(id, typeText)
	}

	ai.typeComboBox.SetActiveID(test.AttachmentFile)

	typeInput := Must(components.NewInput(ai.i18n.InputType, ai.typeComboBox))

	box.PackStart(nameInput, false, false, 4)
	box.PackStart(typeInput, false, false, 4)
	box.PackStart(ai.attachmentSource, false, false, 4)
	box.SetMarginStart(8)
	box.SetMarginEnd(8)

	ai.Add(box)

	return &ai
}

// GetAttachment returns the configured attachment.
func (ai AttachmentInput) GetAttachment() (*test.Attachment, error) {
	var err error
	attachment := test.Attachment{
		Type: ai.typeComboBox.GetActiveID(),
	}

	attachment.Name, err = ai.nameEntry.GetText()
	if err != nil {
		return &attachment, err
	}

	attachment.Src, err = ai.attachmentSource.GetSource()

	return &attachment, err
}

// SetAttachment configures the attachment.
func (ai *AttachmentInput) SetAttachment(attachment *test.Attachment) {
	ai.nameEntry.SetText(attachment.Name)
	ai.typeComboBox.SetActiveID(attachment.Type)
	ai.attachmentSource.SetSource(attachment.Src)
}
