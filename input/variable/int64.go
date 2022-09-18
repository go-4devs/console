package variable

import (
	"fmt"
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Int64(name, description string, opts ...Option) Variable {
	return String(name, description, append(opts, WithParse(CreateInt64, AppendInt64), Value(flag.Int64))...)
}

func CreateInt64(in string) (value.Value, error) {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("create int64:%w", err)
	}

	return value.NewInt64(out), nil
}

func AppendInt64(old value.Value, in string) (value.Value, error) {
	out, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("append int64:%w", err)
	}

	return value.NewInt64s(append(old.Int64s(), out)), nil
}
