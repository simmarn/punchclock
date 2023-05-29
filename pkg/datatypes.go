package punchclock

import (
	"time"
)

type WorkPause struct {
	Start time.Time `json:"Start"`
	End   time.Time `json:"End"`
}

func NewWorkPause() WorkPause {
	return WorkPause{}
}

type WorkDay struct {
	WorkStarted time.Time   `json:"WorkStarted"`
	WorkEnded   time.Time   `json:"WorkEnded"`
	Pauses      []WorkPause `json:"Pauses"`
}

func NewWorkDay(startTime time.Time) *WorkDay {
	day := new(WorkDay)
	day.WorkStarted = startTime
	day.WorkEnded = startTime
	day.Pauses = make([]WorkPause, 0)
	return day
}

func (d *WorkDay) AddPause(p WorkPause) {
	if p.End.Sub(p.Start) > 0 {
		d.Pauses = append(d.Pauses, p)
	}
}

type WorkDayRecord struct {
	WorkDay     WorkDay       `json:"WorkDay"`
	WorkingTime time.Duration `json:"WorkingTime"`
	PauseTime   time.Duration `json:"PauseTime"`
}
