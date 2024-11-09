package test

import (
	"testing"
	"time"

	punchclock "github.com/simmarn/punchclock/pkg"

	"github.com/google/go-cmp/cmp"
)

func TestFileHandler_SaveToFile(t *testing.T) {

	records := CreateRecordsForTest()

	fh := punchclock.NewFileHandler("testrecords.json")
	err := fh.Save(records)
	if err != nil {
		t.Fatalf(`Did not expect %v`, err)
	}

	loadedRecords, err := fh.Load()
	if err != nil {
		t.Fatalf(`Did not expect %v`, err)
	}
	if (cmp.Equal(records, loadedRecords)) == false {
		t.Fatal()
	}
}

func CreateRecordsForTest() []punchclock.WorkDayRecord {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	workday1 := OneWorkdayPlease(time.Date(currentYear, currentMonth-1, 30, 8, 0, 0, 0, time.Local))
	workday2 := OneWorkdayPlease(time.Date(currentYear, currentMonth, 2, 8, 0, 0, 0, time.Local))
	workday3 := OneWorkdayPlease(time.Date(currentYear, currentMonth, 3, 8, 0, 0, 0, time.Local))

	records := []punchclock.WorkDayRecord{
		punchclock.CalculateWorkDay(workday1),
		punchclock.CalculateWorkDay(workday2),
		punchclock.CalculateWorkDay(workday3)}
	return records
}

func CreateRecordsForTestWithSameMonthLastYear() []punchclock.WorkDayRecord {
	currentYear := time.Now().Year()
	currentMonth := time.Now().Month()
	workday0 := OneWorkdayPlease(time.Date(currentYear-1, currentMonth, 29, 8, 0, 0, 0, time.Local))
	workday1 := OneWorkdayPlease(time.Date(currentYear, currentMonth-1, 30, 8, 0, 0, 0, time.Local))
	workday2 := OneWorkdayPlease(time.Date(currentYear, currentMonth, 2, 8, 0, 0, 0, time.Local))
	workday3 := OneWorkdayPlease(time.Date(currentYear, currentMonth, 3, 8, 0, 0, 0, time.Local))

	records := []punchclock.WorkDayRecord{
		punchclock.CalculateWorkDay(workday0),
		punchclock.CalculateWorkDay(workday1),
		punchclock.CalculateWorkDay(workday2),
		punchclock.CalculateWorkDay(workday3)}
	return records
}
