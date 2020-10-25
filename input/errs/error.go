package errs

import (
	"errors"
	"fmt"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrNoArgs         = errors.New("no arguments expected")
	ErrToManyArgs     = errors.New("too many arguments")
	ErrUnexpectedType = errors.New("unexpected type")
	ErrRequired       = errors.New("is required")
	ErrAppend         = errors.New("failed append")
	ErrInvalidName    = errors.New("invalid name")
)

func New(name, t string, err error) Error {
	return Error{
		name: name,
		t:    t,
		err:  err,
	}
}

type Error struct {
	name string
	err  error
	t    string
}

func (o Error) Error() string {
	return fmt.Sprintf("%s: '%s' %s", o.t, o.name, o.err)
}

func (o Error) Is(err error) bool {
	return errors.Is(err, o.err)
}

func (o Error) Unwrap() error {
	return o.err
}

func Option(name string, err error) Error {
	return Error{
		name: name,
		err:  err,
		t:    "option",
	}
}

func Argument(name string, err error) Error {
	return Error{
		name: name,
		err:  err,
		t:    "argument",
	}
}
