package gocask

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNextFileName(t *testing.T) {
	DataDir = "./"
	name, err := nextFileName()

	require.Nil(t, err)
	require.Equal(t, "././gocask_1.data", name)
}

func TestGetFileId(t *testing.T) {
	DataDir = "./"
	name, err := nextFileName()

	require.Nil(t, err)

	file, err := os.Create(name)
	file.Close()

	require.Nil(t, err)

	id, err := getFileId(file.Name())

	require.Nil(t, err)
	require.Equal(t, 1, id)

	err = os.Remove(name)

	require.Nil(t, err)
}
