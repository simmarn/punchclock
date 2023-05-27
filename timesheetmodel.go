package punchclock

import "time"

type TimesheetModel struct {
	records []WorkDayRecord
}

func NewTimesheetModel(records []WorkDayRecord) *TimesheetModel {
	var model = new(TimesheetModel)
	model.records = records
	return model
}

func (ts *TimesheetModel) GetCurrentMonth() (selected []WorkDayRecord) {
	currentMonth := time.Now().Month()
	for _, record := range ts.records {
		if record.WorkDay.WorkStarted.Month() == currentMonth {
			selected = append(selected, record)
		}
	}
	return selected
}

func (ts *TimesheetModel) GetAllRecords() []WorkDayRecord {
	return ts.records
}

func (ts *TimesheetModel) UpdateWorkDay(updated WorkDayRecord) {
	index := ts.GetIndexOf(updated)
	ts.records[index] = updated
}

func (ts *TimesheetModel) GetIndexOf(day WorkDayRecord) int {
	startTime := day.WorkDay.WorkStarted
	for index, record := range ts.records {
		if IsSameDay(record.WorkDay.WorkStarted, startTime) {
			return index
		}
	}
	return -1
}

func IsSameDay(t1 time.Time, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.Month() == t2.Month() && t1.Day() == t2.Day()
}
