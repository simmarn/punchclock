package punchclock

import (
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

func (m *WorkDayModel) WorkingTime() string {
	return fmtDurationDecimal(m.workday.WorkingTime)
	//return m.workday.WorkingTime.String()
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
	return fmt.Sprintf("%02d,%02d", h, 100*m/60)
}
