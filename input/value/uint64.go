package value

import (
	"fmt"
	"strconv"
	"time"
)

var (
	_ ParseValue = Uint64(0)
	_ SliceValue = (Uint64s)(nil)
)

func NewUint64s(in []uint64) Slice {
	return Slice{SliceValue: Uint64s(in)}
}

type Uint64s []uint64

func (u Uint64s) Any() interface{} {
	return u.Uint64s()
}

func (u Uint64s) Unmarshal(val interface{}) error {
	res, ok := val.(*[]uint64)
	if !ok {
		return fmt.Errorf("%w: expect *[]uint64", ErrWrongType)
	}

	*res = u

	return nil
}

func (u Uint64s) Strings() []string {
	return nil
}

func (u Uint64s) Ints() []int {
	return nil
}

func (u Uint64s) Int64s() []int64 {
	return nil
}

func (u Uint64s) Uints() []uint {
	return nil
}

func (u Uint64s) Uint64s() []uint64 {
	out := make([]uint64, len(u))
	copy(out, u)

	return out
}

func (u Uint64s) Float64s() []float64 {
	return nil
}

func (u Uint64s) Bools() []bool {
	return nil
}

func (u Uint64s) Durations() []time.Duration {
	return nil
}

func (u Uint64s) Times() []time.Time {
	return nil
}

func NewUint64(in uint64) Read {
	return Read{ParseValue: Uint64(in)}
}

type Uint64 uint64

func (u Uint64) ParseString() (string, error) {
	return strconv.FormatUint(uint64(u), 10), nil
}

func (u Uint64) ParseInt() (int, error) {
	return int(u), nil
}

func (u Uint64) ParseInt64() (int64, error) {
	return int64(u), nil
}

func (u Uint64) ParseUint() (uint, error) {
	return uint(u), nil
}

func (u Uint64) ParseUint64() (uint64, error) {
	return uint64(u), nil
}

func (u Uint64) ParseFloat64() (float64, error) {
	return float64(u), nil
}

func (u Uint64) ParseBool() (bool, error) {
	return false, fmt.Errorf("uint64 bool:%w", ErrWrongType)
}

func (u Uint64) ParseDuration() (time.Duration, error) {
	return time.Duration(u), nil
}

func (u Uint64) ParseTime() (time.Time, error) {
	return time.Unix(0, int64(0)), nil
}

func (u Uint64) Unmarshal(val interface{}) error {
	res, ok := val.(*uint64)
	if !ok {
		return fmt.Errorf("%w: expect *uint64", ErrWrongType)
	}

	*res = uint64(u)

	return nil
}

func (u Uint64) Any() interface{} {
	return uint64(u)
}
