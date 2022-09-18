package validator

import "gitoa.ru/go-4devs/console/input/value"

func Enum(enum ...string) func(value.Value) error {
	return func(in value.Value) error {
		val := in.String()
		for _, e := range enum {
			if e == val {
				return nil
			}
		}

		return NewError(ErrInvalid, val, enum)
	}
}
