package main

import punchclock "github.com/simmarn/punchclock/pkg"

const (
	RECORDPATH string = "timesheet.json"
)

func main() {
	fh := punchclock.NewFileHandler(RECORDPATH)
	controller := punchclock.NewPunchclockController(fh)
	view := punchclock.NewMainWindowView(controller, controller.Model)
	view.ShowMainWindow()
}
