package desktop

import (
	"github.com/gotk3/gotk3/gtk"
)

func (b Builder) NewBaseForm() *gtk.Box {
	formBox := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6))

	formBox.SetVAlign(gtk.ALIGN_CENTER)
	formBox.SetMarginStart(20)
	formBox.SetMarginEnd(20)

	return formBox
}

func (b Builder) NewInput(
	label *gtk.Label,
	field gtk.IWidget,
) *gtk.Box {
	input := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6))
	input.SetHAlign(gtk.ALIGN_FILL)

	inputLabel := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6))
	inputLabel.SetHAlign(gtk.ALIGN_START)
	inputLabel.PackStart(label, false, false, 0)

	input.PackStart(inputLabel, false, false, 0)
	input.PackStart(field, true, true, 0)

	return input
}

func (b Builder) NewInputGroup(
	legend *gtk.Label,
	inputs ...gtk.IWidget,
) *gtk.Box {
	inputGroup := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8))
	inputGroup.PackStart(legend, false, false, 10)

	for _, input := range inputs {
		inputGroup.PackStart(input, false, false, 8)
	}

	return inputGroup
}
