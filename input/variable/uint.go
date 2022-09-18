package variable

import (
	"fmt"
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Uint(name, description string, opts ...Option) Variable {
	return String(name, description, append(opts, WithParse(CreateUint, AppendUint), Value(flag.Uint))...)
}

func CreateUint(in string) (value.Value, error) {
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("create uint:%w", err)
	}

	return value.NewUint(uint(out)), nil
}

func AppendUint(old value.Value, in string) (value.Value, error) {
	out, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("append uint:%w", err)
	}

	return value.NewUints(append(old.Uints(), uint(out))), nil
}
