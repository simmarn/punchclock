package punchclock

import (
	"fmt"
	"os"
)

type PunchclockController struct {
	storage    RecordStorage
	punchclock *PunchClock
	timesheet  *TimesheetModel
	Model      *PunchclockModel
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
	return c
}

func (c *PunchclockController) Present() {
	c.punchclock.Work()
	today := c.punchclock.GetCurrentWorkDay()
	c.timesheet.UpdateWorkDay(CalculateWorkDay(today))
	c.storage.Save(c.timesheet.GetAllRecords())
}

func (c *PunchclockController) Pause() {
	c.punchclock.Pause()
}

func (c *PunchclockController) GetCurrentMonth() *[]WorkDayModel {
	return NewWorkDayModelSlice(c.timesheet.GetCurrentMonth())
}

func (c *PunchclockController) GetPrevious() []WorkDayRecord {
	return c.timesheet.GetPreviousMonths()
}

func (c *PunchclockController) Refresh() {
	c.Model.CurrentMonth = *NewWorkDayModelSlice(c.timesheet.GetCurrentMonth())
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
