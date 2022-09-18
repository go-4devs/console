package value

import (
	"fmt"
	"time"
)

var (
	_ ParseValue = Float64(0)
	_ SliceValue = Float64s{}
)

func NewFloat64s(in []float64) Slice {
	return Slice{SliceValue: Float64s(in)}
}

type Float64s []float64

func (f Float64s) Any() interface{} {
	return f.Float64s()
}

func (f Float64s) Unmarshal(val interface{}) error {
	v, ok := val.(*[]float64)
	if !ok {
		return fmt.Errorf("%w: expect *[]float64", ErrWrongType)
	}

	*v = f

	return nil
}

func (f Float64s) Strings() []string {
	return nil
}

func (f Float64s) Ints() []int {
	return nil
}

func (f Float64s) Int64s() []int64 {
	return nil
}

func (f Float64s) Uints() []uint {
	return nil
}

func (f Float64s) Uint64s() []uint64 {
	return nil
}

func (f Float64s) Float64s() []float64 {
	out := make([]float64, len(f))
	copy(out, f)

	return out
}

func (f Float64s) Bools() []bool {
	return nil
}

func (f Float64s) Durations() []time.Duration {
	return nil
}

func (f Float64s) Times() []time.Time {
	return nil
}

func NewFloat64(in float64) Read {
	return Read{ParseValue: Float64(in)}
}

type Float64 float64

func (f Float64) Any() interface{} {
	return float64(f)
}

func (f Float64) ParseString() (string, error) {
	return fmt.Sprint(float64(f)), nil
}

func (f Float64) ParseInt() (int, error) {
	return int(f), nil
}

func (f Float64) ParseInt64() (int64, error) {
	return int64(f), nil
}

func (f Float64) ParseUint() (uint, error) {
	return uint(f), nil
}

func (f Float64) ParseUint64() (uint64, error) {
	return uint64(f), nil
}

func (f Float64) ParseFloat64() (float64, error) {
	return float64(f), nil
}

func (f Float64) ParseBool() (bool, error) {
	return false, fmt.Errorf("float64:%w", ErrWrongType)
}

func (f Float64) ParseDuration() (time.Duration, error) {
	return time.Duration(f), nil
}

func (f Float64) ParseTime() (time.Time, error) {
	return time.Unix(0, int64(f*Float64(time.Second))), nil
}

func (f Float64) Unmarshal(in interface{}) error {
	v, ok := in.(*float64)
	if !ok {
		return fmt.Errorf("%w: expect *float64", ErrWrongType)
	}

	*v = float64(f)

	return nil
}
