package punchclock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainWindowView struct {
	controller *PunchclockController
	model      *PunchclockModel
	mainWindow fyne.Window
}

func NewMainWindowView(c *PunchclockController, m *PunchclockModel) *MainWindowView {
	myApp := app.New()
	myWindow := myApp.NewWindow("Punchclock")
	headerLabel := widget.NewLabel("Current working time")

	contentContainer := widget.NewTable(
		func() (int, int) {
			return len(m.CurrentMonth), 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("placeholder")
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {

			var text string
			switch tci.Col {
			case 0:
				text = m.CurrentMonth[tci.Row].Day()
			case 1:
				text = m.CurrentMonth[tci.Row].Start()
			case 2:
				text = m.CurrentMonth[tci.Row].End()
			case 3:
				text = m.CurrentMonth[tci.Row].Pause()
			case 4:
				text = m.CurrentMonth[tci.Row].WorkingTime()
			}
			co.(*widget.Label).SetText(text)
		})
	scrollableContent := container.NewScroll(contentContainer)
	scrollableContent.SetMinSize(fyne.NewSize(1200, 400))

	myWindow.SetContent(container.New(layout.NewVBoxLayout(),
		headerLabel,
		widget.NewSeparator(),
		scrollableContent))
	v := MainWindowView{c, m, myWindow}

	return &v
}

func (v MainWindowView) ShowMainWindow() {
	v.mainWindow.ShowAndRun()
}
