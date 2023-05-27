package punchclock_test

import (
	"simmarn/punchclock"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestFileHandler_SaveToFile(t *testing.T) {
	workday1 := OneWorkdayPlease(time.Date(2023, 2, 15, 8, 0, 0, 0, time.Local))
	workday2 := OneWorkdayPlease(time.Date(2023, 2, 16, 8, 0, 0, 0, time.Local))
	workday3 := OneWorkdayPlease(time.Date(2023, 2, 17, 8, 0, 0, 0, time.Local))

	records := []punchclock.WorkDayRecord{
		punchclock.CalculateWorkDay(workday1),
		punchclock.CalculateWorkDay(workday2),
		punchclock.CalculateWorkDay(workday3)}

	fh := punchclock.NewFileHandler("testrecords.json")
	err := fh.SaveToFile(records)
	if err != nil {
		t.Fatalf(`Did not expect %v`, err)
	}

	loadedRecords, err := fh.LoadFromFile()
	if err != nil {
		t.Fatalf(`Did not expect %v`, err)
	}
	if (cmp.Equal(records, loadedRecords)) == false {
		t.Fatal()
	}
}
