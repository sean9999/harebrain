package harebrain

import (
	"errors"
	"fmt"
)

type HarebrainError struct {
	msg string
}

func (h *HarebrainError) Error() string {
	return h.msg
}

var ErrHareBrain = &HarebrainError{msg: "harebrain"}

var ErrDatabase = fmt.Errorf("%w: database", ErrHareBrain)
var ErrNoSuchRecord = errors.New("no such record")
var ErrNoSuchTable = errors.New("no such table")
