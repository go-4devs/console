package variable

import (
	"errors"
	"fmt"
)

type Error struct {
	Name string
	Err  error
	Type ArgType
}

func (o Error) Error() string {
	return fmt.Sprintf("%s: '%s' %s", o.Type, o.Name, o.Err)
}

func (o Error) Is(err error) bool {
	return errors.Is(err, o.Err)
}

func (o Error) Unwrap() error {
	return o.Err
}

func Err(name string, t ArgType, err error) Error {
	return Error{
		Name: name,
		Type: t,
		Err:  err,
	}
}
