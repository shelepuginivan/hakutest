package main

import (
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
