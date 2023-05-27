package punchclock_test

import (
	punchclock "simmarn/punchclock/pkg"
	"testing"
	"time"

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
	currentMonth := time.Now().Month()
	workday1 := OneWorkdayPlease(time.Date(2023, currentMonth-1, 30, 8, 0, 0, 0, time.Local))
	workday2 := OneWorkdayPlease(time.Date(2023, currentMonth, 2, 8, 0, 0, 0, time.Local))
	workday3 := OneWorkdayPlease(time.Date(2023, currentMonth, 3, 8, 0, 0, 0, time.Local))

	records := []punchclock.WorkDayRecord{
		punchclock.CalculateWorkDay(workday1),
		punchclock.CalculateWorkDay(workday2),
		punchclock.CalculateWorkDay(workday3)}
	return records
}
