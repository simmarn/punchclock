package punchclock

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"time"

	"fyne.io/fyne/v2"
)

const (
	WORKING             string = "Working"
	PAUSED              string = "Paused"
	PREFAUTOPAUSESTART  string = "PrefAutoPauseStart"
	PREFAUTOPAUSEEND    string = "PrefAutoPauseEnd"
	PREFAUTOPAUSEACTIVE string = "PrefAutoPauseActive"
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
	App                fyne.App
	autoPauseToken     int
	prefs              PreferencesWrapper
}

func NewPunchclockController(storage RecordStorage, prefs PreferencesWrapper, app fyne.App) *PunchclockController {
	c := new(PunchclockController)
	c.App = app
	c.prefs = prefs
	c.storage = storage
	records, err := c.storage.Load()
	CheckIfError(err)
	c.timesheet = NewTimesheetModel(records)
	c.punchclock = NewPunchClockFromData(c.timesheet.GetToday())
	c.Model = new(PunchclockModel)
	c.Refresh()
	c.Status = WORKING
	c.displayedTimesheet = CurrentMonth
	c.activateAutoPause()
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

func (c *PunchclockController) SetAutoPauseInterval(timeStart string, timeEnd string) error {
	now := time.Now()
	nextStartTime, err := UpdateTime(now, timeStart)
	if err != nil {
		return err
	}
	nextEndTime, err := UpdateTime(now, timeEnd)
	if err != nil {
		return err
	}
	if nextStartTime.Before(nextEndTime) {
		c.prefs.SetString(PREFAUTOPAUSESTART, timeStart)
		c.prefs.SetString(PREFAUTOPAUSEEND, timeEnd)
	} else {
		return errors.New("auto pause start time must be before end time")
	}
	return nil
}

func (c *PunchclockController) SetAutoPause(active bool) error {
	start := c.prefs.GetString(PREFAUTOPAUSESTART)
	end := c.prefs.GetString(PREFAUTOPAUSEEND)
	if start == "" || end == "" {
		return errors.New("auto pause interval not set")
	}
	c.prefs.SetBool(PREFAUTOPAUSEACTIVE, active)
	c.activateAutoPause()
	return nil
}

func (c *PunchclockController) activateAutoPause() {
	if c.prefs.GetBool(PREFAUTOPAUSEACTIVE) {
		now := time.Now()
		nextStartTime, _ := UpdateTime(now, c.prefs.GetString(PREFAUTOPAUSESTART))
		nextEndTime, _ := UpdateTime(now, c.prefs.GetString(PREFAUTOPAUSEEND))
		rand.Seed(now.UnixNano())
		token := rand.Int() // token to prevent double autopause when setting is changed
		c.autoPauseToken = token

		if now.After(nextEndTime) {
			nextStartTime.Add(24 * time.Hour)
			nextEndTime.Add(24 * time.Hour)
		}

		if now.After(nextStartTime) {
			c.Pause()
		} else {

			go func() {
				time.Sleep(time.Until(nextStartTime))
				if c.autoPauseToken == token {
					c.Pause()
				}
			}()
		}

		go func() {
			time.Sleep(time.Until(nextEndTime))
			if c.autoPauseToken == token {
				c.Present()
				c.activateAutoPause()
			}
		}()
	} else {
		c.autoPauseToken = 0
		c.Present()
	}
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("error: %s", err))
	os.Exit(1)
}
