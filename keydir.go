package gocask

import (
	"io"
	"os"

	"github.com/philip-edekobi/go-cask/internal/serializer"
)

func buildKeyDir(cask *BitCaskHandle) error {
	keydir := cask.KeyDir

	// for each file in datadir
	entries, err := os.ReadDir(DataDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		file, err := os.Open(entry.Name())
		if err != nil {
			return err
		}
		defer file.Close()

		err = load(keydir, file)
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

	for err != nil {
		_, err = file.Read(currentHeader)

		header, err := serializer.DecodeHeader(currentHeader)
		if err != nil {
			return err
		}

		recordBuf := make([]byte, header.KeySize+header.ValueSize+serializer.HEADERLENGTH)
		record, err := serializer.DecodeKV(recordBuf)

		id, err := getFileId(file.Name())
		if err != nil {
			return err
		}

		keydir[record.Key] = buildIndex(record, id, seekIndex)

		nextIdx := int64(record.KeySz) + int64(record.ValueSz) + serializer.HEADERLENGTH
		if nextIdx > maxSize {
			break
		}

		// seekIndex, err = file.Seek(nextIdx, io.SeekStart) // here nextIdx = seekSTart + kv sized
		seekIndex, err = file.Seek(nextIdx, io.SeekCurrent)
	}

	if err != io.EOF {
		return err
	}

	return nil
}
