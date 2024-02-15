package gocask

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadSettings(t *testing.T) {
	var settings Settings

	err := loadSettings(&settings)
	require.Nil(t, err)

	expected := &Settings{
		MaxFileSize: 4096,
	}

	require.Equal(t, *expected, settings)
}
