package gocask

import (
	"os"
	"strconv"
	"strings"
)

func nextFileName() (string, error) {
	name := DataDir + "gocask_"
	number := 1

	dir, err := os.ReadDir("./data/datfiles/")
	if err != nil {
		return "", err
	}

	for _, file := range dir {
		if strings.HasSuffix(file.Name(), ".data") {
			number++
		}
	}

	return name + strconv.Itoa(number) + ".data", nil
}
