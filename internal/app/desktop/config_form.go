package desktop

import (
	"math"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

func (b Builder) NewConfigForm(onSubmit func(cfg *config.Config) error) *gtk.Box {
	box := Must(layouts.NewForm())

	heading := Must(components.NewHeadingH1("Configuration"))

	generalBox := Must(layouts.NewContainer())
	serverBox := Must(layouts.NewContainer())
	submitBox := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8))

	generalHeading := Must(components.NewHeadingH2("General"))
	generalHeading.SetHAlign(gtk.ALIGN_START)

	langComboBox := Must(gtk.ComboBoxTextNew())
	for langCode, lang := range i18n.LanguageCodeMap {
		langComboBox.Append(langCode, lang)
	}
	langComboBox.SetActiveID(b.app.Config.General.Language)
	langInput := Must(components.NewInput("Language", langComboBox))

	testsDirectoryEntry := Must(gtk.EntryNew())
	testsDirectoryEntry.SetText(b.app.Config.General.TestsDirectory)
	testsDirectoryInput := Must(components.NewInput("Tests directory", testsDirectoryEntry))

	resultsDirectoryEntry := Must(gtk.EntryNew())
	resultsDirectoryEntry.SetText(b.app.Config.General.ResultsDirectory)
	resultsDirectoryInput := Must(components.NewInput("Results directory", resultsDirectoryEntry))

	showResultsCheckButton := Must(gtk.CheckButtonNewWithLabel("Show results"))
	showResultsCheckButton.SetActive(b.app.Config.General.ShowResults)

	overwriteResultsCheckButton := Must(gtk.CheckButtonNewWithLabel("Overwrite results"))
	overwriteResultsCheckButton.SetActive(b.app.Config.General.OverwriteResults)

	serverHeading := Must(components.NewHeadingH2("Server"))
	serverHeading.SetHAlign(gtk.ALIGN_START)

	portSpinButton := Must(gtk.SpinButtonNewWithRange(1024, 65535, 1))
	portSpinButton.SetValue(float64(b.app.Config.Server.Port))
	portInput := Must(components.NewInput("Port", portSpinButton))

	maxUploadSizeSpinButton := Must(gtk.SpinButtonNewWithRange(0, math.MaxInt64, 1))
	maxUploadSizeSpinButton.SetValue(float64(b.app.Config.Server.MaxUploadSize))
	maxUploadSizeInput := Must(components.NewInput("Max upload size (bytes)", maxUploadSizeSpinButton))

	serverModeComboBox := Must(gtk.ComboBoxTextNew())
	for modeId, modeName := range config.ServerModeMap {
		serverModeComboBox.Append(modeId, modeName)
	}
	serverModeComboBox.SetActiveID(b.app.Config.Server.Mode)
	serverModeInput := Must(components.NewInput("Mode", serverModeComboBox))

	submitButton := Must(gtk.ButtonNewWithLabel("Save config"))
	submitResult := Must(gtk.LabelNew(""))

	submitButton.Connect("clicked", func() {
		defer time.AfterFunc(time.Second*4, func() {
			submitResult.SetText("")
		})

		cfg := &config.Config{
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

		if err := onSubmit(cfg); err != nil {
			submitResult.SetText("An error occurred!")
			return
		}

		submitResult.SetText("Saved config")
	})

	generalBox.PackStart(generalHeading, false, false, 20)
	generalBox.PackStart(langInput, false, false, 0)
	generalBox.PackStart(testsDirectoryInput, false, false, 0)
	generalBox.PackStart(resultsDirectoryInput, false, false, 0)
	generalBox.PackStart(showResultsCheckButton, false, false, 0)
	generalBox.PackStart(overwriteResultsCheckButton, false, false, 0)

	serverBox.PackStart(serverHeading, false, false, 20)
	serverBox.PackStart(portInput, false, false, 0)
	serverBox.PackStart(maxUploadSizeInput, false, false, 0)
	serverBox.PackStart(serverModeInput, false, false, 0)

	submitBox.PackStart(submitButton, false, false, 0)
	submitBox.PackStart(submitResult, false, false, 0)

	box.PackStart(heading, false, false, 0)
	box.PackStart(generalBox, false, false, 0)
	box.PackStart(serverBox, false, false, 0)
	box.PackStart(submitBox, false, false, 0)

	return box
}
