package variable

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Duration(name, description string, opts ...Option) Variable {
	return String(name, description, append(opts, WithParse(CreateDuration, AppendDuration), Value(flag.Duration))...)
}

func CreateDuration(in string) (value.Value, error) {
	out, err := time.ParseDuration(in)
	if err != nil {
		return nil, fmt.Errorf("create duration:%w", err)
	}

	return value.NewDuration(out), nil
}

func AppendDuration(old value.Value, in string) (value.Value, error) {
	out, err := time.ParseDuration(in)
	if err != nil {
		return nil, fmt.Errorf("append duration:%w", err)
	}

	return value.NewDurations(append(old.Durations(), out)), nil
}
