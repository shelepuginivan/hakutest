package components

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/url"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

// AttachmentSource is a GTK components based on Gtk.Box.
// It allows to select attachment source either from local file or from the
// URL.
type AttachmentSource struct {
	*gtk.Box

	isFile     bool
	fileButton *gtk.FileChooserButton
	urlEntry   *gtk.Entry
}

// NewAttachmentSource returns a new instance of AttachmentSource.
func NewAttachmentSource(
	modeFileLabel,
	modeUrlLabel,
	dialogTitle,
	dialogOpenButtonText,
	dialogCancelButtonText string) (*AttachmentSource, error) {
	var err error

	as := AttachmentSource{}
	as.Box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	modeUrlRadio, err := gtk.RadioButtonNewWithLabel(nil, modeUrlLabel)
	if err != nil {
		return nil, err
	}

	modeFileRadio, err := gtk.RadioButtonNewWithLabelFromWidget(modeUrlRadio, modeFileLabel)
	if err != nil {
		return nil, err
	}

	as.urlEntry, err = gtk.EntryNew()
	if err != nil {
		return nil, err
	}

	fileChooserDialog, err := gtk.FileChooserDialogNewWith2Buttons(
		dialogTitle, nil,
		gtk.FILE_CHOOSER_ACTION_OPEN,
		dialogCancelButtonText, gtk.RESPONSE_CANCEL,
		dialogOpenButtonText, gtk.RESPONSE_ACCEPT,
	)
	if err != nil {
		return nil, err
	}

	as.fileButton, err = gtk.FileChooserButtonNewWithDialog(fileChooserDialog)
	if err != nil {
		return nil, err
	}

	modeUrlRadio.Connect("toggled", func() {
		as.isFile = modeFileRadio.GetActive()
		as.fileButton.SetVisible(as.isFile)
		as.urlEntry.SetVisible(!as.isFile)
	})

	as.Connect("show", func() {
		as.fileButton.SetVisible(as.isFile)
		as.urlEntry.SetVisible(!as.isFile)
	})

	as.PackStart(modeUrlRadio, false, false, 4)
	as.PackStart(modeFileRadio, false, false, 4)
	as.PackStart(as.urlEntry, false, false, 4)
	as.PackStart(as.fileButton, false, false, 4)

	return &as, nil
}

// SetModeFile sets mode of the AttachmentSource to file.
// It hides the URL entry and shows the file chooser button.
func (as *AttachmentSource) SetModeFile() {
	as.isFile = true
	as.fileButton.Show()
	as.urlEntry.Hide()
}

// SetModeUrl sets mode of the AttachmentSource to URL.
// It hides the file chooser button and shows the URL entry.
func (as *AttachmentSource) SetModeUrl() {
	as.isFile = false
	as.fileButton.Hide()
	as.urlEntry.Show()
}

// GetSource returns the entered source.
//
// If mode is set to file, it encodes the file to base64 URL.
// If mode is set to URL, it returns the entered URL.
func (as *AttachmentSource) GetSource() (string, error) {
	if as.isFile {
		path := as.fileButton.GetFilename()

		fileBytes, err := os.ReadFile(path)
		if err != nil {
			return "", err
		}

		mimeType := mime.TypeByExtension(path)
		base64Content := base64.StdEncoding.EncodeToString(fileBytes)

		base64URL := fmt.Sprintf("data:%s;base64,%s", mimeType, base64Content)
		encodedBase64URL := url.PathEscape(base64URL)

		return encodedBase64URL, nil
	}

	return as.urlEntry.GetText()
}

// SetUrl sets the URL of the AttachmentSource.
func (as *AttachmentSource) SetUrl(url string) {
	as.urlEntry.SetText(url)
}
