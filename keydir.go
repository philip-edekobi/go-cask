package gocask

import (
	"io"
	"os"
	"strings"

	"github.com/philip-edekobi/go-cask/internal/serializer"
)

func buildKeyDir(cask *BitCaskHandle) error {
	cask.KeyDir = make(map[string]*Index)
	// for each file in datadir
	entries, err := os.ReadDir(cask.DataDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".data") {
			continue
		}

		file, err := os.Open(cask.DataDir + entry.Name())
		if err != nil {
			return err
		}
		defer file.Close()

		err = load(cask.KeyDir, file)
		if err != nil {
			return err
		}
	}

	return nil
}

func buildIndex(record *serializer.Record, fileID int, position int64) *Index {
	idx := &Index{}

	idx.FileID = fileID
	idx.Size = serializer.HEADERLENGTH + record.KeySz + record.ValueSz
	idx.Timestamp = record.Timestamp
	idx.Position = position

	return idx
}

func load(keydir map[string]*Index, file *os.File) error {
	var seekIndex int64
	var maxSize int64
	currentHeader := make([]byte, 20)

	fileStat, err := file.Stat()
	if err != nil {
		return err
	}
	maxSize = fileStat.Size()

	seekIndex, err = file.Seek(int64(seekIndex), io.SeekStart)

	for err == nil {
		if seekIndex >= maxSize {
			break
		}
		_, err = file.Read(currentHeader)
		if err != nil {
			return err
		}

		header, err := serializer.DecodeHeader(currentHeader)
		if err != nil {
			return err
		}

		recordBuf := make([]byte, header.KeySize+header.ValueSize+serializer.HEADERLENGTH)

		copy(recordBuf[:serializer.HEADERLENGTH], currentHeader)
		_, err = file.Read(recordBuf[serializer.HEADERLENGTH:])
		if err != nil {
			return err
		}

		record, err := serializer.DecodeKV(recordBuf)
		if err != nil {
			return err
		}

		id, err := getFileId(file.Name())
		if err != nil {
			return err
		}

		keydir[record.Key] = buildIndex(record, id, seekIndex)

		// seekIndex, err = file.Seek(nextIdx, io.SeekStart) // here nextIdx = seekSTart + kv sized
		seekIndex, err = file.Seek(0, io.SeekCurrent)
		if err != nil {
			return err
		}
	}

	if err != io.EOF {
		return err
	}

	return nil
}
