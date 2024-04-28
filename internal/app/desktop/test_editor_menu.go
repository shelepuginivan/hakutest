package desktop

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

// EditorMenu is a GTK component based on Gtk.Box.
// It represents the main menu of the test editor.
type EditorMenu struct {
	*gtk.Box

	i18n *i18n.GtkEditorMenuI18n
}

// NewEditorMenu returns a new instance of EditorMenu.
func (b Builder) NewEditorMenu(
	tests []string,
	onOpen func(testName string),
	onCreate func(),
) *EditorMenu {
	box := &EditorMenu{
		Box: Must(layouts.NewForm()),

		i18n: b.app.I18n.Gtk.Editor.Menu,
	}

	boxOpen := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8))

	heading := Must(components.NewHeadingH1(box.i18n.Title))

	testComboBox := Must(gtk.ComboBoxTextNew())
	for _, t := range tests {
		testComboBox.Append(t, t)
	}
	testNameInput := Must(components.NewInput(box.i18n.InputTest, testComboBox))

	buttonOpen := Must(gtk.ButtonNewWithLabel(box.i18n.ButtonOpen))
	buttonOpen.Connect("clicked", func() {
		onOpen(testComboBox.GetActiveID())
	})

	buttonCreate := Must(gtk.ButtonNewWithLabel(box.i18n.ButtonCreate))
	buttonCreate.Connect("clicked", onCreate)

	boxOpen.PackStart(testNameInput, false, false, 4)
	boxOpen.PackStart(buttonOpen, false, false, 4)

	box.PackStart(heading, false, false, 0)
	box.PackStart(boxOpen, false, false, 0)
	box.PackStart(buttonCreate, false, false, 0)

	return box
}
