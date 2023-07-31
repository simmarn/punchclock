package punchclock

import (
	"encoding/json"
	"errors"
	"os"
)

type RecordStorage interface {
	Save(records []WorkDayRecord) error
	Load() ([]WorkDayRecord, error)
}

type FileHandler struct {
	filepath string
}

func NewFileHandler(filepath string) *FileHandler {
	fh := new(FileHandler)
	fh.filepath = filepath
	return fh
}

func (fh *FileHandler) Save(records []WorkDayRecord) error {
	recordsJson, err := json.MarshalIndent(records, "", "  ")
	if err != nil {
		return err
	}
	err = os.WriteFile(fh.filepath, recordsJson, 0664)
	return err
}

func (fh *FileHandler) Load() ([]WorkDayRecord, error) {
	if _, err := os.Stat(fh.filepath); err == nil {
		recordsJson, err := os.ReadFile(fh.filepath)
		if err != nil {
			return nil, err
		}
		records := []WorkDayRecord{}
		err = json.Unmarshal(recordsJson, &records)
		if err != nil {
			return nil, err
		}
		return records, nil
	} else if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	} else {
		return nil, err
	}
}
