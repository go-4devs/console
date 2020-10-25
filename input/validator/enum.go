package validator

import "gitoa.ru/go-4devs/console/input/value"

func Enum(enum ...string) func(value.Value) error {
	return func(in value.Value) error {
		v := in.String()
		for _, e := range enum {
			if e == v {
				return nil
			}
		}

		return NewError(ErrInvalid, v, enum)
	}
}
