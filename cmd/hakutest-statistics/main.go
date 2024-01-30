package main

import (
	"fmt"
	"image/color"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/statistics"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

const appID = "com.github.shelepuginivan.Hakutest.Statistics"

func main() {
	a := app.NewWithID(appID)
	w := a.NewWindow("Hakutest Statistics")

	testService := test.NewService()
	statsService := statistics.NewService(results.NewService())

	formats := []string{statistics.FormatExcel, statistics.FormatImage}

	exportDir, err := os.Getwd()

	if err != nil {
		exportDir = ""
	}

	headerLabel := canvas.NewText("Hakutest Statistics", color.Black)
	headerLabel.TextSize = 36
	headerLabel.Alignment = fyne.TextAlignCenter

	form := statsExportForm(
		w,
		testService.GetTestList(),
		formats,
		exportDir,
		statsService.Export,
		func() {
			a.SendNotification(fyne.NewNotification(
				"Hakutest Statistics",
				"Exported statistics successfully",
			))
		},
		func(err error) {
			a.SendNotification(fyne.NewNotification(
				"Hakutest Statistics",
				fmt.Sprintf("Error occurred: %s", err.Error()),
			))
		},
	)

	w.SetContent(container.NewVBox(
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
		headerLabel,
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		layout.NewSpacer(),
		layout.NewSpacer(),
	))
	w.ShowAndRun()
}
