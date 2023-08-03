package test

import (
	"testing"
	"time"

	punchclock "github.com/simmarn/punchclock/pkg"
	"github.com/stretchr/testify/assert"
)

func TestAutoPause(t *testing.T) {
	assert := assert.New(t)

	controller := punchclock.NewPunchclockController(new(MockedRecordStorage))
	assert.Equal(punchclock.WORKING, controller.Status)

	start := time.Now().Format(punchclock.HHMMSS24h)
	end := time.Now().Add(time.Minute).Format(punchclock.HHMMSS24h)
	err := controller.SetAutoPauseInterval(start, end)
	assert.Nil(err)
	assert.Equal(punchclock.WORKING, controller.Status)

	controller.SetAutoPause(true)
	assert.Equal(punchclock.PAUSED, controller.Status)
}

type MockedRecordStorage struct {
}

func (fh *MockedRecordStorage) Save(records []punchclock.WorkDayRecord) error {
	return nil
}

func (fh *MockedRecordStorage) Load() ([]punchclock.WorkDayRecord, error) {
	return CreateRecordsForTest(), nil
}
