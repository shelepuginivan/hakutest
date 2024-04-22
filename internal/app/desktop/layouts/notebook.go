// Package layouts provides reusable GTK layouts.
package layouts

import "github.com/gotk3/gotk3/gtk"

// NotebookPage represents a page of GtkNotebook.
type NotebookPage struct {
	Child gtk.IWidget
	Label gtk.IWidget
}

// NewNotebook is a wrapper for gtk\_notebook\_new()
// It creates new GtkNotebook with the specified pages.
func NewNotebook(pages ...NotebookPage) (*gtk.Notebook, error) {
	nb, err := gtk.NotebookNew()
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		nb.AppendPage(page.Child, page.Label)
	}

	return nb, nil
}
