package gocask

import "fmt"

type CaskError struct {
	message string
}

func (e CaskError) Error() string {
	return fmt.Sprintf(e.message)
}
