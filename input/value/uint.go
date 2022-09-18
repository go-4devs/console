package value

import (
	"fmt"
	"strconv"
	"time"
)

var (
	_ ParseValue = Uint(0)
	_ SliceValue = (Uints)(nil)
)

func NewUints(in []uint) Slice {
	return Slice{SliceValue: Uints(in)}
}

type Uints []uint

func (u Uints) Any() interface{} {
	return u.Uints()
}

func (u Uints) Unmarshal(val interface{}) error {
	res, ok := val.(*[]uint)
	if !ok {
		return fmt.Errorf("%w: expect *[]uint", ErrWrongType)
	}

	*res = u

	return nil
}

func (u Uints) Strings() []string {
	return nil
}

func (u Uints) Ints() []int {
	return nil
}

func (u Uints) Int64s() []int64 {
	return nil
}

func (u Uints) Uints() []uint {
	out := make([]uint, len(u))
	copy(out, u)

	return out
}

func (u Uints) Uint64s() []uint64 {
	return nil
}

func (u Uints) Float64s() []float64 {
	return nil
}

func (u Uints) Bools() []bool {
	return nil
}

func (u Uints) Durations() []time.Duration {
	return nil
}

func (u Uints) Times() []time.Time {
	return nil
}

func NewUint(in uint) Read {
	return Read{ParseValue: Uint(in)}
}

type Uint uint

func (u Uint) ParseString() (string, error) {
	return strconv.FormatUint(uint64(u), 10), nil
}

func (u Uint) ParseInt() (int, error) {
	return int(u), nil
}

func (u Uint) ParseInt64() (int64, error) {
	return int64(u), nil
}

func (u Uint) ParseUint() (uint, error) {
	return uint(u), nil
}

func (u Uint) ParseUint64() (uint64, error) {
	return uint64(u), nil
}

func (u Uint) ParseFloat64() (float64, error) {
	return float64(u), nil
}

func (u Uint) ParseBool() (bool, error) {
	return false, fmt.Errorf("uint:%w", ErrWrongType)
}

func (u Uint) ParseDuration() (time.Duration, error) {
	return time.Duration(u), nil
}

func (u Uint) ParseTime() (time.Time, error) {
	return time.Unix(0, int64(u)), nil
}

func (u Uint) Unmarshal(val interface{}) error {
	res, ok := val.(*uint)
	if !ok {
		return fmt.Errorf("%w: expect *uint", ErrWrongType)
	}

	*res = uint(u)

	return nil
}

func (u Uint) Any() interface{} {
	return uint(u)
}
