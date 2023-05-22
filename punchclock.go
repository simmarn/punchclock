package punchclock

import (
	"time"
)

type WorkPause struct {
	Start time.Time
	End   time.Time
}

func NewWorkPause() WorkPause {
	return WorkPause{}
}

type WorkDay struct {
	WorkStarted time.Time
	WorkEnded   time.Time
	Pauses      []WorkPause
}

func NewWorkDay(startTime time.Time) *WorkDay {
	day := new(WorkDay)
	day.WorkStarted = startTime
	day.WorkEnded = startTime
	day.Pauses = make([]WorkPause, 0)
	return day
}

type PunchClock struct {
	today        WorkDay
	yesterday    WorkDay
	currentPause WorkPause
}

func NewPunchClock() *PunchClock {
	now := time.Now()
	pc := new(PunchClock)
	today := new(WorkDay)
	today.WorkStarted = now
	today.WorkEnded = now
	pauseSlice := make([]WorkPause, 0)
	today.Pauses = pauseSlice
	pc.today = *today
	return pc
}

func NewPunchClockFromData(today WorkDay) *PunchClock {
	pc := new(PunchClock)
	pc.today = today
	return pc
}

func (pc *PunchClock) Work() {
	now := time.Now()
	if now.Day() == pc.today.WorkStarted.Day() {
		pc.today.WorkEnded = now

		if !pc.currentPause.Start.IsZero() {
			pc.currentPause.End = now
			pc.today.Pauses = append(pc.today.Pauses, pc.currentPause)
			pc.currentPause = NewWorkPause()
		}

	} else {
		pc.yesterday = pc.today
		pc.today = *NewWorkDay(now)
	}
}

func (pc *PunchClock) Pause() {
	if pc.currentPause.Start.IsZero() {
		pc.Work()
		pc.currentPause.Start = time.Now()
	}
}

func (pc *PunchClock) GetCurrentWorkDay() WorkDay {
	return pc.today
}

func (pc *PunchClock) GetPreviousWorkDay() WorkDay {
	return pc.yesterday
}
