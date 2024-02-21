package gocask

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/philip-edekobi/go-cask/internal/dbmanager"
	"github.com/philip-edekobi/go-cask/internal/serializer"
)

var TZ = time.Date(2020, 10, 20, 10, 00, 00, 00, time.UTC).Unix()

func TestBuildIndex(t *testing.T) {
	key := "name"
	val := "adam"

	expectedIdx := &Index{
		Position:  0,
		Size:      28,
		FileID:    1,
		Timestamp: TZ,
	}

	enc := serializer.EncodeKV(key, val)
	r, err := serializer.DecodeKV(enc)

	require.Nil(t, err)

	idx := buildIndex(r, 1, 0)
	require.Equal(t, expectedIdx.Position, idx.Position)
	require.Equal(t, expectedIdx.Size, idx.Size)
	require.Equal(t, expectedIdx.FileID, idx.FileID)
}

func TestLoad(t *testing.T) {
	// make key and value pairs
	keys := []string{"name", "nom"}
	vals := []string{"Adam", "Mohammed"}

	encodings := [][]byte{}
	records := []*serializer.Record{}

	for i := 0; i < 2; i++ {
		encodings = append(encodings, serializer.EncodeKV(keys[i], vals[i]))
	}

	for _, enc := range encodings {
		rec, err := serializer.DecodeKV(enc)

		require.Nil(t, err)
		records = append(records, rec)
	}

	// make 2 files... write the first record to the first one, write all to the other
	file1, err := os.Create("d_1.data")
	require.Nil(t, err)

	file2, err := os.Create("d_2.data")
	require.Nil(t, err)

	dbmanager.WriteToFile(file1, encodings[0])

	dbmanager.WriteToFile(file2, encodings[0])
	dbmanager.WriteToFile(file2, encodings[1])

	err = file1.Sync()
	require.Nil(t, err)

	err = file2.Sync()
	require.Nil(t, err)

	err = file1.Close()
	require.Nil(t, err)

	err = file2.Close()
	require.Nil(t, err)

	// make expected keydirs crafted carefully

	expKeyDir1 := map[string]*Index{
		"name": buildIndex(records[0], 1, 0),
	}

	expKeyDir2 := map[string]*Index{
		"name": buildIndex(records[0], 2, 0),
		"nom":  buildIndex(records[1], 2, 28),
	}

	// run the function
	kd1 := make(map[string]*Index)
	kd2 := make(map[string]*Index)

	file1, err = os.Open("d_1.data")
	require.Nil(t, err)
	defer file1.Close()

	file2, err = os.Open("d_2.data")
	require.Nil(t, err)
	defer file1.Close()

	err = load(kd1, file1)
	require.Nil(t, err)

	err = load(kd2, file2)
	require.Nil(t, err)

	// compare the results
	require.Equal(t, expKeyDir1, kd1)
	require.Equal(t, expKeyDir2, kd2)

	// clean up
	err = os.Remove("d_1.data")
	require.Nil(t, err)

	err = os.Remove("d_2.data")
	require.Nil(t, err)
}

func TestBuildKeyDir(t *testing.T) {
	cask := &BitCaskHandle{}
	cask.DataDir = "./"

	keys := []string{"name", "nom"}
	vals := []string{"Adam", "Mohammed"}

	encodings := [][]byte{}
	records := []*serializer.Record{}
	indexes := []*Index{}

	for i := 0; i < 2; i++ {
		encodings = append(encodings, serializer.EncodeKV(keys[i], vals[i]))
	}

	for _, enc := range encodings {
		rec, err := serializer.DecodeKV(enc)

		require.Nil(t, err)
		records = append(records, rec)
	}

	file, err := os.Create("d_1.data")
	require.Nil(t, err)

	dbmanager.WriteToFile(file, encodings[0])
	dbmanager.WriteToFile(file, encodings[1])

	err = file.Sync()
	require.Nil(t, err)

	err = file.Close()
	require.Nil(t, err)

	file, err = os.Open("d_1.data")
	require.Nil(t, err)

	cask.DBFile = file

	err = buildKeyDir(cask)
	require.Nil(t, err)

	for i, r := range records {
		var pos int64
		if i == 1 {
			pos = 28
		}
		indexes = append(indexes, buildIndex(r, 1, pos))
	}

	for i, n := range keys {
		require.Equal(t, indexes[i], cask.KeyDir[n])
	}

	// clean up
	err = os.Remove("d_1.data")
	require.Nil(t, err)
}
