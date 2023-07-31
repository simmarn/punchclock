package punchclock_test

import (
	punchclock "simmarn/punchclock/pkg"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetPauseTime(t *testing.T) {
	assert := assert.New(t)
	workday := OneWorkdayPlease(time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local))

	editor := punchclock.NewBreakTimeEditor(&workday)

	assert.Equal("01:00", editor.GetPauseTime(0))
	assert.Equal("00:15", editor.GetPauseTime(1))
	assert.Equal("0:00", editor.GetPauseTime(2))
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)
	workday := OneWorkdayPlease(time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local))
	editor := punchclock.NewBreakTimeEditor(&workday)

	err := editor.Add("16:00", "16:05")
	assert.Nil(err)
	assert.Equal(3, len(editor.GetPauses()))
	assert.Equal("00:05", editor.GetPauseTime(2))

	err = editor.Add("15:30", "15:30")
	assert.NotNil(err)
	assert.Equal(3, len(editor.GetPauses()))

	err = editor.Add("15:40", "15:30")
	assert.NotNil(err)
	assert.Equal(3, len(editor.GetPauses()))

	err = editor.Add("20:40", "20:50")
	assert.NotNil(err)
	assert.Equal(3, len(editor.GetPauses()))

	err = editor.Add("16:02", "16:07")
	assert.NotNil(err)
	assert.Equal(3, len(editor.GetPauses()))

	err = editor.Add("15:50", "16:07")
	assert.NotNil(err)
	assert.Equal(3, len(editor.GetPauses()))

	err = editor.Add("16:00", "16:05")
	assert.NotNil(err)
	assert.Equal(3, len(editor.GetPauses()))

	err = editor.Add("16:02", "16:04")
	assert.NotNil(err)
	assert.Equal(3, len(editor.GetPauses()))

	err = editor.Add("15:50", "15:59")
	assert.Nil(err)
	assert.Equal(4, len(editor.GetPauses()))

	err = editor.Add("16:06", "16:10")
	assert.Nil(err)
	assert.Equal(5, len(editor.GetPauses()))
}

func TestDelete(t *testing.T) {
	assert := assert.New(t)
	workday := OneWorkdayPlease(time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local))
	editor := punchclock.NewBreakTimeEditor(&workday)

	editor.Delete(0)
	assert.Equal("00:15", editor.GetPauseTime(0))
	assert.Equal("0:00", editor.GetPauseTime(1))
	assert.Equal(1, len(editor.GetPauses()))

	editor.Delete(0)
	assert.Equal("0:00", editor.GetPauseTime(0))
	assert.Equal(0, len(editor.GetPauses()))

	editor.Delete(0)
	assert.Equal(0, len(editor.GetPauses()))
}

func TestEdit(t *testing.T) {
	assert := assert.New(t)
	workday := OneWorkdayPlease(time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local))
	editor := punchclock.NewBreakTimeEditor(&workday)

	err := editor.Edit(1, "14:45", "15:05")
	assert.Nil(err)
	assert.Equal("00:20", editor.GetPauseTime(1))

	err = editor.Edit(0, "7:59", "8:40")
	assert.NotNil(err)
	assert.Equal("01:00", editor.GetPauseTime(0))

	err = editor.Edit(1, "15:05", "15:00")
	assert.NotNil(err)
	assert.Equal("00:20", editor.GetPauseTime(1))

	assert.Equal(2, len(editor.GetPauses()))

	err = editor.Edit(1, "15:00", "15:00")
	assert.Nil(err)
	assert.Equal(1, len(editor.GetPauses()))

	err = editor.Add("14:15", "14:45")
	assert.Nil(err)
	assert.Equal(2, len(editor.GetPauses()))

	err = editor.Edit(0, "14:15", "14:45")
	assert.NotNil(err)
	assert.Equal(2, len(editor.GetPauses()))

	err = editor.Edit(0, "14:20", "14:40")
	assert.NotNil(err)
	assert.Equal(2, len(editor.GetPauses()))

	err = editor.Edit(0, "14:00", "14:30")
	assert.NotNil(err)
	assert.Equal(2, len(editor.GetPauses()))

	err = editor.Edit(0, "14:30", "15:00")
	assert.NotNil(err)
	assert.Equal(2, len(editor.GetPauses()))
}
