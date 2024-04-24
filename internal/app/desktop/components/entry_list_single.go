package components

import (
	"github.com/gotk3/gotk3/gtk"
)

// EntryListSingleItem is a GTK component based on Gtk.Box.
// It represents an item of the EntryListSingle.
type EntryListSingleItem struct {
	*gtk.Box

	entry *gtk.Entry
	radio *gtk.RadioButton
}

// EntryListSingleValue represents a value of EntryListSingleItem.
type EntryListSingleValue struct {
	Text     string
	Selected bool
}

// EntryListSingle is a GTK component based on Gtk.Frame.
// It wraps ComponentList and contains a list of entries.
// Entries can be added or removed.
// It addition, one entry can be selected.
type EntryListSingle struct {
	*gtk.Frame

	baseRadio *gtk.RadioButton
	cl        *ComponentList[*EntryListSingleItem]
}

// NewEntryListSingleItem returns a new instance of EntryListSingleItem.
func (lst EntryListSingle) NewEntryListSingleItem(v *EntryListSingleValue) (*EntryListSingleItem, error) {
	var err error

	item := EntryListSingleItem{}

	item.Box, err = gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
	if err != nil {
		return nil, err
	}

	item.Connect("show", func(i *gtk.Box) {
		i.ShowAll()
	})

	item.radio, err = gtk.RadioButtonNewFromWidget(lst.baseRadio)
	if err != nil {
		return nil, err
	}

	item.entry, err = gtk.EntryNew()
	if err != nil {
		return nil, err
	}

	item.entry.SetText(v.Text)
	item.radio.SetActive(v.Selected)

	item.PackStart(item.radio, false, false, 4)
	item.PackStart(item.entry, true, true, 4)

	return &item, nil
}

// NewEntryListSingle returns a new instance of EntryListSingle.
func NewEntryListSingle(label, addButtonLabel, removeButtonLabel string) (*EntryListSingle, error) {
	var err error

	lst := EntryListSingle{}

	lst.Frame, err = gtk.FrameNew(label)
	if err != nil {
		return nil, err
	}

	lst.baseRadio, err = gtk.RadioButtonNew(nil)
	if err != nil {
		return nil, err
	}

	lst.cl, err = NewComponentList(
		addButtonLabel,
		removeButtonLabel,
		func() (*EntryListSingleItem, error) {
			return lst.NewEntryListSingleItem(&EntryListSingleValue{})
		},
	)
	if err != nil {
		return nil, err
	}

	lst.Add(lst.cl)

	return &lst, nil
}

// GetValues returns values of the list.
func (lst EntryListSingle) GetValues() []*EntryListSingleValue {
	var values []*EntryListSingleValue

	lst.cl.ForEach(func(item *EntryListSingleItem) {
		text, err := item.entry.GetText()
		if err != nil {
			return
		}

		values = append(values, &EntryListSingleValue{
			Text:     text,
			Selected: item.radio.GetActive(),
		})
	})

	return values
}

// SetValues sets values of the list.
func (lst *EntryListSingle) SetValues(values []*EntryListSingleValue) {
	lst.cl.Clear()
	for _, v := range values {
		item, err := lst.NewEntryListSingleItem(v)
		if err != nil {
			continue
		}
		lst.cl.AddComponent(item)
	}
}
