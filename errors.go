package gocask

import (
	"errors"
	"fmt"
)

type CaskError struct {
	message string
}

func (e CaskError) Error() string {
	return fmt.Sprintf(e.message)
}

var (
	ErrBadKey      = errors.New("invalid key")
	ErrKeyNotFound = errors.New("the required key does not exist")
)
