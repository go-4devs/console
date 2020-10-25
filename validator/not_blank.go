package validator

import (
	"gitoa.ru/go-4devs/console/input"
)

//nolint: gocyclo
func NotBlank(flag input.Flag) func(input.Value) error {
	return func(in input.Value) error {
		switch {
		case flag.IsAny() && in.Any() != nil:
			return nil
		case flag.IsArray():
			return arrayNotBlank(flag, in)
		case flag.IsInt() && in.Int() != 0:
			return nil
		case flag.IsInt64() && in.Int64() != 0:
			return nil
		case flag.IsUint() && in.Uint() != 0:
			return nil
		case flag.IsUint64() && in.Uint64() != 0:
			return nil
		case flag.IsFloat64() && in.Float64() != 0:
			return nil
		case flag.IsDuration() && in.Duration() != 0:
			return nil
		case flag.IsTime() && !in.Time().IsZero():
			return nil
		case flag.IsString() && len(in.String()) > 0:
			return nil
		}

		return ErrNotBlank
	}
}

//nolint: gocyclo,gocognit
func arrayNotBlank(flag input.Flag, in input.Value) error {
	switch {
	case flag.IsInt() && len(in.Ints()) > 0:
		for _, i := range in.Ints() {
			if i == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case flag.IsInt64() && len(in.Int64s()) > 0:
		for _, i := range in.Int64s() {
			if i == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case flag.IsUint() && len(in.Uints()) > 0:
		for _, u := range in.Uints() {
			if u == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case flag.IsUint64() && len(in.Uint64s()) > 0:
		for _, u := range in.Uint64s() {
			if u == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case flag.IsFloat64() && len(in.Float64s()) > 0:
		for _, f := range in.Float64s() {
			if f == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case flag.IsBool() && len(in.Bools()) > 0:
		return nil
	case flag.IsDuration() && len(in.Durations()) > 0:
		for _, d := range in.Durations() {
			if d == 0 {
				return ErrNotBlank
			}
		}

		return nil
	case flag.IsTime() && len(in.Times()) > 0:
		for _, t := range in.Times() {
			if t.IsZero() {
				return ErrNotBlank
			}
		}

		return nil
	case flag.IsString() && len(in.Strings()) > 0:
		for _, st := range in.Strings() {
			if len(st) == 0 {
				return ErrNotBlank
			}
		}

		return nil
	}

	return ErrNotBlank
}
