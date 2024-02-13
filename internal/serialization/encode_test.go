package serialization

import (
	"encoding/binary"
	"hash/crc32"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	testData := []byte("Hi brahhh")
	hash := crcHash(testData)

	expectedHashBytes := make([]byte, 4)

	expectedHash := crc32.Checksum(testData, crc32.MakeTable(crc32.IEEE))

	binary.LittleEndian.PutUint32(expectedHashBytes, expectedHash)
	require.Equal(t, hash, expectedHashBytes)
}

func TestConcatSlices(t *testing.T) {
	slices := [][]byte{[]byte("Hi"), []byte("man")}
	expected := []byte("Himan")

	merged := concatSlices(slices)

	require.Equal(t, expected, merged)
}
