package desktop

import (
	"fmt"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

// ServerForm is a GTK component based on Gtk.Box.
// It provides a form to export test statistics.
type StatsForm struct {
	*gtk.Box

	i18n *i18n.GtkStatsI18n
}

// NewStatsForm returns a new instance of StatsForm.
func (b Builder) NewStatsForm(
	results, formats []string,
	onSubmit func(result, format string) error,
) *StatsForm {
	form := &StatsForm{
		Box: Must(layouts.NewForm()),

		i18n: b.app.I18n.Gtk.Stats,
	}

	inputsBox := Must(layouts.NewContainer())
	submitBox := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8))

	heading := Must(components.NewHeadingH1(form.i18n.Title))

	resultsComboBox := Must(gtk.ComboBoxTextNew())
	for _, r := range results {
		resultsComboBox.Append(r, r)
	}
	resultsComboBox.SetActive(0)
	resultsInput := Must(components.NewInput(form.i18n.InputTest, resultsComboBox))

	formatsComboBox := Must(gtk.ComboBoxTextNew())
	for _, f := range formats {
		formatsComboBox.Append(f, f)
	}
	formatsComboBox.SetActive(0)
	formatsInput := Must(components.NewInput(form.i18n.InputFormat, formatsComboBox))

	submitButton := Must(gtk.ButtonNewWithLabel(b.app.I18n.Statistics.App.SubmitText))
	submitResult := Must(gtk.LabelNew(""))

	submitButton.Connect("clicked", func() {
		defer time.AfterFunc(time.Second*4, func() {
			submitResult.SetText("")
		})

		result := resultsComboBox.GetActiveID()
		format := formatsComboBox.GetActiveID()

		if err := onSubmit(result, format); err != nil {
			submitResult.SetText(fmt.Sprintf(
				form.i18n.LabelError,
				err.Error(),
			))
			return
		}

		submitResult.SetText(form.i18n.LabelSuccess)
	})

	inputsBox.PackStart(resultsInput, false, false, 0)
	inputsBox.PackStart(formatsInput, false, false, 0)

	submitBox.PackStart(submitButton, false, false, 0)
	submitBox.PackStart(submitResult, false, false, 0)

	form.PackStart(heading, false, false, 0)
	form.PackStart(inputsBox, false, false, 0)
	form.PackStart(submitBox, false, false, 0)

	return form
}
