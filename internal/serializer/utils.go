package serializer

import (
	"encoding/binary"
	"hash/crc32"
	"slices"
)

func int32ToBytes(num int32) []byte {
	bytes := make([]byte, 4)

	binary.LittleEndian.PutUint32(bytes, uint32(num))

	return bytes
}

func int64ToBytes(num int64) []byte {
	bytes := make([]byte, 8)

	binary.LittleEndian.PutUint64(bytes, uint64(num))

	return bytes
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

func verifyHash(record []byte) bool {
	checksum := record[:FIELDSIZE]

	hashBytes := make([]byte, 4)

	crc := crc32.Checksum(record[FIELDSIZE:], crc32.MakeTable(crc32.IEEE))

	binary.LittleEndian.PutUint32(hashBytes, crc)

	return slices.Compare(checksum, hashBytes) == 0
}

func getRecordLength(h *Header) int {
	return HEADERLENGTH + h.KeySize + h.ValueSize
}
