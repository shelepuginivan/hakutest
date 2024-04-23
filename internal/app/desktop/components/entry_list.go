package components

import (
	"container/list"

	"github.com/gotk3/gotk3/gtk"
)

// EntryList is a GTK component based on Gtk.Frame.
// It allows to enter multiple strings.
// Entries can be added and removed.
type EntryList struct {
	*gtk.Frame

	entriesBox *gtk.Box
	entries    *list.List
}

// NewEntryList returns a new instance of EntryList.
func NewEntryList(label string, addButtonLabel string) (*EntryList, error) {
	var err error

	el := EntryList{
		entries: list.New(),
	}

	el.Frame, err = gtk.FrameNew(label)
	if err != nil {
		return nil, err
	}

	el.entriesBox, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	buttonAddEntry, err := gtk.ButtonNewWithLabel(addButtonLabel)
	if err != nil {
		return nil, err
	}

	buttonAddEntry.Connect("clicked", func() {
		el.AddEntry("")
	})

	box, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	if err != nil {
		return nil, err
	}

	box.PackStart(el.entriesBox, false, false, 4)
	box.PackStart(buttonAddEntry, false, false, 4)

	el.SetMarginStart(12)
	el.SetMarginEnd(12)
	el.Add(box)

	return &el, nil
}

// AddEntry adds a new entry of the specified value.
func (el *EntryList) AddEntry(value string) error {
	entryBox, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 4)
	if err != nil {
		return err
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		return err
	}

	entry.SetText(value)
	listElement := el.entries.PushBack(entry)

	entryRemoveButton, err := gtk.ButtonNewWithLabel("-")
	if err != nil {
		return err
	}

	entryRemoveButton.Connect("clicked", func() {
		el.entries.Remove(listElement)
		entryBox.Destroy()
	})

	entryBox.PackStart(entry, true, true, 2)
	entryBox.PackStart(entryRemoveButton, false, false, 2)

	el.entriesBox.PackStart(entryBox, false, false, 4)
	el.ShowAll()
	return nil
}

// ClearEntries removes all entries.
func (el *EntryList) ClearEntries() error {
	var err error
	el.entriesBox.Destroy()
	el.entriesBox, err = gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 4)
	return err
}

// GetEntries returns values of each entry.
func (el EntryList) GetEntries() []string {
	var entries []string

	for e := el.entries.Front(); e != nil; e = e.Next() {
		entry, ok := e.Value.(*gtk.Entry)
		if !ok {
			continue
		}

		entryValue, err := entry.GetText()
		if err != nil {
			continue
		}

		entries = append(entries, entryValue)
	}

	return entries
}

// SetEntries sets the entries.
func (el *EntryList) SetEntries(entries []string) error {
	if err := el.ClearEntries(); err != nil {
		return err
	}

	for _, entry := range entries {
		if err := el.AddEntry(entry); err != nil {
			return err
		}
	}

	return nil
}
