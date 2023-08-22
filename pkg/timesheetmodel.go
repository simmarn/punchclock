package punchclock

import (
	"sort"
	"time"
)

type TimesheetModel struct {
	records []WorkDayRecord
}

func NewTimesheetModel(records []WorkDayRecord) *TimesheetModel {
	var model = new(TimesheetModel)
	model.records = records
	return model
}

func (ts *TimesheetModel) GetToday() WorkDay {
	if len(ts.records) > 0 {
		if IsSameDay(GetNow(), ts.records[len(ts.records)-1].WorkDay.WorkStarted) {
			return ts.records[len(ts.records)-1].WorkDay
		}
	}
	return *NewWorkDay(GetNow())
}

func (ts *TimesheetModel) GetCurrentMonth() (selected []WorkDayRecord) {
	currentMonth := GetNow().Month()
	for _, record := range ts.records {
		if record.WorkDay.WorkStarted.Month() == currentMonth {
			selected = append(selected, record)
		}
	}
	return selected
}

func (ts *TimesheetModel) GetPreviousMonths() (selected []WorkDayRecord) {
	currentMonth := GetNow().Month()
	for _, record := range ts.records {
		if record.WorkDay.WorkStarted.Month() != currentMonth {
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
	if index > -1 {
		ts.records[index] = updated
	} else {
		ts.records = append(ts.records, updated)
		ts.SortRecords()
	}
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

func (ts *TimesheetModel) SortRecords() {
	sort.Slice(ts.records, func(i, j int) bool {
		return ts.records[i].WorkDay.WorkStarted.Before(ts.records[j].WorkDay.WorkStarted)
	})
}
