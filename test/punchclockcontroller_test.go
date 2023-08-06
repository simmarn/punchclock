package test

import (
	"os/exec"
	"testing"
	"time"

	mocks "github.com/simmarn/punchclock/mocks/github.com/simmarn/punchclock/pkg"
	punchclock "github.com/simmarn/punchclock/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	RemovePreferencesCommand string = "rm"
	PreferencesLocation      string = "~/.config/fyne/com.github.simmarn.punchclock/preferences.json"
)

func TestAutoPause(t *testing.T) {
	assert := assert.New(t)

	RemoveSettings()

	mockRs := MockRecordStorage(t)
	controller := punchclock.NewPunchclockController(mockRs)
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

	mockRs := MockRecordStorage(t)
	controller := punchclock.NewPunchclockController(mockRs)
	assert.Equal(punchclock.WORKING, controller.Status)

	err := controller.SetAutoPause(true)
	assert.NotNil(err)
}

func RemoveSettings() {
	cmd := exec.Command("rm", PreferencesLocation)
	cmd.Run()
}

func MockRecordStorage(t *testing.T) punchclock.RecordStorage {
	mockRs := mocks.NewRecordStorage(t)
	mockRs.EXPECT().Load().Return(CreateRecordsForTest(), nil)
	mockRs.EXPECT().Save(mock.Anything).Return(nil)
	return mockRs
}
