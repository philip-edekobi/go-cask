package serializer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncoderAndDecoder(t *testing.T) {
	key := "name"
	val := "adam"

	codedBytes := EncodeKV(key, val)

	decodedRecord := DecodeKV(codedBytes)

	require.Equal(t, key, decodedRecord.Key)
	require.Equal(t, val, decodedRecord.Value)
}
