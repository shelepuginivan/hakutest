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

	i := Input{}

	i.Box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	i.SetHAlign(gtk.ALIGN_FILL)

	labelBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 0)
	if err != nil {
		return nil, err
	}

	gtkLabel, err := gtk.LabelNew(label)
	if err != nil {
		return nil, err
	}

	labelBox.SetHAlign(gtk.ALIGN_START)
	labelBox.PackStart(gtkLabel, false, false, 0)

	i.PackStart(labelBox, false, false, 0)
	i.PackStart(field, true, true, 0)

	return &i, nil
}
