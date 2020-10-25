package value

import (
	"errors"
)

var _ AppendValue = (*Read)(nil)

var (
	ErrAppendRead  = errors.New("invalid append data to read value")
	ErrAppendEmpty = errors.New("invalid apped data to empty value")
)

type Read struct {
	Value
}

func (r *Read) Append(string) error {
	return ErrAppendRead
}
