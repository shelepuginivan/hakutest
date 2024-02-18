package main

import (
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"github.com/shelepuginivan/hakutest/internal/app/desktop"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/statistics"
)

const (
	appID   = "com.github.shelepuginivan.Hakutest.Statistics"
	appName = "Hakutest Statistics"
)

func main() {
	a := app.NewWithID(appID)
	w := a.NewWindow(appName)

	appI18n := i18n.New().Statistics.App
	resultsService := results.NewService()
	statsService := statistics.NewService(resultsService)

	formats := []string{statistics.FormatExcel, statistics.FormatImage}

	exportDir, err := os.Getwd()

	if err != nil {
		exportDir = ""
	}

	headerLabel := canvas.NewText(appName, color.Black)
	headerLabel.TextSize = 36
	headerLabel.Alignment = fyne.TextAlignCenter

	form := desktop.NewStatsExportForm(
		w,
		resultsService.GetResultsList(),
		formats,
		exportDir,
		statsService.Export,
		func() {
			a.SendNotification(fyne.NewNotification(
				appName,
				appI18n.SuccessText,
			))
		},
		func(err error) {
			a.SendNotification(fyne.NewNotification(
				appName,
				fmt.Sprintf("%s %v", appI18n.ErrorPrefix, err),
			))
		},
		appI18n,
	)

	content := container.New(
		desktop.NewPaddedLayout(96),
		headerLabel,
		form,
	)

	w.SetContent(content)
	w.ShowAndRun()
}
