package variable

import (
	"fmt"
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Float64(name, description string, opts ...Option) Variable {
	return String(name, description, append(opts, WithParse(CreateFloat64, AppendFloat64), Value(flag.Float64))...)
}

func CreateFloat64(in string) (value.Value, error) {
	out, err := strconv.ParseFloat(in, 10)
	if err != nil {
		return nil, fmt.Errorf("create float64:%w", err)
	}

	return value.NewFloat64(out), nil
}

func AppendFloat64(old value.Value, in string) (value.Value, error) {
	out, err := strconv.ParseFloat(in, 10)
	if err != nil {
		return nil, fmt.Errorf("append float64:%w", err)
	}

	return value.NewFloat64s(append(old.Float64s(), out)), nil
}
