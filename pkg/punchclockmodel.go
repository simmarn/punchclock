package punchclock

type PunchclockModel struct {
	SelectedMonth     []WorkDayModel
	use24hClockFormat bool
}

func (model *PunchclockModel) GetWorkingTime(index int) string {
	if index < 0 || index >= len(model.SelectedMonth) {
		if model.use24hClockFormat {
			return "0:00"
		}
		return "0.00"
	}
	if model.use24hClockFormat {
		return model.SelectedMonth[index].WorkingTimeClockFormat()
	}
	return model.SelectedMonth[index].WorkingDecimalFormat()
}

func (model *PunchclockModel) ToggleTimeFormat() {
	model.use24hClockFormat = !model.use24hClockFormat
}
