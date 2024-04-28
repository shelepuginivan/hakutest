package main

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/hakutest/internal/app/desktop"
	"github.com/shelepuginivan/hakutest/internal/app/desktop/layouts"
	"github.com/shelepuginivan/hakutest/internal/pkg/application"
	"github.com/shelepuginivan/hakutest/internal/pkg/config"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/statistics"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

const (
	AppId = "com.github.shelepuginivan.hakutest"
	Title = "Hakutest"
)

// setupWindow configures the main window of the application.
func setupWindow(app *application.App, g *gtk.Application) *gtk.ApplicationWindow {
	win := desktop.Must(gtk.ApplicationWindowNew(g))
	win.SetTitle(Title)
	win.SetDefaultSize(800, 600)

	builder := desktop.NewBuilder(app)
	configService := config.NewService()
	resultsService := results.NewService(app)
	statisticsService := statistics.NewService(app, resultsService)
	testService := test.NewService(app)

	serverForm := desktop.Must(layouts.NewScrolled(builder.NewServerForm()))

	editor := builder.NewEditor(
		testService.GetTestList(),
		testService.GetTestByName,
		func(t *test.Test) error {
			return testService.SaveToTestsDirectory(t, t.Title)
		},
	)

	statsForm := desktop.Must(layouts.NewScrolled(builder.NewStatsForm(
		resultsService.GetResultsList(),
		[]string{statistics.FormatExcel, statistics.FormatImage},
		func(result, format string) error {
			dest := filepath.Join(xdg.UserDirs.Download, result)
			return statisticsService.Export(result, dest, format)
		},
	)))

	configForm := desktop.Must(layouts.NewScrolled(builder.NewConfigForm(configService.WriteConfig)))

	nb := desktop.Must(layouts.NewNotebook(
		layouts.NotebookPage{
			Child: serverForm,
			Label: desktop.Must(gtk.LabelNew(app.I18n.Gtk.Server.NotebookLabel)),
		},
		layouts.NotebookPage{
			Child: editor,
			Label: desktop.Must(gtk.LabelNew(app.I18n.Gtk.Editor.NotebookLabel)),
		},
		layouts.NotebookPage{
			Child: statsForm,
			Label: desktop.Must(gtk.LabelNew(app.I18n.Gtk.Stats.NotebookLabel)),
		},
		layouts.NotebookPage{
			Child: configForm,
			Label: desktop.Must(gtk.LabelNew(app.I18n.Gtk.Config.NotebookLabel)),
		},
	))

	win.Add(nb)

	return win
}

// activate is a callback to the Gtk.Application `activate` signal.
func activate(g *gtk.Application) {
	app := application.New()
	win := setupWindow(app, g)
	win.ShowAll()
}

func main() {
	g := desktop.Must(gtk.ApplicationNew(AppId, glib.APPLICATION_FLAGS_NONE))
	g.Connect("activate", activate)
	g.Run(os.Args)
}
