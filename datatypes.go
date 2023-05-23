package punchclock

import "time"

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

type WorkDayRecord struct {
	WorkDay     WorkDay
	WorkingTime time.Duration
	PauseTime   time.Duration
}

func CalculateWorkDay(workday WorkDay) WorkDayRecord {
	record := WorkDayRecord{}
	record.WorkDay = workday
	for _, pause := range workday.Pauses {
		record.PauseTime = record.PauseTime + (pause.End.Sub(pause.Start))
	}
	record.WorkingTime = workday.WorkEnded.Sub(record.WorkDay.WorkStarted) - record.PauseTime
	return record
}
