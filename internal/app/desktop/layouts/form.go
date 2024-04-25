package layouts

import "github.com/gotk3/gotk3/gtk"

// NewForm returns a new form layout.
func NewForm() (*gtk.Box, error) {
	b, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 80)
	if err != nil {
		return nil, err
	}

	b.SetHAlign(gtk.ALIGN_FILL)
	b.SetVAlign(gtk.ALIGN_CENTER)
	b.SetMarginTop(80)
	b.SetMarginBottom(80)
	b.SetMarginStart(24)
	b.SetMarginEnd(24)

	return b, nil
}
