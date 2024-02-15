package dbmanager

import (
	"os"
)

func OpenFileRW(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}
