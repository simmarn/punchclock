package punchclock_test

import (
	"simmarn/punchclock"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateWorkDay(t *testing.T) {
	assert := assert.New(t)

	workday := OneWorkdayPlease()

	record := punchclock.CalculateWorkDay(workday)

	assert.Equal(7.75, record.WorkingTime.Hours())
	assert.Equal(75.0, record.PauseTime.Minutes())
}

func OneWorkdayPlease() punchclock.WorkDay {
	workday := punchclock.WorkDay{}
	workday.WorkStarted = time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local)
	workday.WorkEnded = time.Date(2023, 2, 15, 17, 0, 0, 0, time.Local)
	lunchPause := punchclock.NewWorkPause()
	lunchPause.Start = time.Date(2023, 2, 15, 12, 0, 0, 0, time.Local)
	lunchPause.End = time.Date(2023, 2, 15, 13, 0, 0, 0, time.Local)
	fikaPause := punchclock.NewWorkPause()
	fikaPause.Start = time.Date(2023, 2, 15, 15, 0, 0, 0, time.Local)
	fikaPause.End = time.Date(2023, 2, 15, 15, 15, 0, 0, time.Local)
	workday.Pauses = append(workday.Pauses, lunchPause)
	workday.Pauses = append(workday.Pauses, fikaPause)
	return workday
}
