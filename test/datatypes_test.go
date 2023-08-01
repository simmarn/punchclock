package test

import (
	punchclock "simmarn/punchclock/pkg"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateWorkDay(t *testing.T) {
	assert := assert.New(t)

	workday := OneWorkdayPlease(time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local))

	record := punchclock.CalculateWorkDay(workday)

	assert.Equal(7.75, record.WorkingTime.Hours())
	assert.Equal(75.0, record.PauseTime.Minutes())
}

func OneWorkdayPlease(daystart time.Time) punchclock.WorkDay {
	workday := punchclock.WorkDay{}
	workday.WorkStarted = daystart
	workday.WorkEnded = daystart.Add(9 * time.Hour)
	lunchPause := punchclock.NewWorkPause()
	lunchPause.Start = daystart.Add(4 * time.Hour)
	lunchPause.End = daystart.Add(5 * time.Hour)
	fikaPause := punchclock.NewWorkPause()
	fikaPause.Start = daystart.Add(7 * time.Hour)
	fikaPause.End = fikaPause.Start.Add(15 * time.Minute)
	workday.Pauses = append(workday.Pauses, lunchPause)
	workday.Pauses = append(workday.Pauses, fikaPause)
	return workday
}
