package gocask

import (
	"strconv"
)

type CaskOpts struct {
	Role        string
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

func Open(dir string) *BitCaskHandle {
	return nil
}
