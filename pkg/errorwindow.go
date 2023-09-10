package punchclock

import (
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func TerminateOnIoReadError(app fyne.App, err error) {
	if err == nil {
		return
	}

	window := app.NewWindow("Can not start " + app.Metadata().Name)

	errorInfo := container.New(layout.NewVBoxLayout(),
		widget.NewLabel("Error when reading "+FilePath),
		widget.NewLabel(err.Error()))
	exitButton := container.New(layout.NewHBoxLayout(), layout.NewSpacer(), widget.NewButton("OK", func() {
		os.Exit(1)
	}), layout.NewSpacer())
	window.SetContent(container.NewBorder(nil, exitButton, nil, nil, errorInfo))
	window.ShowAndRun()
}
