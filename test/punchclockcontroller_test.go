package test

import (
	"os/exec"
	"testing"
	"time"

	punchclock "github.com/simmarn/punchclock/pkg"
	"github.com/stretchr/testify/assert"
)

const (
	RemovePreferencesCommand string = "rm"
	PreferencesLocation      string = "~/.config/fyne/com.github.simmarn.punchclock/preferences.json"
)

func TestAutoPause(t *testing.T) {
	assert := assert.New(t)

	RemoveSettings()

	controller := punchclock.NewPunchclockController(new(MockedRecordStorage))
	assert.Equal(punchclock.WORKING, controller.Status)

	start := time.Now().Format(punchclock.HHMMSS24h)
	end := time.Now().Add(time.Minute).Format(punchclock.HHMMSS24h)
	err := controller.SetAutoPauseInterval(start, end)
	assert.Nil(err)
	assert.Equal(punchclock.WORKING, controller.Status)

	controller.SetAutoPause(true)
	assert.Equal(punchclock.PAUSED, controller.Status)

	controller.SetAutoPause(false)
	assert.Equal(punchclock.WORKING, controller.Status)

	controller.SetAutoPause(true)
	assert.Equal(punchclock.PAUSED, controller.Status)

	err = controller.SetAutoPauseInterval("11:00", "11:00")
	assert.NotNil(err)

	err = controller.SetAutoPauseInterval("11:00", "12,00")
	assert.NotNil(err)

	err = controller.SetAutoPauseInterval("eleven", "12:00")
	assert.NotNil(err)

	controller.SetAutoPause(false)
}

func TestAutoPauseNotSet(t *testing.T) {
	assert := assert.New(t)
	RemoveSettings()

	controller := punchclock.NewPunchclockController(new(MockedRecordStorage))
	assert.Equal(punchclock.WORKING, controller.Status)

	err := controller.SetAutoPause(true)
	assert.NotNil(err)
}

func RemoveSettings() {
	cmd := exec.Command("rm", PreferencesLocation)
	cmd.Run()
}

type MockedRecordStorage struct {
}

func (fh *MockedRecordStorage) Save(records []punchclock.WorkDayRecord) error {
	return nil
}

func (fh *MockedRecordStorage) Load() ([]punchclock.WorkDayRecord, error) {
	return CreateRecordsForTest(), nil
}
