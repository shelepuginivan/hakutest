package desktop

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
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

func (b GtkComponentsBuilder) NewBaseForm() *gtk.Box {
	formBox := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6))

	formBox.SetVAlign(gtk.ALIGN_CENTER)
	formBox.SetMarginStart(20)
	formBox.SetMarginEnd(20)

	return formBox
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

func (b GtkComponentsBuilder) NewInputGroup(
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

func (b GtkComponentsBuilder) NewStatsForm(
	results, formats []string,
	onSubmit func(result, format string) error,
) *gtk.Box {
	formBox := b.NewBaseForm()

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

	formBox.PackStart(b.NewInput(resultLabel, resultComboBox), false, false, 16)
	formBox.PackStart(b.NewInput(formatLabel, formatComboBox), false, false, 16)
	formBox.PackStart(submitButton, false, false, 4)
	formBox.PackStart(submitResult, false, false, 4)

	return formBox
}

func (b GtkComponentsBuilder) NewConfigForm(onSubmit func(cfg *config.Config) error) *gtk.Box {
	formBox := b.NewBaseForm()

	langComboBox := Must(gtk.ComboBoxTextNew())

	for langCode, lang := range i18n.LanguageCodeMap {
		langComboBox.Append(langCode, lang)
	}

	langComboBox.SetActiveID(b.app.Config.General.Language)

	testsDirectoryEntry := Must(gtk.EntryNew())
	testsDirectoryEntry.SetText(b.app.Config.General.TestsDirectory)

	resultsDirectoryEntry := Must(gtk.EntryNew())
	resultsDirectoryEntry.SetText(b.app.Config.General.ResultsDirectory)

	showResultsCheckButton := Must(gtk.CheckButtonNew())
	showResultsCheckButton.SetActive(b.app.Config.General.ShowResults)

	overwriteResultsCheckButton := Must(gtk.CheckButtonNew())
	overwriteResultsCheckButton.SetActive(b.app.Config.General.OverwriteResults)

	generalConfigGroup := b.NewInputGroup(
		Must(gtk.LabelNew("General")),
		b.NewInput(
			Must(gtk.LabelNew("Language")),
			langComboBox,
		),
		b.NewInput(
			Must(gtk.LabelNew("Show results")),
			showResultsCheckButton,
		),
		b.NewInput(
			Must(gtk.LabelNew("Overwrite results")),
			overwriteResultsCheckButton,
		),
		b.NewInput(
			Must(gtk.LabelNew("Tests directory")),
			testsDirectoryEntry,
		),
		b.NewInput(
			Must(gtk.LabelNew("Results directory")),
			resultsDirectoryEntry,
		),
	)

	portSpinButton := Must(gtk.SpinButtonNewWithRange(1024, 65535, 1))
	portSpinButton.SetValue(float64(b.app.Config.Server.Port))

	maxUploadSizeSpinButton := Must(gtk.SpinButtonNewWithRange(0, math.MaxInt64, 1))
	maxUploadSizeSpinButton.SetValue(float64(b.app.Config.Server.MaxUploadSize))

	serverModeComboBox := Must(gtk.ComboBoxTextNew())

	for modeId, modeName := range config.ServerModeMap {
		serverModeComboBox.Append(modeId, modeName)
	}

	serverModeComboBox.SetActiveID(b.app.Config.Server.Mode)

	serverConfigGroup := b.NewInputGroup(
		Must(gtk.LabelNew("Server")),
		b.NewInput(
			Must(gtk.LabelNew("Port")),
			portSpinButton,
		),
		b.NewInput(
			Must(gtk.LabelNew("Max upload size (bytes)")),
			maxUploadSizeSpinButton,
		),
		b.NewInput(
			Must(gtk.LabelNew("Mode")),
			serverModeComboBox,
		),
	)

	submitButton := Must(gtk.ButtonNewWithLabel("Save config"))
	submitResult := Must(gtk.LabelNew(""))

	submitButton.Connect("clicked", func() {
		cfg := config.Config{
			General: config.GeneralConfig{
				Language:         langComboBox.GetActiveID(),
				TestsDirectory:   Must(testsDirectoryEntry.GetText()),
				ResultsDirectory: Must(resultsDirectoryEntry.GetText()),
				ShowResults:      showResultsCheckButton.GetActive(),
				OverwriteResults: overwriteResultsCheckButton.GetActive(),
			},
			Server: config.ServerConfig{
				Port:          portSpinButton.GetValueAsInt(),
				Mode:          serverModeComboBox.GetActiveID(),
				MaxUploadSize: int64(maxUploadSizeSpinButton.GetValueAsInt()),
			},
		}

		if err := onSubmit(&cfg); err != nil {
			log.Print(err)
			submitResult.SetText("An error occurred!")
		} else {
			submitResult.SetText("Saved config")
		}

		time.AfterFunc(time.Second*4, func() {
			submitResult.SetText("")
		})
	})

	formBox.PackStart(generalConfigGroup, false, false, 16)
	formBox.PackStart(serverConfigGroup, false, false, 16)
	formBox.PackStart(submitButton, false, false, 16)
	formBox.PackStart(submitResult, false, false, 16)

	return formBox
}
