package punchclock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	COLUMNS = 5
)

type MainWindowView struct {
	controller *PunchclockController
	model      *PunchclockModel
	mainWindow fyne.Window
	table      fyne.Widget
}

func NewMainWindowView(c *PunchclockController, m *PunchclockModel) *MainWindowView {
	myApp := app.New()
	myWindow := myApp.NewWindow("Punchclock")
	headerLabel := widget.NewLabel("Punchclock Timesheet")

	table := widget.NewTable(
		func() (int, int) {
			return len(m.CurrentMonth) + 1, COLUMNS
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("placeholder")
			label.Alignment = fyne.TextAlignCenter
			return label
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			var text string
			if tci.Row == 0 {
				switch tci.Col {
				case 0:
					text = "Date"
				case 1:
					text = "Arrived"
				case 2:
					text = "Left"
				case 3:
					text = "Break Time"
				case 4:
					text = "Work Time"
				}
				co.(*widget.Label).TextStyle.Bold = true
			} else {
				row := tci.Row - 1

				switch tci.Col {
				case 0:
					text = m.CurrentMonth[row].Day()
				case 1:
					text = m.CurrentMonth[row].Start()
				case 2:
					text = m.CurrentMonth[row].End()
				case 3:
					text = m.CurrentMonth[row].Pause()
				case 4:
					text = m.CurrentMonth[row].WorkingTime()
				}
			}
			co.(*widget.Label).SetText(text)
		})
	scrollableContent := container.NewScroll(table)
	scrollableContent.SetMinSize(fyne.NewSize(COLUMNS*94, 300))
	refreshButton := widget.NewButton("Work", nil)
	pauseButton := widget.NewButton("Pause", nil)
	buttonContainer := container.NewHBox(refreshButton, pauseButton)

	myWindow.SetContent(container.New(layout.NewVBoxLayout(),
		headerLabel,
		widget.NewSeparator(),
		scrollableContent,
		widget.NewSeparator(),
		buttonContainer))
	v := MainWindowView{c, m, myWindow, table}

	refreshButton.OnTapped = func() {
		c.Present()
		v.refresh()
	}
	pauseButton.OnTapped = func() {
		c.Pause()
		v.refresh()
	}
	return &v
}

func (v *MainWindowView) ShowMainWindow() {
	v.controller.Present()
	v.refresh()
	v.mainWindow.ShowAndRun()
}

func (v *MainWindowView) refresh() {
	v.controller.Refresh()
	v.table.Refresh()
}
