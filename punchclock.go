package main

import (
	"fyne.io/fyne/v2/app"
	punchclock "github.com/simmarn/punchclock/pkg"
)

const (
	RECORDPATH string = "timesheet.json"
)

func main() {
	app := app.NewWithID("com.github.simmarn.punchclock")
	fh := punchclock.NewFileHandler(RECORDPATH)
	prefs := punchclock.NewPrefHandler(app.Preferences())
	controller := punchclock.NewPunchclockController(fh, prefs, app)
	view := punchclock.NewMainWindowView(controller, controller.Model)
	view.ShowMainWindow()
}
