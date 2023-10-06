package test

import (
	"os"
	"testing"
	"time"

	"github.com/simmarn/punchclock/logging"
	punchclock "github.com/simmarn/punchclock/pkg"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	// Write code here to run before tests
	SetNoLogging()
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

func Test_Pause_PauseSaved(t *testing.T) {
	assert := assert.New(t)
	startTime := time.Now().Add((-time.Hour))
	initialdata := CreateWorkDay(startTime)
	pc := punchclock.NewPunchClockFromData(*initialdata)
	FakeTimeTo(startTime.Add(10 * time.Minute))
	pc.Pause()

	assert.Greater(pc.GetCurrentWorkDay().WorkEnded, pc.GetCurrentWorkDay().WorkStarted)
	assert.Equal(0, len(pc.GetCurrentWorkDay().Pauses))

	FakeTimeTo(startTime.Add(30 * time.Minute))
	pc.Work()
	assert.Equal(1, len(pc.GetCurrentWorkDay().Pauses))

	pc.Work()
	assert.Equal(1, len(pc.GetCurrentWorkDay().Pauses))

	pc.Pause()
	pc.Pause()
	pc.Pause()
	FakeTimeTo(startTime.Add(50 * time.Minute))
	pc.Work()
	assert.Equal(2, len(pc.GetCurrentWorkDay().Pauses))

	UnFakeTime()
}

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

func FakeTimeTo(t time.Time) {
	punchclock.GetNow = func() time.Time { return t }
	punchclock.Sleep = func(d time.Duration) { time.Sleep(time.Second) }
}

func UnFakeTime() {
	punchclock.GetNow = time.Now
	punchclock.Sleep = time.Sleep
}

func SetNoLogging() {
	punchclock.Log = logging.Configure(logging.Config{})
}
