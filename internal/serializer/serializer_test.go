package serializer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncoderAndDecoder(t *testing.T) {
	key := "name"
	val := "adam"

	codedBytes := EncodeKV(key, val)

	decodedRecord, err := DecodeKV(codedBytes)

	require.Nil(t, err)
	require.Equal(t, key, decodedRecord.Key)
	require.Equal(t, val, decodedRecord.Value)
}

func TestDecodeHeader(t *testing.T) {
	key := "k1"
	val := "v1"

	encoded := EncodeKV(key, val)
	decoded, decodeErr := DecodeKV(encoded)

	require.Nil(t, decodeErr)

	header, err := DecodeHeader(encoded[:20])

	require.Nil(t, err)
	require.Equal(t, header.CrcChecksum, decoded.CrcChecksum)
	require.Equal(t, header.KeySize, len([]byte(key)))
	require.Equal(t, header.ValueSize, len([]byte(val)))
}
