package punchclock

import (
	"errors"
)

type BreakTimeEditor struct {
	workday *WorkDay
}

func NewBreakTimeEditor(w *WorkDay) *BreakTimeEditor {
	e := new(BreakTimeEditor)
	e.workday = w
	return e
}

func (e *BreakTimeEditor) GetPauses() []WorkPause {
	return e.workday.Pauses
}

func (e *BreakTimeEditor) GetStart(index int) string {
	return e.workday.Pauses[index].Start.Format(HHMMSS24h)
}

func (e *BreakTimeEditor) GetEnd(index int) string {
	return e.workday.Pauses[index].End.Format(HHMMSS24h)
}

func (e *BreakTimeEditor) GetPauseTime(index int) string {
	if len(e.workday.Pauses) > index {
		pause := e.workday.Pauses[index]
		diff := pause.End.Sub(pause.Start)
		return fmtDuration(diff)
	}
	return "0:00"
}

func (e *BreakTimeEditor) Add(startTime string, endTime string) error {
	pause, err := e.convertFromString(startTime, endTime)
	if err != nil {
		return err
	}
	if pause.Start.Equal(pause.End) {
		return errors.New("Pause is has zero length")
	}
	intersects, _ := e.intersectsWithExisting(pause)
	if intersects {
		return errors.New("Pause intersects with other pause")
	}
	e.workday.Pauses = append(e.workday.Pauses, pause)
	return nil
}

func (e *BreakTimeEditor) Delete(index int) {
	if len(e.workday.Pauses) > index {
		newSlice := append(e.workday.Pauses[:index], e.workday.Pauses[index+1:]...)
		e.workday.Pauses = newSlice
	}
}

func (e *BreakTimeEditor) Edit(index int, newStartTime string, newEndTime string) error {
	if len(e.workday.Pauses) > index {
		pause, err := e.convertFromString(newStartTime, newEndTime)
		if err != nil {
			return err
		}
		if pause.Start.Equal(pause.End) {
			e.Delete(index)
			return nil
		}
		intersects, i := e.intersectsWithExisting(pause)
		if intersects && i != index {
			return errors.New("pause intersects with other pause")
		}
		e.workday.Pauses[index] = pause
		return nil
	}
	return errors.New("index out of bounds")
}

func (e *BreakTimeEditor) convertFromString(startTime string, endTime string) (WorkPause, error) {
	start, err := UpdateTime(e.workday.WorkStarted, startTime)
	if err != nil {
		return NewWorkPause(), err
	}
	end, err := UpdateTime(e.workday.WorkStarted, endTime)
	if err != nil {
		return NewWorkPause(), err
	}
	if start.After(end) {
		return NewWorkPause(), errors.New("start time is before end time")
	}
	if start.After(e.workday.WorkStarted) && end.Before(e.workday.WorkEnded) {
		return NewWorkPauseWithData(start, end), nil
	} else {
		return NewWorkPause(), errors.New("pause is not within boundaries for current work day")
	}
}

func (e *BreakTimeEditor) intersectsWithExisting(pause WorkPause) (bool, int) {
	for index, p := range e.workday.Pauses {
		if intersects(pause, p) {
			return true, index
		}
	}
	return false, -1
}

func intersects(p1 WorkPause, p2 WorkPause) bool {
	if p1.End.Before(p2.Start) {
		return false
	}
	if p1.Start.After(p2.End) {
		return false
	}
	return true
}
