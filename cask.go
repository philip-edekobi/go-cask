package gocask

import (
	"os"
	"strconv"

	"github.com/philip-edekobi/go-cask/internal/dbmanager"
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
	ValueSZ   int
	ValuePos  int
	Timestamp int64
}

type BitCaskHandle struct {
	KeyDir map[string]Index
	DBFile *os.File
}

func (b BitCaskHandle) Get(key string) (string, error) {
	val, ok := b.KeyDir[key]
	if !ok {
		return "", CaskError{"key not found in db"}
	}

	// TODO: write a function to extrace the value from the db file
	return strconv.Itoa(val.ValuePos), nil
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
	buildKeyDir(cask)

	// return bitcask instance
	return cask
}
