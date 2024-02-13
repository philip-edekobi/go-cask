package serializer

import (
	"encoding/binary"
	"hash/crc32"
)

func crcHash(inputData []byte) []byte {
	hashBytes := make([]byte, 4)

	hash := crc32.Checksum(inputData, crc32.MakeTable(crc32.IEEE))

	binary.LittleEndian.PutUint32(hashBytes, hash)

	return hashBytes
}
