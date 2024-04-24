package components

import (
	"github.com/gotk3/gotk3/gtk"
)

// EntryListMultipleItem is a GTK component based on Gtk.Box.
// It represents an item of the EntryListMultiple.
type EntryListMultipleItem struct {
	*gtk.Box

	entry *gtk.Entry
	check *gtk.CheckButton
}

// EntryListMultipleValue represents a value of EntryListMultipleItem.
type EntryListMultipleValue struct {
	Text     string
	Selected bool
}

// EntryListSingle is a GTK component based on Gtk.Frame.
// It wraps ComponentList and contains a list of entries.
// Entries can be added or removed.
// It addition, many entries can be selected.
type EntryListMultiple struct {
	*gtk.Frame

	cl *ComponentList[*EntryListMultipleItem]
}

// NewEntryListMultipleItem returns a new instance of EntryListMultipleItem.
func (lst EntryListMultiple) NewEntryListMultipleItem(v *EntryListMultipleValue) (*EntryListMultipleItem, error) {
	var err error

	item := EntryListMultipleItem{}

	item.Box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
	if err != nil {
		return nil, err
	}

	item.check, err = gtk.CheckButtonNew()
	if err != nil {
		return nil, err
	}

	item.entry, err = gtk.EntryNew()
	if err != nil {
		return nil, err
	}

	item.check.SetActive(v.Selected)
	item.entry.SetText(v.Text)

	item.PackStart(item.check, false, false, 4)
	item.PackStart(item.entry, true, true, 4)

	return &item, nil
}

// NewEntryListMultiple returns a new instance of EntryListMultiple.
func NewEntryListMultiple(label, addButtonLabel, removeButtonLabel string) (*EntryListMultiple, error) {
	var err error

	lst := EntryListMultiple{}

	lst.Frame, err = gtk.FrameNew(label)
	if err != nil {
		return nil, err
	}

	lst.cl, err = NewComponentList(
		addButtonLabel,
		removeButtonLabel,
		func() (*EntryListMultipleItem, error) {
			return lst.NewEntryListMultipleItem(&EntryListMultipleValue{})
		},
	)

	if err != nil {
		return nil, err
	}

	lst.Add(lst.cl)

	return &lst, nil
}

// GetValues returns values of the list.
func (lst *EntryListMultiple) GetValues() []*EntryListMultipleValue {
	var values []*EntryListMultipleValue

	lst.cl.ForEach(func(item *EntryListMultipleItem) {
		text, err := item.entry.GetText()
		if err != nil {
			return
		}

		values = append(values, &EntryListMultipleValue{
			Text:     text,
			Selected: item.check.GetActive(),
		})
	})

	return values
}

// SetValues sets values of the list.
func (lst *EntryListMultiple) SetValues(values []*EntryListMultipleValue) {
	lst.cl.Clear()
	for _, v := range values {
		item, err := lst.NewEntryListMultipleItem(v)
		if err != nil {
			continue
		}

		lst.cl.AddComponent(item)
	}
}
