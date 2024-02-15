package gocask

import (
	"os"
)

func buildKeyDir(cask *BitCaskHandle) {
	keydir := cask.KeyDir

	// for each file in datadir
	// load(keydir, file)
	// file.close()
}

func buildIndex(timestamp int64)

func load(keydir map[string]Index, file *os.File) {}
