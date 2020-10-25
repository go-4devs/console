package value

import (
	"time"
)

type Empty struct{}

func (e *Empty) Append(string) error {
	return ErrAppendEmpty
}

func (e *Empty) String() string {
	return ""
}

func (e *Empty) Int() int {
	return 0
}

func (e *Empty) Int64() int64 {
	return 0
}

func (e *Empty) Uint() uint {
	return 0
}

func (e *Empty) Uint64() uint64 {
	return 0
}

func (e *Empty) Float64() float64 {
	return 0
}

func (e *Empty) Bool() bool {
	return false
}

func (e *Empty) Duration() time.Duration {
	return 0
}

func (e *Empty) Time() time.Time {
	return time.Time{}
}

func (e *Empty) Strings() []string {
	return nil
}

func (e *Empty) Ints() []int {
	return nil
}

func (e *Empty) Int64s() []int64 {
	return nil
}

func (e *Empty) Uints() []uint {
	return nil
}

func (e *Empty) Uint64s() []uint64 {
	return nil
}

func (e *Empty) Float64s() []float64 {
	return nil
}

func (e *Empty) Bools() []bool {
	return nil
}

func (e *Empty) Durations() []time.Duration {
	return nil
}

func (e *Empty) Times() []time.Time {
	return nil
}

func (e *Empty) Any() interface{} {
	return nil
}
