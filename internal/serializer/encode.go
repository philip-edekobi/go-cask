package serializer

import (
	"time"
)

func EncodeKV(key, value string) []byte {
	tstamp := time.Now().Unix()
	keySize := len([]byte(key))
	valueSize := len([]byte(value))

	tstampBytes := int64ToBytes(tstamp)
	keySizeBytes := int32ToBytes(int32(keySize))
	valueSizeBytes := int32ToBytes(int32(valueSize))

	// TODO: upgrade to go 1.22 and change this nonsense to slices.Concat
	data := concatSlices(
		[][]byte{tstampBytes, keySizeBytes, valueSizeBytes, []byte(key), []byte(value)},
	)

	checksum := crcHash(data)

	// TODO: upgrade to go 1.22 and change this nonsense to slices.Concat
	return concatSlices([][]byte{checksum, data})
}
