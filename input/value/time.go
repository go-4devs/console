package value

import (
	"fmt"
	"time"
)

var (
	_ ParseValue = Time{time.Now()}
	_ SliceValue = (Times)(nil)
)

func NewTimes(in []time.Time) Slice {
	return Slice{SliceValue: Times(in)}
}

type Times []time.Time

func (t Times) Any() interface{} {
	return t.Times()
}

func (t Times) Unmarshal(val interface{}) error {
	res, ok := val.(*[]time.Time)
	if !ok {
		return fmt.Errorf("%w: expect *[]time.Time", ErrWrongType)
	}

	*res = t

	return nil
}

func (t Times) Strings() []string {
	return nil
}

func (t Times) Ints() []int {
	return nil
}

func (t Times) Int64s() []int64 {
	return nil
}

func (t Times) Uints() []uint {
	return nil
}

func (t Times) Uint64s() []uint64 {
	return nil
}

func (t Times) Float64s() []float64 {
	return nil
}

func (t Times) Bools() []bool {
	return nil
}

func (t Times) Durations() []time.Duration {
	return nil
}

func (t Times) Times() []time.Time {
	out := make([]time.Time, len(t))
	copy(out, t)

	return out
}

func NewTime(in time.Time) Read {
	return Read{ParseValue: Time{Time: in}}
}

type Time struct {
	time.Time
}

func (t Time) ParseString() (string, error) {
	return t.Format(time.RFC3339), nil
}

func (t Time) ParseInt() (int, error) {
	return int(t.Unix()), nil
}

func (t Time) ParseInt64() (int64, error) {
	return t.Unix(), nil
}

func (t Time) ParseUint() (uint, error) {
	return uint(t.Unix()), nil
}

func (t Time) ParseUint64() (uint64, error) {
	return uint64(t.Unix()), nil
}

func (t Time) ParseFloat64() (float64, error) {
	return float64(t.UnixNano()), nil
}

func (t Time) ParseBool() (bool, error) {
	return false, fmt.Errorf("time bool:%w", ErrWrongType)
}

func (t Time) ParseDuration() (time.Duration, error) {
	return 0, fmt.Errorf("time duration:%w", ErrWrongType)
}

func (t Time) ParseTime() (time.Time, error) {
	return t.Time, nil
}

func (t Time) Unmarshal(val interface{}) error {
	res, ok := val.(*time.Time)
	if !ok {
		return fmt.Errorf("%w: expect *time.Time", ErrWrongType)
	}

	*res = t.Time

	return nil
}

func (t Time) Any() interface{} {
	return t.Time
}
