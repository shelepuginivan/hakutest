package main

import (
	"errors"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

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

func statsExportForm(
	parent fyne.Window,
	tests []string,
	formats []string,
	initialPath string,
	exportFunc func(testName, dest, format string) error,
	onSuccess func(),
	onError func(err error),
) *widget.Form {
	directoryButton := chooseDirectoryButton(parent, initialPath)
	testSelect := widget.NewSelect(tests, func(_ string) {})
	formatSelect := widget.NewSelect(formats, func(_ string) {})

	return &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Test", Widget: testSelect},
			{Text: "Format", Widget: formatSelect},
			{Text: "Export to", Widget: directoryButton},
		},
		OnSubmit: func() {
			testName := testSelect.Selected
			format := formatSelect.Selected
			dest := filepath.Join(directoryButton.Text, testName)

			if testName == "" {
				onError(errors.New("test is not chosen"))
				return
			}

			if format == "" {
				onError(errors.New("format is not chosen"))
				return
			}

			err := exportFunc(testName, dest, format)

			if err != nil {
				onError(err)
				return
			}

			onSuccess()
		},
		OnCancel: parent.Close,
	}
}
