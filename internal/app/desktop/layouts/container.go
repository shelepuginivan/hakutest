package layouts

import "github.com/gotk3/gotk3/gtk"

// NewContainer returns a new container layout.
func NewContainer() (*gtk.Box, error) {
	b, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 28)
	if err != nil {
		return nil, err
	}

	b.SetHAlign(gtk.ALIGN_FILL)
	b.SetVAlign(gtk.ALIGN_CENTER)

	return b, nil
}
