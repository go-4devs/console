package variable

import (
	"gitoa.ru/go-4devs/console/input/value"
)

func String(name, description string, opts ...Option) Variable {
	return New(name, description, opts...)
}

func CreateString(in string) (value.Value, error) {
	return value.NewString(in), nil
}

func AppendString(old value.Value, in string) (value.Value, error) {
	return value.NewStrings(append(old.Strings(), in)), nil
}
