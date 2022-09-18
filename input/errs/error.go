package errs

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrNoArgs         = errors.New("no arguments expected")
	ErrToManyArgs     = errors.New("too many arguments")
	ErrUnexpectedType = errors.New("unexpected type")
	ErrRequired       = errors.New("is required")
	ErrAppend         = errors.New("failed append")
	ErrInvalidName    = errors.New("invalid name")
	ErrWrongType      = errors.New("wrong type")
)
