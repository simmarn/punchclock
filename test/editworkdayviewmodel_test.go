package punchclock_test

import (
	punchclock "simmarn/punchclock/pkg"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

/*func TestTimeEntry(t *testing.T) {
	assert := assert.New(t)

	timeEntry := punchclock.NewTimeEntry(time.UnixMilli(0))
	baseTime := time.Date(2023, 2, 14, 8, 35, 0, 0, time.Local)
	expected := time.Date(2023, 2, 14, 7, 42, 0, 0, time.Local)

	timeEntry.SetTime(baseTime)

	test.Type(timeEntry.Entry, "7:42\n")
	test.DoubleTap(timeEntry.Entry)
	assert.Equal(expected, timeEntry.GetTime())
}*/

func TestUpdateTime_ValidTime(t *testing.T) {
	assert := assert.New(t)
	baseTime := time.Date(2023, 2, 14, 8, 35, 0, 0, time.Local)
	expected := time.Date(2023, 2, 14, 7, 42, 0, 0, time.Local)

	newTime, err := punchclock.UpdateTime(baseTime, "7:42")
	assert.Nil(err)
	assert.Equal(expected, newTime)

	expected = time.Date(2023, 2, 14, 0, 0, 0, 0, time.Local)
	newTime, err = punchclock.UpdateTime(baseTime, "0:00")
	assert.Nil(err)
	assert.Equal(expected, newTime)

	expected = time.Date(2023, 2, 14, 23, 1, 0, 0, time.Local)
	newTime, err = punchclock.UpdateTime(baseTime, "23:1")
	assert.Nil(err)
	assert.Equal(expected, newTime)
}

func TestUpdateTime_InvalidInput(t *testing.T) {
	assert := assert.New(t)
	baseTime := time.Date(2023, 2, 14, 8, 35, 0, 0, time.Local)

	newTime, err := punchclock.UpdateTime(baseTime, "7,42")
	assert.NotNil(err)
	assert.Equal(baseTime, newTime)

	newTime, err = punchclock.UpdateTime(baseTime, "24:00")
	assert.NotNil(err)
	assert.Equal(baseTime, newTime)

	newTime, err = punchclock.UpdateTime(baseTime, "22:60")
	assert.NotNil(err)
	assert.Equal(baseTime, newTime)
}
