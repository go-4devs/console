package validator

import (
	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

//nolint: gocyclo
func NotBlank(f flag.Flag) func(value.Value) error {
	return func(in value.Value) error {
		switch {
		case f.IsAny() && in.Any() != nil:
			return nil
		case f.IsArray():
			return arrayNotBlank(f, in)
		case f.IsInt() && in.Int() != 0:
			return nil
		case f.IsInt64() && in.Int64() != 0:
			return nil
		case f.IsUint() && in.Uint() != 0:
			return nil
		case f.IsUint64() && in.Uint64() != 0:
			return nil
		case f.IsFloat64() && in.Float64() != 0:
			return nil
		case f.IsDuration() && in.Duration() != 0:
			return nil
		case f.IsTime() && !in.Time().IsZero():
			return nil
		case f.IsString() && len(in.String()) > 0:
			return nil
		}

		return ErrNotBlank
	}
}

//nolint: gocyclo,gocognit
func arrayNotBlank(f flag.Flag, in value.Value) error {
	switch {
	case f.IsInt() && len(in.Ints()) > 0:
		for _, i := range in.Ints() {
			if i == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case f.IsInt64() && len(in.Int64s()) > 0:
		for _, i := range in.Int64s() {
			if i == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case f.IsUint() && len(in.Uints()) > 0:
		for _, u := range in.Uints() {
			if u == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case f.IsUint64() && len(in.Uint64s()) > 0:
		for _, u := range in.Uint64s() {
			if u == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case f.IsFloat64() && len(in.Float64s()) > 0:
		for _, f := range in.Float64s() {
			if f == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case f.IsBool() && len(in.Bools()) > 0:
		return nil
	case f.IsDuration() && len(in.Durations()) > 0:
		for _, d := range in.Durations() {
			if d == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case f.IsTime() && len(in.Times()) > 0:
		for _, t := range in.Times() {
			if t.IsZero() {
				return ErrNotBlank
			}
		}

		return nil
	case f.IsString() && len(in.Strings()) > 0:
		for _, st := range in.Strings() {
			if len(st) == 0 {
				return ErrNotBlank
			}
		}

		return nil
	}

	return ErrNotBlank
}
