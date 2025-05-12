package test

import (
	"testing"

	mocks "github.com/simmarn/punchclock/mocks/github.com/simmarn/punchclock/pkg"
	punchclock "github.com/simmarn/punchclock/pkg"
	"github.com/stretchr/testify/mock"
)

/*
	func TestAutoPause(t *testing.T) {
		assert := assert.New(t)

		mockRs := MockRecordStorage(t)
		mockPrefs := mocks.NewPreferencesWrapper(t)
		mockPrefs.EXPECT().SetString(mock.Anything, mock.Anything)
		mockPrefs.EXPECT().SetBool(mock.Anything, mock.Anything)
		mockPrefs.EXPECT().GetBool(punchclock.PREFAUTOPAUSEACTIVE).Return(false).Once()
		controller := punchclock.NewPunchclockController(mockRs, mockPrefs, nil)
		mockPrefs.AssertCalled(t, "GetBool", punchclock.PREFAUTOPAUSEACTIVE)
		assert.Equal(punchclock.WORKING, GetStatus(controller))

		start := time.Now().Format(punchclock.HHMMSS24h)
		end := time.Now().Add(time.Minute).Format(punchclock.HHMMSS24h)
		mockPrefs.EXPECT().GetString(punchclock.PREFAUTOPAUSESTART).Return(start)
		mockPrefs.EXPECT().GetString(punchclock.PREFAUTOPAUSEEND).Return(end)
		err := controller.SetAutoPauseInterval(start, end)
		mockPrefs.AssertCalled(t, "SetString", punchclock.PREFAUTOPAUSESTART, start)
		mockPrefs.AssertCalled(t, "SetString", punchclock.PREFAUTOPAUSEEND, end)
		mockPrefs.AssertNumberOfCalls(t, "SetString", 2)
		assert.Nil(err)
		assert.Equal(punchclock.WORKING, GetStatus(controller))

		mockPrefs.EXPECT().GetBool(punchclock.PREFAUTOPAUSEACTIVE).Return(true).Once()
		err = controller.SetAutoPause(true)
		assert.Nil(err)
		assert.Equal(punchclock.PAUSED, GetStatus(controller))

		mockPrefs.EXPECT().GetBool(punchclock.PREFAUTOPAUSEACTIVE).Return(false).Once()
		controller.SetAutoPause(false)
		assert.Equal(punchclock.WORKING, GetStatus(controller))

		mockPrefs.EXPECT().GetBool(punchclock.PREFAUTOPAUSEACTIVE).Return(true).Once()
		controller.SetAutoPause(true)
		assert.Equal(punchclock.PAUSED, GetStatus(controller))

		err = controller.SetAutoPauseInterval("11:00", "11:00")
		assert.NotNil(err)
		mockPrefs.AssertNumberOfCalls(t, "SetString", 2)

		err = controller.SetAutoPauseInterval("11:00", "12,00")
		assert.NotNil(err)
		mockPrefs.AssertNumberOfCalls(t, "SetString", 2)

		err = controller.SetAutoPauseInterval("eleven", "12:00")
		assert.NotNil(err)
		mockPrefs.AssertNumberOfCalls(t, "SetString", 2)
	}

	func TestAutoPauseNotSet(t *testing.T) {
		assert := assert.New(t)

		mockPrefs := mocks.NewPreferencesWrapper(t)
		mockPrefs.EXPECT().GetBool(punchclock.PREFAUTOPAUSEACTIVE).Return(false)
		mockPrefs.EXPECT().GetString(punchclock.PREFAUTOPAUSESTART).Return("")
		mockPrefs.EXPECT().GetString(punchclock.PREFAUTOPAUSEEND).Return("")

		mockRs := MockRecordStorage(t)
		controller := punchclock.NewPunchclockController(mockRs, mockPrefs, nil)
		assert.Equal(punchclock.WORKING, GetStatus(controller))

		err := controller.SetAutoPause(true)
		assert.NotNil(err)
	}
*/
func MockRecordStorage(t *testing.T) punchclock.RecordStorage {
	mockRs := mocks.NewRecordStorage(t)
	mockRs.EXPECT().Load().Return(CreateRecordsForTest(), nil)
	mockRs.EXPECT().Save(mock.Anything).Return(nil)
	return mockRs
}

func GetStatus(c *punchclock.PunchclockController) string {
	status, _ := c.Status.Get()
	return status
}
