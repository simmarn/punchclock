package punchclock

import (
	"errors"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
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
	Status             binding.String
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
	TerminateOnIoReadError(app, err)
	c.timesheet = NewTimesheetModel(records)
	c.punchclock = NewPunchClockFromData(c.timesheet.GetToday())
	c.Model = new(PunchclockModel)
	c.Status = binding.NewString()
	c.Status.Set(WORKING)
	c.Refresh()
	c.displayedTimesheet = CurrentMonth
	c.activateAutoPause()
	return c
}

func (c *PunchclockController) Present() {
	c.punchclock.Work()
	today := c.punchclock.GetCurrentWorkDay()
	c.timesheet.UpdateWorkDay(CalculateWorkDay(today))
	c.storage.Save(c.timesheet.GetAllRecords())
	c.Status.Set(WORKING)
}

func (c *PunchclockController) Pause() {
	c.punchclock.Pause()
	today := c.punchclock.GetCurrentWorkDay()
	c.timesheet.UpdateWorkDay(CalculateWorkDay(today))
	c.Status.Set(PAUSED)
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
	status, _ := c.Status.Get()
	if status == WORKING {
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
	now := GetNow()
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

func (c *PunchclockController) GetAutoPauseStart() string {
	return c.prefs.GetString(PREFAUTOPAUSESTART)
}

func (c *PunchclockController) GetAutoPauseEnd() string {
	return c.prefs.GetString(PREFAUTOPAUSEEND)
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

func (c *PunchclockController) GetAutoPause() bool {
	return c.prefs.GetBool(PREFAUTOPAUSEACTIVE)
}

func (c *PunchclockController) activateAutoPause() {
	if c.prefs.GetBool(PREFAUTOPAUSEACTIVE) {
		now := GetNow()
		nextStartTime, _ := UpdateTime(now, c.prefs.GetString(PREFAUTOPAUSESTART))
		nextEndTime, _ := UpdateTime(now, c.prefs.GetString(PREFAUTOPAUSEEND))
		random := rand.New(rand.NewSource(now.UnixNano()))
		token := random.Int() // token to prevent double autopause when setting is changed
		c.autoPauseToken = token

		if now.After(nextEndTime) {
			nextStartTime = nextStartTime.Add(24 * time.Hour)
			nextEndTime = nextEndTime.Add(24 * time.Hour)
		}

		Log.Info().Msg("Auto pause will start " + nextStartTime.Format(time.RFC3339))
		Log.Info().Msg("Auto pause will end " + nextEndTime.Format(time.RFC3339))

		if now.After(nextStartTime) {
			Log.Info().Msg("Auto pausing...")
			c.Pause()
		} else {

			go func() {
				Sleep(time.Until(nextStartTime.Add(-time.Minute)))
				if c.autoPauseToken == token {
					Log.Info().Msg("Auto pausing...")
					c.Pause()
				}
			}()
		}

		go func() {
			Sleep(time.Until(nextEndTime.Add(5 * time.Second)))
			if c.autoPauseToken == token {
				Log.Info().Msg("Auto pause ended")
				c.Present()
				c.activateAutoPause()
			}
		}()
	} else {
		c.autoPauseToken = 0
		Log.Info().Msg("Auto pause disabled")
		c.Present()
	}
}

func (c *PunchclockController) GetUnworkedDayBefore(t time.Time) *WorkDay {
	return NewWorkDay(t.Add(-24 * time.Hour))
}
