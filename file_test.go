package gocask

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextFileName(t *testing.T) {
	DataDir = "./"
	name, err := nextFileName()

	require.Nil(t, err)
	require.Equal(t, "./gocask_1.data", name)
}

func TestGetFileId(t *testing.T) {
	DataDir = "."
	name, err := nextFileName()

	require.Nil(t, err)

	id, err := getFileId(name)

	require.Nil(t, err)
	require.Equal(t, 1, id)
}

func TestGetFName(t *testing.T) {
	n := 5
	name := getFileNameFromID(n)

	require.Equal(t, DataDir+"gocask_5.data", name)
}
