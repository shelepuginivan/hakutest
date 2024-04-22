package desktop

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
)

type GtkComponentsBuilder struct {
	app *application.App
	win *gtk.Window
}

func NewGtkComponentsBuilder(app *application.App, win *gtk.Window) *GtkComponentsBuilder {
	return &GtkComponentsBuilder{
		app: app,
		win: win,
	}
}

func (b GtkComponentsBuilder) NewInput(
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

func (b GtkComponentsBuilder) NewStatsForm(
	results, formats []string,
	onSubmit func(result, format string) error,
) *gtk.Box {
	formBox := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6))

	formBox.SetVAlign(gtk.ALIGN_CENTER)
	formBox.SetHAlign(gtk.ALIGN_CENTER)

	headerLabel := Must(gtk.LabelNew("Statistics Export"))

	resultLabel := Must(gtk.LabelNew(b.app.I18n.Statistics.App.LabelTest))
	resultComboBox := Must(gtk.ComboBoxTextNew())

	for _, result := range results {
		resultComboBox.AppendText(result)
	}

	resultComboBox.SetActive(0)

	formatLabel := Must(gtk.LabelNew(b.app.I18n.Statistics.App.LabelFormat))
	formatComboBox := Must(gtk.ComboBoxTextNew())

	for _, format := range formats {
		formatComboBox.AppendText(format)
	}

	formatComboBox.SetActive(0)

	submitButton := Must(gtk.ButtonNewWithLabel(b.app.I18n.Statistics.App.SubmitText))
	submitResult := Must(gtk.LabelNew(""))

	submitButton.Connect("clicked", func() {
		result := resultComboBox.GetActiveText()
		format := formatComboBox.GetActiveText()

		err := onSubmit(result, format)

		if err != nil {
			submitResult.SetText(fmt.Sprintf(
				"%s %s",
				b.app.I18n.Statistics.App.ErrorPrefix,
				err.Error(),
			))
		} else {
			submitResult.SetText(b.app.I18n.Statistics.App.SuccessText)
		}

		time.AfterFunc(time.Second*4, func() {
			submitResult.SetText("")
		})
	})

	formBox.PackStart(headerLabel, false, false, 24)
	formBox.PackStart(b.NewInput(resultLabel, resultComboBox), false, false, 16)
	formBox.PackStart(b.NewInput(formatLabel, formatComboBox), false, false, 16)
	formBox.PackStart(submitButton, false, false, 4)
	formBox.PackStart(submitResult, false, false, 4)

	return formBox
}
