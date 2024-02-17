package gocask

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	name, err := nextFileName()

	require.Nil(t, err)

	db := Open("")

	require.Equal(t, "./data/datfiles/", DataDir)
	require.Equal(t, name, db.DBFile.Name())
}
