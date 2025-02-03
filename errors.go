package harebrain

import (
	"fmt"
)

type Herror struct {
	msg   string
	child error
}

func (h *Herror) Error() string {
	return h.msg
}

func (h *Herror) Unwrap() error {
	return h.child
}

func (h *Herror) Wrap(e error) {
	h.child = e
}

var ErrHareBrain = &Herror{msg: "harebrained"}

var ErrDatabase = fmt.Errorf("%w: database", ErrHareBrain)

//var ErrNoSuchRecord = errors.New("no such record")
//var ErrNoSuchTable = errors.New("no such table")
