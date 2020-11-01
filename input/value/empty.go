package value

import (
	"time"
)

// nolint: gochecknoglobals
var (
	emptyValue = &empty{}
)

func Empty() Value {
	return emptyValue
}

func IsEmpty(v Value) bool {
	return v == nil || v == emptyValue
}

type empty struct{}

func (e *empty) Append(string) error {
	return ErrAppendEmpty
}

func (e *empty) String() string {
	return ""
}

func (e *empty) Int() int {
	return 0
}

func (e *empty) Int64() int64 {
	return 0
}

func (e *empty) Uint() uint {
	return 0
}

func (e *empty) Uint64() uint64 {
	return 0
}

func (e *empty) Float64() float64 {
	return 0
}

func (e *empty) Bool() bool {
	return false
}

func (e *empty) Duration() time.Duration {
	return 0
}

func (e *empty) Time() time.Time {
	return time.Time{}
}

func (e *empty) Strings() []string {
	return nil
}

func (e *empty) Ints() []int {
	return nil
}

func (e *empty) Int64s() []int64 {
	return nil
}

func (e *empty) Uints() []uint {
	return nil
}

func (e *empty) Uint64s() []uint64 {
	return nil
}

func (e *empty) Float64s() []float64 {
	return nil
}

func (e *empty) Bools() []bool {
	return nil
}

func (e *empty) Durations() []time.Duration {
	return nil
}

func (e *empty) Times() []time.Time {
	return nil
}

func (e *empty) Any() interface{} {
	return nil
}
