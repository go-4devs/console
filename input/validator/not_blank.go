package validator

import (
	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

//nolint:gocyclo,cyclop
func NotBlank(fl flag.Flag) func(value.Value) error {
	return func(in value.Value) error {
		switch {
		case fl.IsAny() && in.Any() != nil:
			return nil
		case fl.IsArray():
			return arrayNotBlank(fl, in)
		case fl.IsInt() && in.Int() != 0:
			return nil
		case fl.IsInt64() && in.Int64() != 0:
			return nil
		case fl.IsUint() && in.Uint() != 0:
			return nil
		case fl.IsUint64() && in.Uint64() != 0:
			return nil
		case fl.IsFloat64() && in.Float64() != 0:
			return nil
		case fl.IsDuration() && in.Duration() != 0:
			return nil
		case fl.IsTime() && !in.Time().IsZero():
			return nil
		case fl.IsString() && len(in.String()) > 0:
			return nil
		}

		return ErrNotBlank
	}
}

//nolint:gocyclo,gocognit,cyclop
func arrayNotBlank(fl flag.Flag, in value.Value) error {
	switch {
	case fl.IsInt() && len(in.Ints()) > 0:
		for _, i := range in.Ints() {
			if i == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case fl.IsInt64() && len(in.Int64s()) > 0:
		for _, i := range in.Int64s() {
			if i == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case fl.IsUint() && len(in.Uints()) > 0:
		for _, u := range in.Uints() {
			if u == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case fl.IsUint64() && len(in.Uint64s()) > 0:
		for _, u := range in.Uint64s() {
			if u == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case fl.IsFloat64() && len(in.Float64s()) > 0:
		for _, f := range in.Float64s() {
			if f == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case fl.IsBool() && len(in.Bools()) > 0:
		return nil
	case fl.IsDuration() && len(in.Durations()) > 0:
		for _, d := range in.Durations() {
			if d == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case fl.IsTime() && len(in.Times()) > 0:
		for _, t := range in.Times() {
			if t.IsZero() {
				return ErrNotBlank
			}
		}

		return nil
	case fl.IsString() && len(in.Strings()) > 0:
		for _, st := range in.Strings() {
			if len(st) == 0 {
				return ErrNotBlank
			}
		}

		return nil
	}

	return ErrNotBlank
}
