package variable

import (
	"fmt"
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Int(name, description string, opts ...Option) Variable {
	return New(name, description, append(opts, WithParse(CreateInt, AppendInt), Value(flag.Int))...)
}

func AppendInt(old value.Value, in string) (value.Value, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return nil, fmt.Errorf("append int:%w", err)
	}

	return value.NewInts(append(old.Ints(), out)), nil
}

func CreateInt(in string) (value.Value, error) {
	out, err := strconv.Atoi(in)
	if err != nil {
		return nil, fmt.Errorf("create int:%w", err)
	}

	return value.NewInt(out), nil
}
