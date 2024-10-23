package test

import (
	"testing"
	"time"

	punchclock "github.com/simmarn/punchclock/pkg"

	"github.com/stretchr/testify/assert"
)

func TestSetStart(t *testing.T) {
	assert := assert.New(t)
	day := OneWorkDayModelPlease()
	assert.Equal("07.75", day.WorkingDecimalFormat())

	assert.Nil(day.SetStart("7:45"))
	assert.Equal("08.00", day.WorkingDecimalFormat())

	assert.NotNil(day.SetStart("seven"))
	assert.Equal("08.00", day.WorkingDecimalFormat())

	assert.NotNil(day.SetStart("23:00"))
	assert.Equal("08.00", day.WorkingDecimalFormat())
}

func TestSetEnd(t *testing.T) {
	assert := assert.New(t)
	day := OneWorkDayModelPlease()
	assert.Equal("07.75", day.WorkingDecimalFormat())

	assert.Nil(day.SetEnd("17:15"))
	assert.Equal("08.00", day.WorkingDecimalFormat())

	assert.NotNil(day.SetEnd("seven"))
	assert.Equal("08.00", day.WorkingDecimalFormat())

	assert.NotNil(day.SetEnd("7:59"))
	assert.Equal("08.00", day.WorkingDecimalFormat())
}

func OneWorkDayModelPlease() punchclock.WorkDayModel {
	workday := OneWorkdayPlease(time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local))
	record := punchclock.CalculateWorkDay(workday)
	model := punchclock.NewWorkDayModel(record)
	return model
}
