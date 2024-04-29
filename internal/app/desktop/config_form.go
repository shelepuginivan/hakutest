package desktop

import (
	"fmt"
	"math"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/components"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
)

// ConfigGeneralContainer is a GTK component based on Gtk.Form.
// It provides general configuration fields for the ConfigForm.
type ConfigGeneralContainer struct {
	*gtk.Box

	i18n                  *i18n.GtkConfigGeneralI18n
	langComboBox          *gtk.ComboBoxText
	testsDirectoryEntry   *gtk.FileChooserButton
	resultsDirectoryEntry *gtk.FileChooserButton
	showResultsCheck      *gtk.CheckButton
	overwriteResultsCheck *gtk.CheckButton
}

// NewConfigGeneralContainer returns a new instance of ConfigGeneralContainer.
func (b Builder) NewConfigGeneralContainer() *ConfigGeneralContainer {
	c := &ConfigGeneralContainer{
		Box: Must(layouts.NewContainer()),

		i18n:         b.app.I18n.Gtk.Config.General,
		langComboBox: Must(gtk.ComboBoxTextNew()),
	}

	heading := Must(components.NewHeadingH2(c.i18n.Title))
	heading.SetHAlign(gtk.ALIGN_START)

	for _, lang := range i18n.AvailableLanguages {
		c.langComboBox.Append(lang, i18n.LanguageCodeMap[lang])
	}
	c.langComboBox.SetActiveID(b.app.Config.General.Language)
	langInput := Must(components.NewInput(c.i18n.InputLang, c.langComboBox))

	testsDirectoryDialog := Must(gtk.FileChooserDialogNewWith2Buttons(
		c.i18n.FileDialogTitle, nil,
		gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
		c.i18n.FileDialogButtonCancel, gtk.RESPONSE_CANCEL,
		c.i18n.FileDialogButtonOpen, gtk.RESPONSE_ACCEPT,
	))

	c.testsDirectoryEntry = Must(gtk.FileChooserButtonNewWithDialog(testsDirectoryDialog))
	c.testsDirectoryEntry.SetCurrentFolder(b.app.Config.General.TestsDirectory)
	testsDirectoryInput := Must(components.NewInput(
		c.i18n.InputTestsDirectory,
		c.testsDirectoryEntry,
	))
	testsDirectoryInput.SetTooltipText(c.i18n.InputTestsDirectoryTooltip)

	resultsDirectoryDialog := Must(gtk.FileChooserDialogNewWith2Buttons(
		c.i18n.FileDialogTitle, nil,
		gtk.FILE_CHOOSER_ACTION_SELECT_FOLDER,
		c.i18n.FileDialogButtonCancel, gtk.RESPONSE_CANCEL,
		c.i18n.FileDialogButtonOpen, gtk.RESPONSE_ACCEPT,
	))

	c.resultsDirectoryEntry = Must(gtk.FileChooserButtonNewWithDialog(resultsDirectoryDialog))
	c.resultsDirectoryEntry.SetCurrentFolder(b.app.Config.General.ResultsDirectory)
	resultsDirectoryInput := Must(components.NewInput(
		c.i18n.InputResultsDirectory,
		c.resultsDirectoryEntry,
	))
	resultsDirectoryInput.SetTooltipText(c.i18n.InputResultsDirectoryTooltip)

	c.showResultsCheck = Must(gtk.CheckButtonNewWithLabel(c.i18n.InputShowResults))
	c.showResultsCheck.SetTooltipText(c.i18n.InputShowResultsTooltip)
	c.showResultsCheck.SetActive(b.app.Config.General.ShowResults)

	c.overwriteResultsCheck = Must(gtk.CheckButtonNewWithLabel(c.i18n.InputOverwriteResults))
	c.overwriteResultsCheck.SetTooltipText(c.i18n.InputOverwriteResultsTooltip)
	c.overwriteResultsCheck.SetActive(b.app.Config.General.OverwriteResults)

	c.PackStart(heading, false, false, 20)
	c.PackStart(langInput, false, false, 0)
	c.PackStart(testsDirectoryInput, false, false, 0)
	c.PackStart(resultsDirectoryInput, false, false, 0)
	c.PackStart(c.showResultsCheck, false, false, 0)
	c.PackStart(c.overwriteResultsCheck, false, false, 0)

	return c
}

// GetValues returns values of the general config inputs.
func (c ConfigGeneralContainer) GetValues() *config.GeneralConfig {
	return &config.GeneralConfig{
		Language:         c.langComboBox.GetActiveID(),
		TestsDirectory:   c.testsDirectoryEntry.GetFilename(),
		ResultsDirectory: c.resultsDirectoryEntry.GetFilename(),
		ShowResults:      c.showResultsCheck.GetActive(),
		OverwriteResults: c.overwriteResultsCheck.GetActive(),
	}
}

