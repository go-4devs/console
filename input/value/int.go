package value

import (
	"fmt"
	"strconv"
	"time"
)

var (
	_ ParseValue = Int(0)
	_ SliceValue = Ints{}
)

func NewInts(in []int) Slice {
	return Slice{SliceValue: Ints(in)}
}

type Ints []int

func (i Ints) Unmarshal(in interface{}) error {
	val, ok := in.(*[]int)
	if !ok {
		return fmt.Errorf("%w: expect *[]int", ErrWrongType)
	}

	*val = i

	return nil
}

func (i Ints) Any() interface{} {
	return i.Ints()
}

func (i Ints) Strings() []string {
	return nil
}

func (i Ints) Ints() []int {
	out := make([]int, len(i))
	copy(out, i)

	return out
}

func (i Ints) Int64s() []int64 {
	return nil
}

func (i Ints) Uints() []uint {
	return nil
}

func (i Ints) Uint64s() []uint64 {
	return nil
}

func (i Ints) Float64s() []float64 {
	return nil
}

func (i Ints) Bools() []bool {
	return nil
}

func (i Ints) Durations() []time.Duration {
	return nil
}

func (i Ints) Times() []time.Time {
	return nil
}

func NewInt(in int) Read {
	return Read{ParseValue: Int(in)}
}

type Int int

func (i Int) Unmarshal(in interface{}) error {
	v, ok := in.(*int)
	if !ok {
		return fmt.Errorf("%w: expect *int", ErrWrongType)
	}

	*v = int(i)

	return nil
}

func (i Int) ParseString() (string, error) {
	return strconv.Itoa(int(i)), nil
}

func (i Int) ParseInt() (int, error) {
	return int(i), nil
}

func (i Int) ParseInt64() (int64, error) {
	return int64(i), nil
}

func (i Int) ParseUint() (uint, error) {
	return uint(i), nil
}

func (i Int) ParseUint64() (uint64, error) {
	return uint64(i), nil
}

func (i Int) ParseFloat64() (float64, error) {
	return float64(i), nil
}

func (i Int) ParseBool() (bool, error) {
	return false, fmt.Errorf("int:%w", ErrWrongType)
}

func (i Int) ParseDuration() (time.Duration, error) {
	return time.Duration(i), nil
}

func (i Int) ParseTime() (time.Time, error) {
	return time.Unix(0, int64(i)), nil
}

func (i Int) Any() interface{} {
	return int(i)
}
