package punchclock

const (
	YYYYMMDD  = "2006-01-02"
	HHMMSS24h = "15:04:05"
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
	return m.workday.PauseTime.String()
}

func (m *WorkDayModel) WorkingTime() string {
	return m.workday.WorkingTime.String()
}
