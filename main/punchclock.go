package main

import (
	"simmarn/punchclock"
)

const (
	RECORDPATH string = "timesheet.json"
)

func main() {
	fh := punchclock.NewFileHandler(RECORDPATH)
	controller := punchclock.NewPunchclockController(fh)
	view := punchclock.NewMainWindowView(controller, controller.Model)
	view.ShowMainWindow()
}
