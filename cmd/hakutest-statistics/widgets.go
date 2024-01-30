package main

import (
	"errors"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
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
	appI18n i18n.StatsAppI18n,
) *widget.Form {
	directoryButton := chooseDirectoryButton(parent, initialPath)
	testSelect := widget.NewSelect(tests, func(_ string) {})
	formatSelect := widget.NewSelect(formats, func(_ string) {})

	testSelect.PlaceHolder = appI18n.SelectText
	formatSelect.PlaceHolder = appI18n.SelectText

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: appI18n.LabelTest, Widget: testSelect},
			{Text: appI18n.LabelFormat, Widget: formatSelect},
			{Text: appI18n.LabelDirectory, Widget: directoryButton},
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

	form.SubmitText = appI18n.SubmitText
	form.CancelText = appI18n.CancelText

	return form
}
