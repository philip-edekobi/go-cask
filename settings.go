package gocask

import (
	"encoding/json"
	"os"
)

const datafile = "./data/default_settings.json"

type Settings struct {
	MaxFileSize int `json:"max_file_size_bytes"` // in bytes
}

func loadSettings(s *Settings) error {
	jsonFile, err := os.ReadFile(datafile)
	if err != nil {
		return err
	}

	json.Unmarshal(jsonFile, s)

	return nil
}
