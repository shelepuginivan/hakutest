package components

import (
	"encoding/base64"
	"fmt"
	"mime"
	"net/url"
	"os"

	"github.com/gotk3/gotk3/gtk"
)

type AttachmentMode = string

const (
	AttachmentSourceModeLoaded AttachmentMode = "loaded"
	AttachmentSourceModeFile   AttachmentMode = "file"
	AttachmentSourceModeURL    AttachmentMode = "url"
)

// AttachmentSource is a GTK components based on Gtk.Box.
// It allows to select attachment source either from local file or from the
// URL.
type AttachmentSource struct {
	*gtk.Box

	mode         AttachmentMode
	loadedSource string

	stack *gtk.Stack

	loadedLabel *gtk.Label
	fileButton  *gtk.FileChooserButton
	urlEntry    *gtk.Entry

	baseRadio     *gtk.RadioButton
	modeUrlRadio  *gtk.RadioButton
	modeFileRadio *gtk.RadioButton
}

// NewAttachmentSource returns a new instance of AttachmentSource.
func NewAttachmentSource(
	modeFileLabel,
	modeUrlLabel,
	modeLoadedLabel,
	dialogTitle,
	dialogOpenButtonText,
	dialogCancelButtonText string,
) (*AttachmentSource, error) {
	var err error

	as := &AttachmentSource{
		mode: AttachmentSourceModeURL,
	}

	as.Box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	as.baseRadio, err = gtk.RadioButtonNew(nil)
	if err != nil {
		return nil, err
	}

	as.modeUrlRadio, err = gtk.RadioButtonNewWithLabelFromWidget(as.baseRadio, modeUrlLabel)
	if err != nil {
		return nil, err
	}

	as.modeFileRadio, err = gtk.RadioButtonNewWithLabelFromWidget(as.baseRadio, modeFileLabel)
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

	as.loadedLabel, err = gtk.LabelNew(modeLoadedLabel)
	if err != nil {
		return nil, err
	}

	as.stack, err = gtk.StackNew()
	if err != nil {
		return nil, err
	}

	as.stack.AddNamed(as.urlEntry, AttachmentSourceModeURL)
	as.stack.AddNamed(as.fileButton, AttachmentSourceModeFile)
	as.stack.AddNamed(as.loadedLabel, AttachmentSourceModeLoaded)

	as.modeFileRadio.Connect("toggled", func(w *gtk.RadioButton) {
		if w.GetActive() {
			as.SetMode(AttachmentSourceModeFile)
		}
	})

	as.modeUrlRadio.Connect("toggled", func(w *gtk.RadioButton) {
		if w.GetActive() {
			as.SetMode(AttachmentSourceModeURL)
		}
	})

	as.Connect("show", func() {
		as.SetMode(as.mode)
	})

	as.PackStart(as.modeUrlRadio, false, false, 4)
	as.PackStart(as.modeFileRadio, false, false, 4)
	as.PackStart(as.stack, false, false, 4)

	return as, nil
}

func (as *AttachmentSource) SetMode(mode AttachmentMode) {
	as.mode = mode
	as.baseRadio.SetActive(mode == AttachmentSourceModeLoaded)
	as.modeFileRadio.SetActive(mode == AttachmentSourceModeFile)
	as.modeUrlRadio.SetActive(mode == AttachmentSourceModeURL)

	as.stack.SetVisibleChildName(mode)
}

// GetSource returns the entered source.
//
// If mode is set to file, it encodes the file to base64 URL.
// If mode is set to URL, it returns the entered URL.
func (as *AttachmentSource) GetSource() (string, error) {
	if as.mode == AttachmentSourceModeLoaded {
		return as.loadedSource, nil
	}

	if as.mode == AttachmentSourceModeURL {
		return as.urlEntry.GetText()
	}

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

// SetSource sets the URL of the AttachmentSource.
func (as *AttachmentSource) SetSource(src string) {
	_, err := url.ParseRequestURI(src)
	if err != nil {
		as.SetMode(AttachmentSourceModeLoaded)
		as.loadedSource = src
		return
	}

	as.SetMode(AttachmentSourceModeURL)
	as.urlEntry.SetText(src)
}
