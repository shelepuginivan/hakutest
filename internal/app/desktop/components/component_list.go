package components

import (
	"container/list"

	"github.com/gotk3/gotk3/gtk"
)

// ComponentListItem is a GTK component based on Gtk.Box.
// It wraps items of the ComponentList.
// This component contains a remove button. When clicked, the item is removed
// from the list.
type ComponentListItem[T gtk.IWidget] struct {
	*gtk.Box

	w T
}

// ComponentList is a GTK component based on Gtk.Box.
// It wraps a box with components.
// Each component must be of the same type.
// Components can be added and removed.
type ComponentList[T gtk.IWidget] struct {
	*gtk.Box

	list              *list.List
	componentBox      *gtk.Box
	removeButtonLabel string
}

// NewComponentList returns a new instance of NewComponentList.
func NewComponentList[T gtk.IWidget](
	addButtonLabel,
	removeButtonLabel string,
	componentConstructor func() (T, error),
) (*ComponentList[T], error) {
	var err error

	cl := ComponentList[T]{
		list:              list.New(),
		removeButtonLabel: removeButtonLabel,
	}

	cl.Box, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8)
	if err != nil {
		return nil, err
	}

	cl.componentBox, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	buttonAdd, err := gtk.ButtonNewWithLabel(addButtonLabel)
	if err != nil {
		return nil, err
	}

	buttonAdd.Connect("clicked", func() {
		w, err := componentConstructor()
		if err != nil {
			return
		}
		cl.AddComponent(w)
	})

	cl.PackStart(cl.componentBox, false, false, 4)
	cl.PackStart(buttonAdd, false, false, 4)
	cl.SetMarginStart(8)
	cl.SetMarginEnd(8)

	return &cl, nil
}

// Clear removes all inner components.
func (cl *ComponentList[T]) Clear() {
	for i := cl.list.Front(); i != nil; i = i.Next() {
		if w, ok := i.Value.(*ComponentListItem[T]); ok {
			cl.list.Remove(i)
			w.Destroy()
		}
	}
}

// AddComponent adds a new component to the list.
func (cl *ComponentList[T]) AddComponent(w T) error {
	var err error

	box := &ComponentListItem[T]{}

	box.Box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
	if err != nil {
		return err
	}

	removeButton, err := gtk.ButtonNewWithLabel(cl.removeButtonLabel)
	if err != nil {
		return err
	}

	box.w = w

	box.PackStart(box.w, true, true, 4)
	box.PackStart(removeButton, false, false, 4)

	listItem := cl.list.PushBack(box)

	removeButton.Connect("clicked", func() {
		box.Destroy()
		cl.list.Remove(listItem)
	})

	cl.componentBox.PackStart(box, false, false, 4)
	cl.componentBox.ShowAll()

	return nil
}

// ForEach performs a callback cb on each component in the list.
func (cl ComponentList[T]) ForEach(cb func(w T)) {
	for i := cl.list.Front(); i != nil; i = i.Next() {
		if box, ok := i.Value.(*ComponentListItem[T]); ok {
			cb(box.w)
		}
	}
}
