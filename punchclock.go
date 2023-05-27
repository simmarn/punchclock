package punchclock

import (
	"time"
)

func CalculateWorkDay(workday WorkDay) WorkDayRecord {
	record := WorkDayRecord{}
	record.WorkDay = workday
	for _, pause := range workday.Pauses {
		record.PauseTime = record.PauseTime + (pause.End.Sub(pause.Start))
	}
	record.WorkingTime = workday.WorkEnded.Sub(record.WorkDay.WorkStarted) - record.PauseTime
	return record
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
