package punchclock

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func Test_Work_WorkEndedUpdated(t *testing.T) {
	assert := assert.New(t)
	initialdata := CreateWorkDay(time.Now().Add(-time.Hour))
	pc := NewPunchClockFromData(*initialdata)
	assert.Equal(pc.GetCurrentWorkDay().WorkEnded, pc.GetCurrentWorkDay().WorkStarted)

	pc.Work()

	assert.Greater(pc.GetCurrentWorkDay().WorkEnded, pc.GetCurrentWorkDay().WorkStarted)
}

func Test_Work_NewDay_WorkEndedUpdatedWithNewDate(t *testing.T) {
	assert := assert.New(t)
	initialdata := CreateWorkDay(time.Now().Add(-time.Hour * 24))
	pc := NewPunchClockFromData(*initialdata)

	pc.Work()

	assert.NotEqual(initialdata.WorkStarted.Day(), pc.GetCurrentWorkDay().WorkStarted.Day())
	assert.NotEqual(initialdata.WorkEnded.Day(), pc.GetCurrentWorkDay().WorkEnded.Day())
	assert.Equal(initialdata.WorkStarted, pc.GetPreviousWorkDay().WorkStarted)
	assert.Equal(initialdata.WorkEnded, pc.GetPreviousWorkDay().WorkEnded)
}

func Test_Pause_PauseSaved(t *testing.T) {
	assert := assert.New(t)
	initialdata := CreateWorkDay(time.Now().Add(-time.Hour))
	pc := NewPunchClockFromData(*initialdata)

	pc.Pause()

	assert.Greater(pc.GetCurrentWorkDay().WorkEnded, pc.GetCurrentWorkDay().WorkStarted)
	assert.Equal(0, len(pc.GetCurrentWorkDay().Pauses))

	time.Sleep(1 * time.Second)
	pc.Work()
	assert.Equal(1, len(pc.GetCurrentWorkDay().Pauses))

	pc.Work()
	assert.Equal(1, len(pc.GetCurrentWorkDay().Pauses))

	pc.Pause()
	pc.Pause()
	pc.Pause()
	time.Sleep(1 * time.Second)
	pc.Work()
	assert.Equal(2, len(pc.GetCurrentWorkDay().Pauses))
}

func CreateWorkDay(initTime time.Time) *WorkDay {
	started := initTime
	today := new(WorkDay)
	today.WorkStarted = started
	today.WorkEnded = started
	pauseSlice := make([]WorkPause, 0)
	today.Pauses = pauseSlice
	return today
}
