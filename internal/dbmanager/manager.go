package dbmanager

import (
	"io"
	"os"
)

func OpenFileRW(filename string) (*os.File, error) {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ReadNBytesFromFileAt(file *os.File, num int, pos int64) ([]byte, error) {
	buf := make([]byte, num)

	_, err := file.Seek(pos, io.SeekStart)
	if err != nil {
		return buf, err
	}

	_, err = file.Read(buf)
	if err != nil {
		return buf, err
	}

	return buf, nil
}

func WriteToFile(file *os.File, data []byte) error {
	_, err := file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
