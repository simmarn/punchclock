package punchclock_test

import (
	"os"
	punchclock "simmarn/punchclock/pkg"
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
	pc := punchclock.NewPunchClockFromData(*initialdata)
	assert.Equal(pc.GetCurrentWorkDay().WorkEnded, pc.GetCurrentWorkDay().WorkStarted)

	pc.Work()

	assert.Greater(pc.GetCurrentWorkDay().WorkEnded, pc.GetCurrentWorkDay().WorkStarted)
}

func Test_Work_NewDay_WorkEndedUpdatedWithNewDate(t *testing.T) {
	assert := assert.New(t)
	initialdata := CreateWorkDay(time.Now().Add(-time.Hour * 24))
	pc := punchclock.NewPunchClockFromData(*initialdata)

	pc.Work()

	assert.NotEqual(initialdata.WorkStarted.Day(), pc.GetCurrentWorkDay().WorkStarted.Day())
	assert.NotEqual(initialdata.WorkEnded.Day(), pc.GetCurrentWorkDay().WorkEnded.Day())
	assert.Equal(initialdata.WorkStarted, pc.GetPreviousWorkDay().WorkStarted)
	assert.Equal(initialdata.WorkEnded, pc.GetPreviousWorkDay().WorkEnded)
}

/*func Test_Pause_PauseSaved(t *testing.T) {
	assert := assert.New(t)
	initialdata := CreateWorkDay(time.Now().Add(-time.Hour))
	pc := punchclock.NewPunchClockFromData(*initialdata)

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
}*/

func Test_SetCurrentWorkDay_NoCurrent_Ignored(t *testing.T) {
	assert := assert.New(t)

	initialdata := CreateWorkDay(time.Now())
	pc := punchclock.NewPunchClockFromData(*initialdata)

	yesterday := CreateWorkDay(time.Now().Add(-24 * time.Hour))
	pc.SetCurrentWorkDay(*yesterday)

	assert.Equal(initialdata.WorkStarted, pc.GetCurrentWorkDay().WorkStarted)
}

func Test_SetCurrentWorkDay_TodayModified_Changed(t *testing.T) {
	assert := assert.New(t)

	initialdata := CreateWorkDay(time.Now())
	pc := punchclock.NewPunchClockFromData(*initialdata)

	todayModified := CreateWorkDay(time.Now().Add(-1 * time.Hour))
	pc.SetCurrentWorkDay(*todayModified)

	assert.Equal(todayModified.WorkStarted, pc.GetCurrentWorkDay().WorkStarted)
}

func CreateWorkDay(initTime time.Time) *punchclock.WorkDay {
	started := initTime
	today := new(punchclock.WorkDay)
	today.WorkStarted = started
	today.WorkEnded = started
	pauseSlice := make([]punchclock.WorkPause, 0)
	today.Pauses = pauseSlice
	return today
}
