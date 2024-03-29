package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/simmarn/punchclock/logging"
	punchclock "github.com/simmarn/punchclock/pkg"
)

const (
	RECORDPATH string = "timesheet.json"
)

func main() {
	punchclock.Log = logging.Configure(logging.Config{
		FileLoggingEnabled:    false,
		EncodeLogsAsJson:      true,
		ConsoleLoggingEnabled: false,
		Directory:             ".",
		Filename:              "punchclock.log",
		MaxSize:               1,
		MaxBackups:            3,
		MaxAge:                30})

	app := app.NewWithID("com.github.simmarn.punchclock")
	fh := punchclock.NewFileHandler(RECORDPATH)
	prefs := punchclock.NewPreferencesWrapper(app.Preferences())
	controller := punchclock.NewPunchclockController(fh, prefs, app)
	view := punchclock.NewMainWindowView(controller, controller.Model)
	punchclock.Log.Info().Msg("Starting " + app.UniqueID())
	view.ShowMainWindow()
}
