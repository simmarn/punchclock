package punchclock

import (
	"errors"
	"fmt"
	"time"
)

const (
	YYYYMMDD  = "2006-01-02"
	HHMMSS24h = "15:04"
)

type WorkDayModel struct {
	workday WorkDayRecord
}

func NewWorkDayModel(record WorkDayRecord) WorkDayModel {
	return WorkDayModel{record}
}

func NewWorkDayModelSlice(records []WorkDayRecord) *[]WorkDayModel {
	var modelSlice []WorkDayModel
	for _, record := range records {
		modelSlice = append(modelSlice, WorkDayModel{record})
	}
	return &modelSlice
}

func (m *WorkDayModel) Day() string {
	return m.workday.WorkDay.WorkStarted.Format(YYYYMMDD)
}

func (m *WorkDayModel) Start() string {
	return m.workday.WorkDay.WorkStarted.Format(HHMMSS24h)
}

func (m *WorkDayModel) End() string {
	return m.workday.WorkDay.WorkEnded.Format(HHMMSS24h)
}

func (m *WorkDayModel) Pause() string {
	return fmtDuration(m.workday.PauseTime)
}

func (m *WorkDayModel) WorkingDecimalFormat() string {
	return fmtDurationDecimal(m.workday.WorkingTime)
}

func (m *WorkDayModel) WorkingTimeClockFormat() string {
	return fmtDuration(m.workday.WorkingTime)
}

func (m *WorkDayModel) SetPauses(pauses []WorkPause) {
	m.workday.WorkDay.Pauses = pauses
	m.recalculate()
}

func (m *WorkDayModel) SetStart(timeStr string) error {
	time, err := UpdateTime(m.workday.WorkDay.WorkStarted, timeStr)
	if err != nil {
		return err
	}
	if time.After(m.workday.WorkDay.WorkEnded) {
		return errors.New("work start must be before work end")
	}
	m.workday.WorkDay.WorkStarted = time
	m.recalculate()
	return nil
}

func (m *WorkDayModel) SetEnd(timeStr string) error {
	time, err := UpdateTime(m.workday.WorkDay.WorkEnded, timeStr)
	if err != nil {
		return err
	}
	if time.Before(m.workday.WorkDay.WorkStarted) {
		return errors.New("work end must be after work start")
	}
	m.workday.WorkDay.WorkEnded = time
	m.recalculate()
	return nil
}

func (m *WorkDayModel) recalculate() {
	m.workday = CalculateWorkDay(m.workday.WorkDay)
}

func fmtDuration(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d:%02d", h, m)
}

func fmtDurationDecimal(d time.Duration) string {
	d = d.Round(time.Minute)
	h := d / time.Hour
	d -= h * time.Hour
	m := d / time.Minute
	return fmt.Sprintf("%02d.%02d", h, 100*m/60)
}

func UpdateTime(baseTime time.Time, s string) (time.Time, error) {
	timeString := s
	var h, m int
	_, err := fmt.Sscanf(timeString, "%d:%d", &h, &m)
	if err != nil {
		return baseTime, err
	}
	if h > 23 || m > 59 {
		return baseTime, errors.New("not valid time (hh:mm)")
	}
	newTime := time.Date(baseTime.Year(), baseTime.Month(), baseTime.Day(), h, m, 0, 0, time.Local)
	return newTime, nil
}
