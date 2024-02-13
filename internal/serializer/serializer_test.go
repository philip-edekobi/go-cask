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
