package variable

import (
	"fmt"
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Bool(name, description string, opts ...Option) Variable {
	return String(name, description, append(opts, WithParse(CreateBool, AppendBool), Value(flag.Bool))...)
}

func CreateBool(in string) (value.Value, error) {
	out, err := strconv.ParseBool(in)
	if err != nil {
		return nil, fmt.Errorf("create bool:%w", err)
	}

	return value.NewBool(out), nil
}

func AppendBool(old value.Value, in string) (value.Value, error) {
	out, err := strconv.ParseBool(in)
	if err != nil {
		return nil, fmt.Errorf("create bool:%w", err)
	}

	return value.NewBools(append(old.Bools(), out)), nil
}