// ConfigServerContainer is a GTK component based on Gtk.Form.
// It provides server configuration fields for the ConfigForm.
type ConfigServerContainer struct {
	*gtk.Box

	i18n              *i18n.GtkConfigServerI18n
	portSpin          *gtk.SpinButton
	maxUploadSizeSpin *gtk.SpinButton
	modeComboBox      *gtk.ComboBoxText
}

// NewConfigServerContainer returns a new instance of ConfigServerContainer.
func (b Builder) NewConfigServerContainer() *ConfigServerContainer {
	c := &ConfigServerContainer{
		Box: Must(layouts.NewContainer()),

		i18n:              b.app.I18n.Gtk.Config.Server,
		portSpin:          Must(gtk.SpinButtonNewWithRange(1024, 65535, 1)),
		maxUploadSizeSpin: Must(gtk.SpinButtonNewWithRange(0, math.MaxInt64, 1)),
		modeComboBox:      Must(gtk.ComboBoxTextNew()),
	}

	heading := Must(components.NewHeadingH2(c.i18n.Title))
	heading.SetHAlign(gtk.ALIGN_START)

	c.portSpin.SetValue(float64(b.app.Config.Server.Port))
	portInput := Must(components.NewInput(c.i18n.InputPort, c.portSpin))
	portInput.SetTooltipText(c.i18n.InputPortTooltip)

	c.maxUploadSizeSpin.SetValue(float64(b.app.Config.Server.MaxUploadSize))
	maxUploadSizeInput := Must(components.NewInput(
		c.i18n.InputMaxUploadSize,
		c.maxUploadSizeSpin,
	))
	maxUploadSizeInput.SetTooltipText(c.i18n.InputMaxUploadSizeTooltip)

	for modeId, modeName := range config.ServerModeMap {
		c.modeComboBox.Append(modeId, modeName)
	}
	c.modeComboBox.SetActiveID(b.app.Config.Server.Mode)
	serverModeInput := Must(components.NewInput(c.i18n.InputMode, c.modeComboBox))
	serverModeInput.SetTooltipText(c.i18n.InputModeTooltip)

	c.PackStart(heading, false, false, 20)
	c.PackStart(portInput, false, false, 0)
	c.PackStart(maxUploadSizeInput, false, false, 0)
	c.PackStart(serverModeInput, false, false, 0)

	return c
}

// GetValues returns values of the server config inputs.
func (c ConfigServerContainer) GetValues() *config.ServerConfig {
	return &config.ServerConfig{
		Port:          c.portSpin.GetValueAsInt(),
		MaxUploadSize: int64(c.maxUploadSizeSpin.GetValueAsInt()),
		Mode:          c.modeComboBox.GetActiveID(),
	}
}

// ConfigForm is a GTK component based on Gtk.Box.
// It provides a form to configure Hakutest.
type ConfigForm struct {
	*gtk.Box

	i18n *i18n.GtkConfigI18n
}

// NewConfigForm returns a new instance of ConfigForm.
func (b Builder) NewConfigForm(onSubmit func(cfg *config.Config) error) *ConfigForm {
	form := &ConfigForm{
		Box: Must(layouts.NewForm()),

		i18n: b.app.I18n.Gtk.Config,
	}

	heading := Must(components.NewHeadingH1(form.i18n.Title))
	general := b.NewConfigGeneralContainer()
	server := b.NewConfigServerContainer()
	submitBox := Must(gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 8))

	submitButton := Must(gtk.ButtonNewWithLabel(form.i18n.ButtonSave))
	submitResult := Must(gtk.LabelNew(""))

	submitButton.Connect("clicked", func() {
		defer time.AfterFunc(time.Second*4, func() {
			submitResult.SetText("")
		})

		cfg := &config.Config{
			General: general.GetValues(),
			Server:  server.GetValues(),
		}

		if err := onSubmit(cfg); err != nil {
			submitResult.SetText(fmt.Sprintf(
				form.i18n.LabelError,
				err.Error(),
			))
			return
		}

		submitResult.SetText(form.i18n.LabelSuccess)
	})

	submitBox.PackStart(submitButton, false, false, 0)
	submitBox.PackStart(submitResult, false, false, 0)

	form.PackStart(heading, false, false, 0)
	form.PackStart(general, false, false, 0)
	form.PackStart(server, false, false, 0)
	form.PackStart(submitBox, false, false, 0)

	return form
}
