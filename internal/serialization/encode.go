package serialization

import (
	"encoding/binary"
	"hash/crc32"
	"time"
)

func encodeKV(key, value string) []byte {
	tstamp := time.Now().Unix()
	keySize := len([]byte(key))
	valueSize := len([]byte(value))

	tstampBytes := make([]byte, 8)
	keySizeBytes := make([]byte, 4)
	valueSizeBytes := make([]byte, 4)

	binary.LittleEndian.PutUint32(tstampBytes, uint32(keySize))
	binary.LittleEndian.PutUint32(tstampBytes, uint32(valueSize))
	binary.LittleEndian.PutUint64(tstampBytes, uint64(tstamp))

	// TODO: upgrade to go 1.22 and change this nonsense to slices.Concat
	data := concatSlices(
		[][]byte{tstampBytes, keySizeBytes, valueSizeBytes, []byte(key), []byte(value)},
	)

	checksum := crcHash(data)

	return concatSlices([][]byte{checksum, data})
}

func concatSlices(slices [][]byte) []byte {
	var combinedLen int

	for _, s := range slices {
		combinedLen += len(s)
	}

	result := make([]byte, combinedLen)

	var i int
	for _, s := range slices {
		i += copy(result[i:], s)
	}

	return result
}

func crcHash(inputData []byte) []byte {
	hashBytes := make([]byte, 4)

	hash := crc32.Checksum(inputData, crc32.MakeTable(crc32.IEEE))

	binary.LittleEndian.PutUint32(hashBytes, hash)

	return hashBytes
}
