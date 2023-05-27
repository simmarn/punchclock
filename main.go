package main

import (
	"fmt"
	"os"
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

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
