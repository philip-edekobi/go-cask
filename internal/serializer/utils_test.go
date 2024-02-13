package serializer

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConcatSlices(t *testing.T) {
	slices := [][]byte{[]byte("Hi"), []byte("man")}
	expected := []byte("Himan")

	merged := concatSlices(slices)

	require.Equal(t, expected, merged)
}

func TestVerifyHash(t *testing.T) {
	record := EncodeKV("name", "adam")

	require.True(t, verifyHash(record))
}
