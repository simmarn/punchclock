package punchclock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type EditBreakTimeView struct {
	Pauses      []WorkPause
	OnSubmitted func()
	list        *widget.List
}

func (v *EditBreakTimeView) Show(parent *fyne.Window, day WorkDay) {
	editor := NewBreakTimeEditor(&day)

	list := widget.NewList(
		func() int {
			return len(editor.GetPauses())
		},
		func() fyne.CanvasObject {
			start := widget.NewEntry()
			end := widget.NewEntry()
			start.Validator = validation.NewTime(HHMMSS24h)
			end.Validator = validation.NewTime(HHMMSS24h)
			time := widget.NewLabel("0:55")
			delete := widget.NewButtonWithIcon("", theme.DeleteIcon(), nil)
			co := container.New(layout.NewGridLayout(4), start, end, time, delete)
			return co
		},
		func(i widget.ListItemID, co fyne.CanvasObject) {
			start := co.(*fyne.Container).Objects[0].(*widget.Entry)
			end := co.(*fyne.Container).Objects[1].(*widget.Entry)
			time := co.(*fyne.Container).Objects[2].(*widget.Label)
			delete := co.(*fyne.Container).Objects[3].(*widget.Button)
			start.SetText(editor.GetStart(i))
			end.SetText(editor.GetEnd(i))
			updateOnValidation := func(err error) {
				if err == nil {
					if start.Validate() == nil && end.Validate() == nil {
						editor.Edit(i, start.Text, end.Text)
						time.SetText(editor.GetPauseTime(i))
					}
				}
			}
			start.SetOnValidationChanged(updateOnValidation)
			end.SetOnValidationChanged(updateOnValidation)
			time.SetText(editor.GetPauseTime(i))
			delete.OnTapped = func() {
				editor.Delete(i)
				v.refresh()
			}
		})
	v.list = list
	scroll := container.NewScroll(list)
	scroll.SetMinSize(fyne.NewSize(400, 140))
	newStart := NewValidatedTimeEntry()
	newEnd := NewValidatedTimeEntry()
	add := widget.NewButtonWithIcon("Add", theme.ContentAddIcon(), func() {
		err := editor.Add(newStart.Text, newEnd.Text)
		if err == nil {
			v.refresh()
			newStart.Clear()
			newEnd.Clear()
		}
	})
	addList := container.New(layout.NewHBoxLayout(), widget.NewLabel("New pause"), newStart, newEnd, add)
	dialogGui := container.NewBorder(widget.NewLabel("Pauses"), addList, nil, nil, scroll)

	dialog.ShowCustomConfirm("Edit break time", "Submit", "Cancel", dialogGui, func(b bool) {
		if b {
			v.Pauses = editor.GetPauses()
			v.OnSubmitted()
		}
	}, *parent)
}

func (v *EditBreakTimeView) refresh() {
	v.list.Refresh()
}
