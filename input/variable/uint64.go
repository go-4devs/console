package variable

import (
	"fmt"
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Uint64(name, descriontion string, opts ...Option) Variable {
	return String(name, descriontion, append(opts, WithParse(CreateUint64, AppendUint64), Value(flag.Uint64))...)
}

func CreateUint64(in string) (value.Value, error) {
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("create uint64:%w", err)
	}

	return value.NewUint64(out), nil
}

func AppendUint64(old value.Value, in string) (value.Value, error) {
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("append uint64:%w", err)
	}

	return value.NewUint64s(append(old.Uint64s(), out)), nil
}
