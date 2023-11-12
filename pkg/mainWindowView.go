package punchclock

import (
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

const (
	COLUMNS              = 5
	PREFAPPWIDTH  string = "ApplicationWidth"
	PREFAPPHEIGHT string = "ApplicationHeight"
)

type MainWindowView struct {
	controller *PunchclockController
	model      *PunchclockModel
	mainWindow fyne.Window
	table      *widget.Table
}

func NewMainWindowView(c *PunchclockController, m *PunchclockModel) *MainWindowView {
	myWindow := c.App.NewWindow("Punchclock")
	width := c.prefs.GetFloatWithFallback(PREFAPPWIDTH, COLUMNS*99)
	height := c.prefs.GetFloatWithFallback(PREFAPPHEIGHT, 600)
	myWindow.Resize(fyne.NewSize(float32(width), float32(height)))
	myWindow.SetMainMenu(SetMainMenu(c, myWindow))
	v := MainWindowView{}
	headerLabel := widget.NewLabel("Punchclock Timesheet")
	selectTimesheet := widget.NewButton(PreviousMonth.toString(), nil)
	header := container.NewHBox(headerLabel, layout.NewSpacer(), selectTimesheet)

	table := widget.NewTableWithHeaders(
		func() (int, int) {
			return len(m.SelectedMonth), COLUMNS
		},
		func() fyne.CanvasObject {
			label := NewTappableLabel("placeholder", nil)
			label.Alignment = fyne.TextAlignCenter
			return label
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			l := co.(*TappableLabel)
			row := tci.Row

			switch tci.Col {
			case 0:
				l.SetText(m.SelectedMonth[row].Day())
			case 1:
				l.SetText(m.SelectedMonth[row].Start())
				l.OnTapped = func() {
					e := widget.NewEntry()
					e.Text = l.Text
					e.Validator = validation.NewTime(HHMMSS24h)
					formItem := widget.NewFormItem("Arrived", e)
					dialog.ShowForm(m.SelectedMonth[row].Day(), "Submit", "Cancel", []*widget.FormItem{formItem},
						func(b bool) {
							if b && e.Validate() == nil {
								m.SelectedMonth[row].SetStart(e.Text)
								l.SetText(m.SelectedMonth[row].Start())
								c.Update(m.SelectedMonth[row].workday)
								v.refresh()
							}
						},
						myWindow)
				}
			case 2:
				l.SetText(m.SelectedMonth[row].End())
				l.OnTapped = func() {
					e := widget.NewEntry()
					e.Text = l.Text
					e.Validator = validation.NewTime(HHMMSS24h)
					formItem := widget.NewFormItem("Left", e)
					dialog.ShowForm(m.SelectedMonth[row].Day(), "Submit", "Cancel", []*widget.FormItem{formItem},
						func(b bool) {
							if b && e.Validate() == nil {
								m.SelectedMonth[row].SetEnd(e.Text)
								l.SetText(m.SelectedMonth[row].End())
								c.Update(m.SelectedMonth[row].workday)
								v.refresh()
							}
						},
						myWindow)
				}
			case 3:
				l.SetText(m.SelectedMonth[row].Pause())
				l.OnTapped = func() {
					p := new(EditBreakTimeView)
					p.OnSubmitted = func() {
						m.SelectedMonth[row].SetPauses(p.Pauses)
						l.SetText(m.SelectedMonth[row].Pause())
						c.Update(m.SelectedMonth[row].workday)
						v.refresh()
					}
					p.Show(&myWindow, m.SelectedMonth[row].workday.WorkDay)
				}
			case 4:
				l.SetText(m.SelectedMonth[row].WorkingTime())
			}
		})
	table.SetColumnWidth(5, 30)
	table.CreateHeader = func() fyne.CanvasObject {
		return widget.NewLabelWithStyle("", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	}
	table.UpdateHeader = func(id widget.TableCellID, template fyne.CanvasObject) {
		if id.Row == -1 {
			switch id.Col {
			case 0:
				template.(*widget.Label).SetText("Date")
			case 1:
				template.(*widget.Label).SetText("Arrived")
			case 2:
				template.(*widget.Label).SetText("Left")
			case 3:
				template.(*widget.Label).SetText("Break Time")
			case 4:
				template.(*widget.Label).SetText("Work Time")
			}
		}
		if id.Col == -1 {
			template.(*widget.Label).SetText(strconv.Itoa(id.Row + 1))
		}
	}
	table.ShowHeaderColumn = false
	refreshButton := widget.NewButton("Work", nil)
	pauseButton := widget.NewButton("Pause", nil)
	statusLabel := widget.NewLabelWithData(c.Status)
	statusLabel.TextStyle.Bold = true
	buttonContainer := container.NewHBox(refreshButton, pauseButton, layout.NewSpacer(), statusLabel)

	myWindow.SetContent(container.NewBorder(
		header,
		buttonContainer,
		nil,
		nil,
		table))
	v.controller = c
	v.model = m
	v.mainWindow = myWindow
	v.table = table

	selectTimesheet.OnTapped = func() {
		table.ScrollToTop()
		selectTimesheet.Text = c.ToggleSelectedMonth()
		selectTimesheet.Refresh()
		v.refresh()
	}

	refreshButton.OnTapped = func() {
		c.Present()
		v.refresh()
	}
	pauseButton.OnTapped = func() {
		c.Pause()
		v.refresh()
	}

	v.mainWindow.SetOnClosed(v.onClosed)

	go func() {
		interval := 5 * time.Minute
		startloop := GetNow().Truncate(interval).Add(interval + time.Second)
		c.Refresh()
		v.refresh()
		Sleep(time.Until(startloop))
		for range time.Tick(interval) {
			c.Refresh()
			v.refresh()
		}
	}()

	return &v
}

func (v *MainWindowView) ShowMainWindow() {
	v.refresh()
	v.mainWindow.ShowAndRun()
}

func (v *MainWindowView) refresh() {
	v.controller.Refresh()
	v.table.Refresh()
}

func (v *MainWindowView) onClosed() {
	Log.Info().Msg("Closing " + v.controller.App.UniqueID())
	v.controller.prefs.SetFloat(PREFAPPWIDTH, float64(v.mainWindow.Canvas().Size().Width))
	v.controller.prefs.SetFloat(PREFAPPHEIGHT, float64(v.mainWindow.Canvas().Size().Height))
	v.refresh()
}
