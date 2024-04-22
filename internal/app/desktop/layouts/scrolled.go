package layouts

import "github.com/gotk3/gotk3/gtk"

// NewScrolled a wrapper for gtk\_scrolled\_window\_new().
// It creates new GtkScrolledWindow with specified widget w as child.
func NewScrolled(w gtk.IWidget) (*gtk.ScrolledWindow, error) {
	scrolled, err := gtk.ScrolledWindowNew(nil, nil)
	if err != nil {
		return nil, err
	}

	scrolled.Add(w)
	return scrolled, nil
}
