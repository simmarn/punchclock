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
	today.WorkStarted = RoundDown(now)
	today.WorkEnded = RoundUp(now)
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
		pc.today.WorkEnded = RoundUp(now)

		if !pc.currentPause.Start.IsZero() {
			pc.currentPause.End = RoundDown(now)
			pc.today.AddPause(pc.currentPause)
			pc.currentPause = NewWorkPause()
		}

	} else {
		pc.yesterday = pc.today
		pc.today = *NewWorkDay(RoundDown(now))
	}
}

func (pc *PunchClock) Pause() {
	if pc.currentPause.Start.IsZero() {
		pc.Work()
		pc.currentPause.Start = RoundUp(time.Now())
	}
}

func (pc *PunchClock) GetCurrentWorkDay() WorkDay {
	if pc.currentPause.Start.IsZero() {
		return pc.today
	} else {
		now := time.Now()
		temp_today := pc.today
		temp_today.WorkEnded = RoundUp(now)
		temp_pause := pc.currentPause
		temp_pause.End = RoundDown(now)
		temp_today.AddPause(temp_pause)
		return temp_today
	}
}

func (pc *PunchClock) SetCurrentWorkDay(wd WorkDay) {
	if pc.today.WorkStarted.Day() == wd.WorkEnded.Day() {
		pc.today = wd
	}
}

func (pc *PunchClock) GetPreviousWorkDay() WorkDay {
	return pc.yesterday
}

func RoundDown(t time.Time) time.Time {
	return t.Truncate(5 * time.Minute)
}

func RoundUp(t time.Time) time.Time {
	return RoundDown(t).Add(5 * time.Minute)
}
