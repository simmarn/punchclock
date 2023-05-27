package punchclock_test

import (
	"simmarn/punchclock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimesheetModel_GetCurrentMonth(t *testing.T) {
	assert := assert.New(t)
	records := CreateRecordsForTest()

	timesheet := punchclock.NewTimesheetModel(records)

	currentMonth := timesheet.GetCurrentMonth()
	assert.Equal(2, len(currentMonth))
	assert.Equal(3, len(timesheet.GetAllRecords()))
}

func TestTimesheetModel_UpdateWorkDay(t *testing.T) {
	assert := assert.New(t)
	records := CreateRecordsForTest()

	ts := punchclock.NewTimesheetModel(records)

	edited := records[1]
	worktimeBeforeEdit := ts.GetAllRecords()[1].WorkingTime
	workstart := edited.WorkDay.WorkStarted
	pauses := edited.WorkDay.Pauses
	pauses = append(pauses, punchclock.WorkPause{workstart.Add(60 * time.Minute), workstart.Add(90 * time.Minute)})
	edited.WorkDay.Pauses = pauses
	edited = punchclock.CalculateWorkDay(edited.WorkDay)

	ts.UpdateWorkDay(edited)

	assert.Less(ts.GetAllRecords()[1].WorkingTime, worktimeBeforeEdit)
	assert.Equal(3, len(ts.GetAllRecords()))
}
