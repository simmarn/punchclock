package punchclock

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/widget"
)

const (
	emptyEntry = "12:00"
)

type ValidatedTimeEntry struct {
	widget.Entry
}

func NewValidatedTimeEntry() *ValidatedTimeEntry {
	entry := &ValidatedTimeEntry{}
	entry.ExtendBaseWidget(entry)
	entry.Validator = validation.NewTime(HHMMSS24h)
	entry.SetText(emptyEntry)
	return entry
}

func (e *ValidatedTimeEntry) Clear() {
	e.SetText(emptyEntry)
}

type TappableLabel struct {
	widget.Label
	OnTapped func() `json:"-"`
}

func NewTappableLabel(text string, tapped func()) *TappableLabel {
	label := &TappableLabel{
		OnTapped: tapped,
	}
	label.ExtendBaseWidget(label)
	label.SetText(text)
	return label
}

func (t *TappableLabel) Tapped(*fyne.PointEvent) {
	if t.OnTapped != nil {
		t.OnTapped()
	}
}

func (t *TappableLabel) TappedSecondary(*fyne.PointEvent) {}
