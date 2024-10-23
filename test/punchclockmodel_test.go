package test

import (
	"testing"

	punchclock "github.com/simmarn/punchclock/pkg"
	"github.com/stretchr/testify/assert"
)

func TestToggleTimeFormat(t *testing.T) {
	assert := assert.New(t)

	records := CreateRecordsForTest()
	model := new(punchclock.PunchclockModel)
	workdayModelSlice := punchclock.NewWorkDayModelSlice(records)

	model.SelectedMonth = *workdayModelSlice

	assert.Equal("0.00", model.GetWorkingTime(-1))
	assert.Equal("07.75", model.GetWorkingTime(0))
	assert.Equal("07.75", model.GetWorkingTime(1))
	assert.Equal("07.75", model.GetWorkingTime(2))
	assert.Equal("0.00", model.GetWorkingTime(3))

	model.ToggleTimeFormat()
	assert.Equal("0:00", model.GetWorkingTime(-1))
	assert.Equal("07:45", model.GetWorkingTime(0))
	assert.Equal("07:45", model.GetWorkingTime(1))
	assert.Equal("07:45", model.GetWorkingTime(2))
	assert.Equal("0:00", model.GetWorkingTime(3))

	model.ToggleTimeFormat()
	assert.Equal("0.00", model.GetWorkingTime(-1))
	assert.Equal("07.75", model.GetWorkingTime(0))
	assert.Equal("07.75", model.GetWorkingTime(1))
	assert.Equal("07.75", model.GetWorkingTime(2))
	assert.Equal("0.00", model.GetWorkingTime(3))
}
