package gocask

import (
	"os"
	"strconv"
	"strings"
)

func nextFileName(dataDir string) (string, error) {
	name := dataDir + "gocask_"
	number := 1

	dir, err := os.ReadDir(dataDir)
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

func getFileId(name string) (int, error) {
	// name has format {name}_{num}.data... we're interested in num
	sections := strings.Split(name, "_")
	numStr := strings.Replace(sections[1], ".data", "", -1)

	num, err := strconv.Atoi(numStr)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func getFileNameFromID(dataDir string, id int) string {
	return dataDir + "gocask_" + strconv.Itoa(id) + ".data"
}
