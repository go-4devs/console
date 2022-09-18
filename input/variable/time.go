package variable

import (
	"fmt"
	"time"

	"gitoa.ru/go-4devs/console/input/errs"
	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

const (
	ParamFormat = "format"
)

func Time(name, description string, opts ...Option) Variable {
	return String(name, description, append(opts,
		WithParamParse(CreateTime, AppendTime),
		WithParam(ParamFormat, RFC3339),
		Value(flag.Time),
	)...)
}

func RFC3339(in interface{}) error {
	v, ok := in.(*string)
	if !ok {
		return fmt.Errorf("%w: expect *string got %T", errs.ErrWrongType, in)
	}

	*v = time.RFC3339

	return nil
}

func CreateTime(param Param) Create {
	var (
		formatErr error
		format    string
	)

	formatErr = param.Value(ParamFormat, &format)

	return func(in string) (value.Value, error) {
		if formatErr != nil {
			return nil, fmt.Errorf("create format:%w", formatErr)
		}

		out, err := time.Parse(format, in)
		if err != nil {
			return nil, fmt.Errorf("create time:%w", err)
		}

		return value.NewTime(out), nil
	}
}

func AppendTime(param Param) Append {
	var (
		formatErr error
		format    string
	)

	formatErr = param.Value(ParamFormat, &format)

	return func(old value.Value, in string) (value.Value, error) {
		if formatErr != nil {
			return nil, fmt.Errorf("append format:%w", formatErr)
		}

		out, err := time.Parse(format, in)
		if err != nil {
			return nil, fmt.Errorf("append time:%w", err)
		}

		return value.NewTimes(append(old.Times(), out)), nil
	}
}
