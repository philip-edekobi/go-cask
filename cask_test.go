package gocask

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	name, err := nextFileName("./data/datfiles/")
	require.Nil(t, err)

	db, err := Open("./data/datfiles/")
	require.Nil(t, err)

	require.Equal(t, "./data/datfiles/", db.DataDir)
	require.Equal(t, name, db.DBFile.Name())

	db2, err := Open("./testDir/")
	require.Nil(t, err)

	require.Equal(t, "./testDir/", db2.DataDir)

	err = db.DBFile.Close()
	require.Nil(t, err)

	err = db2.DBFile.Close()
	require.Nil(t, err)

	err = os.Remove(name)
	require.Nil(t, err)

	err = os.Remove(db2.DBFile.Name())
	require.Nil(t, err)
}

func TestClose(t *testing.T) {
	name, err := nextFileName("./data/datfiles/")
	require.Nil(t, err)

	db, err := Open("./data/datfiles/")
	require.Nil(t, err)

	err = db.Close()
	require.Nil(t, err)

	err = db.DBFile.Close()
	require.ErrorIs(t, err, os.ErrClosed)

	err = os.Remove(name)
	require.Nil(t, err)
}

func TestGet(t *testing.T) {
	k := "name"
	v := "albert"

	name, err := nextFileName("./data/datfiles/")
	require.Nil(t, err)

	cask, err := Open("./data/datfiles/")
	require.Nil(t, err)

	err = cask.Set(k, v)
	require.Nil(t, err)

	val, err := cask.Get(k)
	require.Nil(t, err)
	require.Equal(t, v, val)

	val, err = cask.Get("alpha")
	require.Equal(t, "", val)
	require.ErrorIs(t, ErrKeyNotFound, err)
	cask.Close()

	name2, err := nextFileName("./data/datfiles/")
	require.Nil(t, err)

	cask, err = Open("./data/datfiles/")
	require.Nil(t, err)
	defer cask.Close()

	err = cask.Set("hey", "hi")
	require.Nil(t, err)

	val, err = cask.Get(k)
	require.Nil(t, err)
	require.Equal(t, v, val)

	err = os.Remove(name)
	require.Nil(t, err)

	err = os.Remove(name2)
	require.Nil(t, err)
}

func TestSet(t *testing.T) {
	DataDir := "./testDir/"

	testCases := []struct {
		keys      []string
		vals      []string
		totalSize int64
	}{
		{
			keys:      []string{"name"},
			vals:      []string{"adam"},
			totalSize: 28,
		},
		{
			keys:      []string{"name", "age"},
			vals:      []string{"adam", "17"},
			totalSize: 53,
		},
		{
			keys: []string{""},
			vals: []string{"adam"},
		},
	}

	for i, tc := range testCases {
		cask, err := Open(DataDir)
		require.Nil(t, err)

		if i == 2 {
			err := cask.Set(tc.keys[0], tc.vals[0])
			require.ErrorIs(t, err, ErrBadKey)

			err = cask.Sync()
			require.Nil(t, err)

			err = cask.Close()
			require.Nil(t, err)

			err = os.Remove(cask.DBFile.Name())
			require.Nil(t, err)

			continue
		}

		for i := 0; i < len(tc.keys); i++ {
			err := cask.Set(tc.keys[i], tc.vals[i])
			require.Nil(t, err)

			err = cask.Sync()
			require.Nil(t, err)
		}

		fileStat, err := cask.DBFile.Stat()
		require.Nil(t, err)
		require.Equal(t, tc.totalSize, fileStat.Size())

		err = cask.Close()
		require.Nil(t, err)

		err = os.Remove(cask.DBFile.Name())
		require.Nil(t, err)
	}
}

func TestDelete(t *testing.T) {
	keys := []string{"name", "age"}
	vals := []string{"adam", "23"}

	name, err := nextFileName("./data/datfiles/")
	require.Nil(t, err)

	cask, err := Open("")
	require.Nil(t, err)
	defer cask.Close()

	for i := 0; i < len(keys); i++ {
		cask.Set(keys[i], vals[i])
	}

	lKeys := cask.ListKeys()
	require.ElementsMatch(t, keys, lKeys)

	err = cask.Delete("age")
	require.Nil(t, err)

	v, err := cask.Get("age")
	require.Nil(t, err)

	require.Equal(t, "", v)

	lKeys = cask.ListKeys()
	require.Equal(t, keys[0], lKeys[0])

	err = os.Remove(name)
	require.Nil(t, err)
}
