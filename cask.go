package gocask

import (
	"io"
	"os"

	"github.com/philip-edekobi/go-cask/internal/dbmanager"
	"github.com/philip-edekobi/go-cask/internal/serializer"
)

const (
	SettingsFile = "./data/default_settings.json"
)

var DataDir = "./data/datfiles/"

var settings Settings

type CaskOpts struct {
	SyncOnWrite bool
}

type Index struct {
	FileID    int
	Size      int
	Position  int64
	Timestamp int64
}

type BitCaskHandle struct {
	KeyDir map[string]*Index
	DBFile *os.File
}

func (b BitCaskHandle) Get(key string) (string, error) {
	index, ok := b.KeyDir[key]
	if !ok {
		return "", ErrKeyNotFound
	}

	rawData, err := dbmanager.ReadNBytesFromFileAt(b.DBFile, index.Size, index.Position)
	if err != nil {
		return "", err
	}

	dataRecord, err := serializer.DecodeKV(rawData)
	if err != nil {
		return "", err
	}

	return dataRecord.Value, nil
}

func (b BitCaskHandle) Set(key, val string) error {
	if len(key) == 0 {
		return ErrBadKey
	}
	pos, err := b.DBFile.Seek(0, io.SeekEnd)
	if err != nil {
		return err
	}

	rawData := serializer.EncodeKV(key, val)

	err = dbmanager.WriteToFile(b.DBFile, rawData)
	if err != nil {
		return err
	}

	record, err := serializer.DecodeKV(rawData)
	if err != nil {
		return err
	}

	id, err := getFileId(b.DBFile.Name())
	if err != nil {
		return err
	}

	b.KeyDir[key] = buildIndex(record, id, pos)

	return nil
}

func (b BitCaskHandle) ListKeys() []string {
	keys := []string{}

	for k := range b.KeyDir {
		keys = append(keys, k)
	}

	return keys
}

func (b BitCaskHandle) Delete(key string) error {
	err := b.Set(key, "")

	return err
}

func (b BitCaskHandle) Sync() error {
	err := b.DBFile.Sync()

	return err
}

func (b BitCaskHandle) Close() error {
	err := b.DBFile.Close()

	return err
}

func (b BitCaskHandle) Merge() error {
	// create a new map[string]*Record and new ".data" File
	// for each file in dataDir:
	// read them and update their shii in the map, delete keys with empty value fields
	// make a "tmp.data" file and decode and write every encoded kv pair to it
	// rename the file to the #1 file and delete all the files in the dataDir
	// open a new bitcask in that directory and return that???

	return nil
}

func init() {
	loadSettings(&settings)
}

func Open(dir string) (*BitCaskHandle, error) {
	if len(dir) > 0 {
		if dir[len(dir)-1] != '/' {
			dir += "/"
		}
		DataDir = dir
	}
	cask := &BitCaskHandle{}

	// TODO: check if there are other bitcask instances

	newFileName, err := nextFileName()
	if err != nil {
		return nil, err
	}

	dbFile, err := dbmanager.OpenFileRW(newFileName)
	if err != nil {
		return nil, err
	}
	cask.DBFile = dbFile

	// TODO: if there are other bitcasks, copy their KeyDir
	err = buildKeyDir(cask)
	if err != nil {
		return nil, err
	}

	// return bitcask instance
	return cask, nil
}
