package punchclock

import (
	"fmt"
	"os"
)

const (
	WORKING string = "Working"
	PAUSED  string = "Paused"
)

type Selected uint8

const (
	CurrentMonth Selected = iota
	PreviousMonth
)

func (s Selected) toString() string {
	switch s {
	case CurrentMonth:
		return "Current"
	case PreviousMonth:
		return "Previous"
	default:
		return "Invalid value"
	}
}

type PunchclockController struct {
	storage            RecordStorage
	punchclock         *PunchClock
	timesheet          *TimesheetModel
	Model              *PunchclockModel
	Status             string
	displayedTimesheet Selected
}

func NewPunchclockController(storage RecordStorage) *PunchclockController {
	c := new(PunchclockController)
	c.storage = storage
	records, err := c.storage.Load()
	CheckIfError(err)
	c.timesheet = NewTimesheetModel(records)
	c.punchclock = NewPunchClockFromData(c.timesheet.GetToday())
	c.Model = new(PunchclockModel)
	c.Refresh()
	c.Status = WORKING
	c.displayedTimesheet = CurrentMonth
	return c
}

func (c *PunchclockController) Present() {
	c.punchclock.Work()
	today := c.punchclock.GetCurrentWorkDay()
	c.timesheet.UpdateWorkDay(CalculateWorkDay(today))
	c.storage.Save(c.timesheet.GetAllRecords())
	c.Status = WORKING
}

func (c *PunchclockController) Pause() {
	c.punchclock.Pause()
	today := c.punchclock.GetCurrentWorkDay()
	c.timesheet.UpdateWorkDay(CalculateWorkDay(today))
	c.Status = PAUSED
}

func (c *PunchclockController) ToggleSelectedMonth() string {
	if c.displayedTimesheet == CurrentMonth {
		c.displayedTimesheet = PreviousMonth
		return CurrentMonth.toString()
	} else {
		c.displayedTimesheet = CurrentMonth
		return PreviousMonth.toString()
	}
}

func (c *PunchclockController) Refresh() {
	if c.Status == WORKING {
		c.Present()
	} else {
		c.Pause()
	}
	if c.displayedTimesheet == CurrentMonth {
		c.Model.SelectedMonth = *NewWorkDayModelSlice(c.timesheet.GetCurrentMonth())
	} else {
		c.Model.SelectedMonth = *NewWorkDayModelSlice(c.timesheet.GetPreviousMonths())
	}
}

func (c *PunchclockController) Update(day WorkDayRecord) {
	c.punchclock.SetCurrentWorkDay(day.WorkDay)
	c.timesheet.UpdateWorkDay(day)
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
