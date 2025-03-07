package test

import (
	"fmt"
	"testing"
	"time"

	punchclock "github.com/simmarn/punchclock/pkg"

	"github.com/stretchr/testify/assert"
)

func TestTimesheetModel_GetCurrentMonth(t *testing.T) {
	assert := assert.New(t)
	records := CreateRecordsForTestWithSameMonthLastYear()

	timesheet := punchclock.NewTimesheetModel(records)

	currentMonth := timesheet.GetCurrentMonth()
	assert.Equal(2, len(currentMonth))
	assert.Equal(4, len(timesheet.GetAllRecords()))
	assert.Equal(1, len(timesheet.GetPreviousMonths()))
}

func TestTimesheetModel_UpdateWorkDay_UpdateExising(t *testing.T) {
	assert := assert.New(t)
	records := CreateRecordsForTest()

	ts := punchclock.NewTimesheetModel(records)

	edited := records[1]
	worktimeBeforeEdit := ts.GetAllRecords()[1].WorkingTime
	workstart := edited.WorkDay.WorkStarted
	pauses := edited.WorkDay.Pauses
	pauses = append(pauses, punchclock.WorkPause{Start: workstart.Add(60 * time.Minute), End: workstart.Add(90 * time.Minute)})
	edited.WorkDay.Pauses = pauses
	edited = punchclock.CalculateWorkDay(edited.WorkDay)

	ts.UpdateWorkDay(edited)

	assert.Less(ts.GetAllRecords()[1].WorkingTime, worktimeBeforeEdit)
	assert.Equal(3, len(ts.GetAllRecords()))
}

func TestTimesheetModel_UpdateWorkDay_AddNew(t *testing.T) {
	assert := assert.New(t)
	records := CreateRecordsForTest()
	origanNumberOfRecords := len(records)

	ts := punchclock.NewTimesheetModel(records)

	newDay := OneWorkdayPlease(records[1].WorkDay.WorkStarted.Add(-24 * time.Hour))
	newRecord := punchclock.CalculateWorkDay(newDay)

	fmt.Println("Original record dates:")
	for _, r := range records {
		fmt.Println(r.WorkDay.WorkStarted.String())
	}
	fmt.Println("Record to add")
	fmt.Println(newRecord.WorkDay.WorkStarted.String())

	ts.UpdateWorkDay(newRecord)

	assert.Greater(len(ts.GetAllRecords()), origanNumberOfRecords)
	AreSortedByDate(assert, ts.GetAllRecords())
}

func AreSortedByDate(assert *assert.Assertions, records []punchclock.WorkDayRecord) bool {
	for i := 1; i < len(records); i++ {
		assert.Greater(records[i].WorkDay.WorkEnded, records[i-1].WorkDay.WorkEnded)
	}
	return true
}
