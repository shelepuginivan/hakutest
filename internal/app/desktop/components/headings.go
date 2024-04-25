package components

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

// NewHeadingH1 returns a new Gtk.Label with `size='xx-large'` and `weight='bold'`.
func NewHeadingH1(str string) (*gtk.Label, error) {
	h, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}

	h.SetMarkup(fmt.Sprintf("<span size='xx-large' weight='bold'>%s</span>", str))

	return h, nil
}

// NewHeadingH2 returns a new Gtk.Label with `size='x-large'` and `weight='bold'`.
func NewHeadingH2(str string) (*gtk.Label, error) {
	h, err := gtk.LabelNew("")
	if err != nil {
		return nil, err
	}

	h.SetMarkup(fmt.Sprintf("<span size='x-large' weight='bold'>%s</span>", str))

	return h, nil
}
