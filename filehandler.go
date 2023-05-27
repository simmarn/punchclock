package punchclock

import (
	"encoding/json"
	"os"
)

type FileHandler struct {
	filepath string
}

func NewFileHandler(filepath string) *FileHandler {
	fh := new(FileHandler)
	fh.filepath = filepath
	return fh
}

func (fh *FileHandler) SaveToFile(records []WorkDayRecord) error {
	recordsJson, err := json.Marshal(records)
	if err != nil {
		return err
	}
	err = os.WriteFile(fh.filepath, recordsJson, 0664)
	return err
}

func (fh *FileHandler) LoadFromFile() ([]WorkDayRecord, error) {
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
}
