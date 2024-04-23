package desktop

import (
	"log"
	"math"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

func (b Builder) NewConfigForm(onSubmit func(cfg *config.Config) error) *gtk.Box {
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
