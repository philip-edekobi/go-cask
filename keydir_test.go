package gocask

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

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

// func TestLoad() {}

// func TestBuildKeyDir() {}
