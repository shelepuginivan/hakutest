package components

import "github.com/gotk3/gotk3/gtk"

// Input is a GTK component based on Gtk.Box.
// It is an input field with text label and widget.
type Input struct {
	*gtk.Box
}

// NewInput returns a new instance of Input.
func NewInput(label string, field gtk.IWidget) (*Input, error) {
	var err error

	i := &Input{}

	i.Box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}
	i.SetHAlign(gtk.ALIGN_FILL)

	fieldLabel, err := gtk.LabelNew(label)
	if err != nil {
		return nil, err
	}
	fieldLabel.SetHAlign(gtk.ALIGN_START)

	i.PackStart(fieldLabel, false, false, 0)
	i.PackStart(field, true, true, 0)

	return i, nil
}
