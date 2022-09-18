package value

import (
	"fmt"
	"time"
)

var (
	_ ParseValue = Duration(0)
	_ SliceValue = Durations{}
)

func NewDurations(in []time.Duration) Slice {
	return Slice{SliceValue: Durations(in)}
}

type Durations []time.Duration

func (d Durations) Unmarshal(val interface{}) error {
	v, ok := val.(*[]time.Duration)
	if !ok {
		return fmt.Errorf("%w: expect: *[]time.Duration got: %T", ErrWrongType, val)
	}

	*v = d

	return nil
}

func (d Durations) Any() interface{} {
	return d.Durations()
}

func (d Durations) Strings() []string {
	return nil
}

func (d Durations) Ints() []int {
	return nil
}

func (d Durations) Int64s() []int64 {
	return nil
}

func (d Durations) Uints() []uint {
	return nil
}

func (d Durations) Uint64s() []uint64 {
	return nil
}

func (d Durations) Float64s() []float64 {
	return nil
}

func (d Durations) Bools() []bool {
	return nil
}

func (d Durations) Durations() []time.Duration {
	out := make([]time.Duration, len(d))
	copy(out, d)

	return out
}

func (d Durations) Times() []time.Time {
	return nil
}

func NewDuration(in time.Duration) Read {
	return Read{ParseValue: Duration(in)}
}

type Duration time.Duration

func (d Duration) ParseDuration() (time.Duration, error) {
	return time.Duration(d), nil
}

func (d Duration) ParseString() (string, error) {
	return time.Duration(d).String(), nil
}

func (d Duration) ParseInt() (int, error) {
	return int(d), nil
}

func (d Duration) ParseInt64() (int64, error) {
	return int64(d), nil
}

func (d Duration) ParseUint() (uint, error) {
	return uint(d), nil
}

func (d Duration) ParseUint64() (uint64, error) {
	return uint64(d), nil
}

func (d Duration) ParseFloat64() (float64, error) {
	return float64(d), nil
}

func (d Duration) ParseBool() (bool, error) {
	return false, fmt.Errorf("duration:%w", ErrWrongType)
}

func (d Duration) ParseTime() (time.Time, error) {
	return time.Time{}, fmt.Errorf("duration:%w", ErrWrongType)
}

func (d Duration) Unmarshal(val interface{}) error {
	v, ok := val.(*time.Duration)
	if !ok {
		return fmt.Errorf("%w: expect: *[]time.Duration got: %T", ErrWrongType, val)
	}

	*v = time.Duration(d)

	return nil
}

func (d Duration) Any() interface{} {
	return time.Duration(d)
}
