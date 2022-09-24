package console

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrNotFound         = errors.New("command not found")
	ErrCommandNil       = errors.New("console: Register command is nil")
	ErrExecuteNil       = errors.New("console: execute is nil")
	ErrCommandDuplicate = errors.New("console: duplicate command")
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
