package components

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

// AttachmentInput is a GTK component based on Gtk.Frame.
// It allows to configure attachment for the test task.
type AttachmentInput struct {
	*gtk.Frame

	nameEntry        *gtk.Entry
	typeComboBox     *gtk.ComboBoxText
	attachmentSource *AttachmentSource
}

// NewAttachmentInput returns a new instance of AttachmentInput.
func NewAttachmentInput(
	label,
	nameLabel,
	typeLabel,
	modeFileLabel,
	modeUrlLabel,
	dialogTitle,
	dialogOpenButtonText,
	dialogCancelButtonText string,
	typeMap map[string]string,
) (*AttachmentInput, error) {
	var err error

	ai := AttachmentInput{}

	ai.Frame, err = gtk.FrameNew(label)
	if err != nil {
		return nil, err
	}

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	ai.nameEntry, err = gtk.EntryNew()
	if err != nil {
		return nil, err
	}

	nameInput, err := NewInput(nameLabel, ai.nameEntry)
	if err != nil {
		return nil, err
	}

	ai.typeComboBox, err = gtk.ComboBoxTextNew()
	if err != nil {
		return nil, err
	}

	for id, typeText := range typeMap {
		ai.typeComboBox.Append(id, typeText)
	}

	typeInput, err := NewInput(typeLabel, ai.typeComboBox)
	if err != nil {
		return nil, err
	}

	ai.attachmentSource, err = NewAttachmentSource(
		modeFileLabel,
		modeUrlLabel,
		dialogTitle,
		dialogOpenButtonText,
		dialogCancelButtonText,
	)
	if err != nil {
		return nil, err
	}

	box.PackStart(nameInput, false, false, 4)
	box.PackStart(typeInput, false, false, 4)
	box.PackStart(ai.attachmentSource, false, false, 4)

	ai.Add(box)

	return &ai, err
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
