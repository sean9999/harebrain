package harebrain

import (
	"errors"
	"fmt"
)

type Herror struct {
	msg string
}

func (h *Herror) Error() string {
	return h.msg
}

var ErrHareBrain = &Herror{msg: "harebrained"}

var ErrDatabase = fmt.Errorf("%w: database", ErrHareBrain)
var ErrNoSuchRecord = errors.New("no such record")
var ErrNoSuchTable = errors.New("no such table")
