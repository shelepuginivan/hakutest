package main

import (
	"fmt"
	"image/color"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/shelepuginivan/hakutest/internal/pkg/results"
	"github.com/shelepuginivan/hakutest/internal/pkg/statistics"
	"github.com/shelepuginivan/hakutest/internal/pkg/test"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hakutest Statistics")

	exportDir, err := os.Getwd()

	if err != nil {
		exportDir = ""
	}

	headerLabel := canvas.NewText("Hakutest Statistics", color.Black)
	headerLabel.TextSize = 36
	headerLabel.Alignment = fyne.TextAlignCenter

	statusLabel := widget.NewLabel("")

	testService := test.NewService()
	resultsService := results.NewService()

	formats := []string{statistics.FormatExcel, statistics.FormatImage}

	directoryButton := chooseDirectoryButton(w, exportDir)
	testSelect := widget.NewSelect(testService.GetTestList(), func(_ string) {})
	formatSelect := widget.NewSelect(formats, func(_ string) {})

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Test", Widget: testSelect},
			{Text: "Format", Widget: formatSelect},
			{Text: "Export to", Widget: directoryButton},
		},
		OnSubmit: func() {
			testName := testSelect.Selected
			format := formatSelect.Selected

			if testName == "" || format == "" {
				statusLabel.SetText("Please select test and format")
				return
			}

			res, err := resultsService.GetResultsOfTest(testName)

			if err != nil {
				statusLabel.SetText(errorLabel(err))
				return
			}

			stats := statistics.New(res)
			exportPath := filepath.Join(directoryButton.Text, testName)

			if err := stats.Export(exportPath, formatSelect.Selected); err != nil {
				statusLabel.SetText(errorLabel(err))
				return
			}

			statusLabel.SetText(fmt.Sprintf("Successfully exported to: %s", exportPath))
		},
		OnCancel: w.Close,
	}

	w.SetContent(container.NewVBox(
		layout.NewSpacer(),
		headerLabel,
		layout.NewSpacer(),
		form,
		layout.NewSpacer(),
		container.NewCenter(statusLabel),
		layout.NewSpacer(),
	))
	w.ShowAndRun()
}

func chooseDirectoryButton(parent fyne.Window, initialPath string) *widget.Button {
	button := widget.NewButton(initialPath, func() {})

	directoryDialog := dialog.NewFolderOpen(func(lu fyne.ListableURI, err error) {
		if err != nil || lu == nil {
			return
		}

		button.SetText(lu.Path())
	}, parent)

	button.OnTapped = directoryDialog.Show
	return button
}

func errorLabel(err error) string {
	return fmt.Sprintf("Error: %s", err.Error())
}
