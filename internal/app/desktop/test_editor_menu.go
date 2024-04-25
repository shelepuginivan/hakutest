package desktop

import (
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
)

func (b Builder) NewEditorMenu(
	tests []string,
	onOpen func(testName string),
	onCreate func(),
) *gtk.Box {
	box := Must(layouts.NewForm())

	boxOpen := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8))

	heading := Must(components.NewHeadingH1("Test editor"))

	testComboBox := Must(gtk.ComboBoxTextNew())
	for _, t := range tests {
		testComboBox.Append(t, t)
	}
	testNameInput := Must(components.NewInput("Choose test", testComboBox))

	buttonOpen := Must(gtk.ButtonNewWithLabel("Open test"))
	buttonOpen.Connect("clicked", func() {
		onOpen(testComboBox.GetActiveID())
	})

	buttonCreate := Must(gtk.ButtonNewWithLabel("Create test"))
	buttonCreate.Connect("clicked", onCreate)

	boxOpen.PackStart(testNameInput, false, false, 4)
	boxOpen.PackStart(buttonOpen, false, false, 4)

	box.PackStart(heading, false, false, 0)
	box.PackStart(boxOpen, false, false, 0)
	box.PackStart(buttonCreate, false, false, 0)

	return box
}
