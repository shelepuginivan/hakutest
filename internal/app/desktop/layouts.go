// Package desktop provides shared Fyne UI components: layouts, forms, etc.
package desktop

import (
	"fyne.io/fyne/v2"
)

// PaddedLayout stacks objects vertically.
// It adds Padding between each object and around all objects.
// Objects are displayed at their MinSize, both vertically and horizontally.
type PaddedLayout struct {
	Padding float32
}

// NewPaddedLayout creates a new PaddedLayout instance with specified padding.
func NewPaddedLayout(padding float32) *PaddedLayout {
	return &PaddedLayout{Padding: padding}
}

func (p *PaddedLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := 2*p.Padding, p.Padding

	for _, o := range objects {
		childSize := o.MinSize()

		w = fyne.Max(w, childSize.Width+2*p.Padding)
		h += childSize.Height + p.Padding
	}

	return fyne.NewSize(w, h)
}

func (p *PaddedLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	pos := fyne.NewPos(p.Padding, p.Padding)

	for _, o := range objects {
		size := o.MinSize()
		o.Resize(size)
		o.Move(pos)

		pos = pos.Add(fyne.NewPos(0, size.Height+p.Padding))
	}
}
