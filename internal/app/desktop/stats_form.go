package desktop

import (
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
)

func (b Builder) NewStatsForm(
	results, formats []string,
	onSubmit func(result, format string) error,
) *gtk.Box {
	box := Must(layouts.NewForm())

	inputsBox := Must(layouts.NewContainer())
	submitBox := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8))

	heading := Must(components.NewHeadingH1("Export statistics"))

	resultsComboBox := Must(gtk.ComboBoxTextNew())
	for _, r := range results {
		resultsComboBox.Append(r, r)
	}
	resultsComboBox.SetActive(0)
	resultsInput := Must(components.NewInput("Test", resultsComboBox))

	formatsComboBox := Must(gtk.ComboBoxTextNew())
	for _, f := range formats {
		formatsComboBox.Append(f, f)
	}
	formatsComboBox.SetActive(0)
	formatsInput := Must(components.NewInput("Format", formatsComboBox))

	submitButton := Must(gtk.ButtonNewWithLabel(b.app.I18n.Statistics.App.SubmitText))
	submitResult := Must(gtk.LabelNew(""))

	submitButton.Connect("clicked", func() {
		defer time.AfterFunc(time.Second*4, func() {
			submitResult.SetText("")
		})

		result := resultsComboBox.GetActiveID()
		format := formatsComboBox.GetActiveID()

		if err := onSubmit(result, format); err != nil {
			submitResult.SetText("An error occurred!")
			return
		}

		submitResult.SetText("Statistics exported to Downloads")
	})

	inputsBox.PackStart(resultsInput, false, false, 0)
	inputsBox.PackStart(formatsInput, false, false, 0)

	submitBox.PackStart(submitButton, false, false, 0)
	submitBox.PackStart(submitResult, false, false, 0)

	box.PackStart(heading, false, false, 0)
	box.PackStart(inputsBox, false, false, 0)
	box.PackStart(submitBox, false, false, 0)

	return box
}
