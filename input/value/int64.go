package value

import (
	"fmt"
	"strconv"
	"time"
)

var (
	_ ParseValue = Int64(0)
	_ SliceValue = Int64s{}
)

func NewInt64s(in []int64) Slice {
	return Slice{SliceValue: Int64s(in)}
}

type Int64s []int64

func (i Int64s) Any() interface{} {
	return i.Int64s()
}

func (i Int64s) Unmarshal(val interface{}) error {
	v, ok := val.(*[]int64)
	if !ok {
		return fmt.Errorf("%w: expect *[]int64", ErrWrongType)
	}

	*v = i

	return nil
}

func (i Int64s) Strings() []string {
	return nil
}

func (i Int64s) Ints() []int {
	return nil
}

func (i Int64s) Int64s() []int64 {
	out := make([]int64, len(i))
	copy(out, i)

	return out
}

func (i Int64s) Uints() []uint {
	return nil
}

func (i Int64s) Uint64s() []uint64 {
	return nil
}

func (i Int64s) Float64s() []float64 {
	return nil
}

func (i Int64s) Bools() []bool {
	return nil
}

func (i Int64s) Durations() []time.Duration {
	return nil
}

func (i Int64s) Times() []time.Time {
	return nil
}

func NewInt64(in int64) Read {
	return Read{ParseValue: Int64(in)}
}

type Int64 int64

func (i Int64) Any() interface{} {
	return int64(i)
}

func (i Int64) ParseString() (string, error) {
	return strconv.FormatInt(int64(i), 10), nil
}

func (i Int64) ParseInt() (int, error) {
	return int(i), nil
}

func (i Int64) ParseInt64() (int64, error) {
	return int64(i), nil
}

func (i Int64) ParseUint() (uint, error) {
	return uint(i), nil
}

func (i Int64) ParseUint64() (uint64, error) {
	return uint64(i), nil
}

func (i Int64) ParseFloat64() (float64, error) {
	return float64(i), nil
}

func (i Int64) ParseBool() (bool, error) {
	return false, fmt.Errorf("int64:%w", ErrWrongType)
}

func (i Int64) ParseDuration() (time.Duration, error) {
	return time.Duration(i), nil
}

func (i Int64) ParseTime() (time.Time, error) {
	return time.Unix(0, int64(i)), nil
}

func (i Int64) Unmarshal(val interface{}) error {
	v, ok := val.(*int64)
	if !ok {
		return fmt.Errorf("%w: expect *int64", ErrWrongType)
	}

	*v = int64(i)

	return nil
}
