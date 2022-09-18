package value

import (
	"fmt"
	"time"
)

var (
	_ ParseValue = Bool(false)
	_ SliceValue = Bools{}
)

func NewBools(in []bool) Slice {
	return Slice{SliceValue: Bools(in)}
}

type Bools []bool

func (b Bools) Any() interface{} {
	return b.Bools()
}

func (b Bools) Unmarshal(val interface{}) error {
	v, ok := val.(*[]bool)
	if !ok {
		return fmt.Errorf("%w: expect: *[]bool got: %T", ErrWrongType, val)
	}

	*v = b

	return nil
}

func (b Bools) Strings() []string {
	return nil
}

func (b Bools) Ints() []int {
	return nil
}

func (b Bools) Int64s() []int64 {
	return nil
}

func (b Bools) Uints() []uint {
	return nil
}

func (b Bools) Uint64s() []uint64 {
	return nil
}

func (b Bools) Float64s() []float64 {
	return nil
}

func (b Bools) Bools() []bool {
	out := make([]bool, len(b))
	copy(out, b)

	return out
}

func (b Bools) Durations() []time.Duration {
	return nil
}

func (b Bools) Times() []time.Time {
	return nil
}

func NewBool(in bool) Read {
	return Read{ParseValue: Bool(in)}
}

type Bool bool

func (b Bool) Unmarshal(val interface{}) error {
	v, ok := val.(*bool)
	if !ok {
		return fmt.Errorf("%w: expect: *bool got: %T", ErrWrongType, val)
	}

	*v = bool(b)

	return nil
}

func (b Bool) ParseString() (string, error) {
	return fmt.Sprintf("%v", b), nil
}

func (b Bool) ParseInt() (int, error) {
	if b {
		return 1, nil
	}

	return 0, nil
}

func (b Bool) ParseInt64() (int64, error) {
	if b {
		return 1, nil
	}

	return 0, nil
}

func (b Bool) ParseUint() (uint, error) {
	if b {
		return 1, nil
	}

	return 0, nil
}

func (b Bool) ParseUint64() (uint64, error) {
	if b {
		return 1, nil
	}

	return 0, nil
}

func (b Bool) ParseFloat64() (float64, error) {
	if b {
		return 1, nil
	}

	return 0, nil
}

func (b Bool) ParseBool() (bool, error) {
	return bool(b), nil
}

func (b Bool) ParseDuration() (time.Duration, error) {
	return 0, fmt.Errorf("bool to duration:%w", ErrWrongType)
}

func (b Bool) ParseTime() (time.Time, error) {
	return time.Time{}, fmt.Errorf("bool to time:%w", ErrWrongType)
}

func (b Bool) Any() interface{} {
	return bool(b)
}
