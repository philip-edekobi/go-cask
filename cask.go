package gocask

import (
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
		return "", CaskError{"key not found in db"}
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
	rawData := serializer.EncodeKV(key, val)

	err := dbmanager.WriteToFile(b.DBFile, rawData)
	if err != nil {
		return err
	}

	return nil
}

func (b BitCaskHandle) ListKeys() []string {
	keys := []string{}

	for k := range b.KeyDir {
		keys = append(keys, k)
	}

	return keys
}

func init() {
	loadSettings(&settings)
}

func Open(dir string) *BitCaskHandle {
	if len(dir) > 0 {
		DataDir = dir
	}
	cask := &BitCaskHandle{}
	// check if there are other bitcask instances

	// get file next name
	newFileName, err := nextFileName()
	if err != nil {
		panic(err)
	}

	// create and open file
	dbFile, err := dbmanager.OpenFileRW(newFileName)
	if err != nil {
		panic(err)
	}
	cask.DBFile = dbFile

	// read previous files and build KeyDir
	// TODO: if there are other bitcasks, copy their KeyDir
	err = buildKeyDir(cask)
	if err != nil {
		panic(err)
	}

	// return bitcask instance
	return cask
}
