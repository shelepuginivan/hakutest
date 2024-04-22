package components

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop"
)

func (b Builder) NewStatsForm(
	results, formats []string,
	onSubmit func(result, format string) error,
) *gtk.Box {
	formBox := b.NewBaseForm()

	resultLabel := desktop.Must(gtk.LabelNew(b.app.I18n.Statistics.App.LabelTest))
	resultComboBox := desktop.Must(gtk.ComboBoxTextNew())

	for _, result := range results {
		resultComboBox.AppendText(result)
	}

	resultComboBox.SetActive(0)

	formatLabel := desktop.Must(gtk.LabelNew(b.app.I18n.Statistics.App.LabelFormat))
	formatComboBox := desktop.Must(gtk.ComboBoxTextNew())

	for _, format := range formats {
		formatComboBox.AppendText(format)
	}

	formatComboBox.SetActive(0)

	submitButton := desktop.Must(gtk.ButtonNewWithLabel(b.app.I18n.Statistics.App.SubmitText))
	submitResult := desktop.Must(gtk.LabelNew(""))

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

	formBox.PackStart(b.NewInput(resultLabel, resultComboBox), false, false, 16)
	formBox.PackStart(b.NewInput(formatLabel, formatComboBox), false, false, 16)
	formBox.PackStart(submitButton, false, false, 4)
	formBox.PackStart(submitResult, false, false, 4)

	return formBox
}
