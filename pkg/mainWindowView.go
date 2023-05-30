package punchclock

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	COLUMNS = 6
)

type MainWindowView struct {
	controller     *PunchclockController
	model          *PunchclockModel
	mainWindow     fyne.Window
	table          fyne.Widget
	hiddenObjects  []fyne.CanvasObject // TODO write bug on show/hide setting on windows resize
	visibleObjects []fyne.CanvasObject
}

func NewMainWindowView(c *PunchclockController, m *PunchclockModel) *MainWindowView {
	myApp := app.New()
	myWindow := myApp.NewWindow("Punchclock")
	headerLabel := widget.NewLabel("Punchclock Timesheet")
	var visibleObjects []fyne.CanvasObject
	var hiddenObjects []fyne.CanvasObject

	table := widget.NewTable(
		func() (int, int) {
			return len(m.CurrentMonth) + 1, COLUMNS
		},
		func() fyne.CanvasObject {
			label := widget.NewLabel("placeholder")
			label.Alignment = fyne.TextAlignCenter
			button := widget.NewButtonWithIcon("", theme.DocumentCreateIcon(), nil)
			button.Hide()
			return container.NewMax(label, button)
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			l := co.(*fyne.Container).Objects[0].(*widget.Label)
			b := co.(*fyne.Container).Objects[1].(*widget.Button)
			if tci.Row == 0 {
				switch tci.Col {
				case 0:
					l.SetText("Date")
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 1:
					l.SetText("Arrived")
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 2:
					l.SetText("Left")
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 3:
					l.SetText("Break Time")
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 4:
					l.SetText("Work Time")
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 5:
					l.Hide()
					hiddenObjects = append(hiddenObjects, b)
					hiddenObjects = append(hiddenObjects, b)
				}
				l.TextStyle.Bold = true
			} else {
				row := tci.Row - 1

				switch tci.Col {
				case 0:
					l.SetText(m.CurrentMonth[row].Day())
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 1:
					l.SetText(m.CurrentMonth[row].Start())
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 2:
					l.SetText(m.CurrentMonth[row].End())
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 3:
					l.SetText(m.CurrentMonth[row].Pause())
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 4:
					l.SetText(m.CurrentMonth[row].WorkingTime())
					visibleObjects = append(visibleObjects, l)
					hiddenObjects = append(hiddenObjects, b)
				case 5:
					l.Hide()
					b.Show()
					visibleObjects = append(visibleObjects, b)
					hiddenObjects = append(hiddenObjects, l)
				}
			}
		})
	table.SetColumnWidth(5, 30)
	scrollableContent := container.NewScroll(table)
	scrollableContent.SetMinSize(fyne.NewSize((COLUMNS-1)*94+40, 300))
	refreshButton := widget.NewButton("Work", nil)
	pauseButton := widget.NewButton("Pause", nil)
	status := binding.NewString()
	status.Set(c.Status)
	statusLabel := widget.NewLabelWithData(status)
	statusLabel.TextStyle.Bold = true
	buttonContainer := container.NewHBox(refreshButton, pauseButton, layout.NewSpacer(), statusLabel)

	myWindow.SetContent(container.New(layout.NewVBoxLayout(),
		headerLabel,
		widget.NewSeparator(),
		scrollableContent,
		widget.NewSeparator(),
		buttonContainer))
	v := MainWindowView{c, m, myWindow, table, hiddenObjects, visibleObjects}

	refreshButton.OnTapped = func() {
		c.Present()
		v.refresh()
		status.Set(c.Status)
	}
	pauseButton.OnTapped = func() {
		c.Pause()
		v.refresh()
		status.Set(c.Status)
	}

	v.mainWindow.SetOnClosed(v.refresh)

	go func() {
		for range time.Tick(5 * time.Minute) {
			c.Refresh()
			v.refresh()
		}
	}()

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

	for _, o := range v.hiddenObjects {
		o.Hide()
	}
	for _, o := range v.visibleObjects {
		o.Show()
	}
}
