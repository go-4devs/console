package validator

import "gitoa.ru/go-4devs/console/input"

func Enum(enum ...string) func(input.Value) error {
	return func(in input.Value) error {
		v := in.String()
		for _, e := range enum {
			if e == v {
				return nil
			}
		}

		return NewError(ErrInvalid, v, enum)
	}
}
