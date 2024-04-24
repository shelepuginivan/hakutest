package desktop

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// AttachmentInput is a GTK component based on Gtk.Frame.
// It allows to configure attachment for the test task.
type AttachmentInput struct {
	*gtk.Frame

	nameEntry        *gtk.Entry
	typeComboBox     *gtk.ComboBoxText
	attachmentSource *components.AttachmentSource
}

// NewAttachmentInput returns a new instance of AttachmentInput.
func (b Builder) NewAttachmentInput() *AttachmentInput {
	ai := AttachmentInput{
		Frame: Must(gtk.FrameNew("")),

		nameEntry:    Must(gtk.EntryNew()),
		typeComboBox: Must(gtk.ComboBoxTextNew()),
		attachmentSource: Must(components.NewAttachmentSource(
			"Choose file",
			"Enter URL",
			"Open file",
			"Open",
			"Cancel",
		)),
	}

	box := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4))

	nameInput := Must(components.NewInput("Name", ai.nameEntry))

	typeMap := map[string]string{
		test.AttachmentFile:  "File",
		test.AttachmentImage: "Image",
		test.AttachmentAudio: "Audio",
		test.AttachmentVideo: "Video",
	}

	for id, typeText := range typeMap {
		ai.typeComboBox.Append(id, typeText)
	}

	ai.typeComboBox.SetActiveID(test.AttachmentFile)

	typeInput := Must(components.NewInput("Type", ai.typeComboBox))

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
	ai.attachmentSource.SetUrl(attachment.Src)
}
