package errs

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrWrongType       = errors.New("wrong type")
	ErrNotFound        = errors.New("not found")
	ErrCommandNil      = errors.New("command is nil")
	ErrExecuteNil      = errors.New("execute is nil")
	ErrCommandDplicate = errors.New("duplicate command")
)

type AlternativesError struct {
	Alt []string
	Err error
}

func (e AlternativesError) Error() string {
	return fmt.Sprintf("%s, alternatives: [%s]", e.Err, strings.Join(e.Alt, ","))
}

func (e AlternativesError) Is(err error) bool {
	return errors.Is(e.Err, err)
}

func (e AlternativesError) Unwrap() error {
	return e.Err
}
