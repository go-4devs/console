package input

import (
	"time"
)

type Value interface {
	String() string
	Int() int
	Int64() int64
	Uint() uint
	Uint64() uint64
	Float64() float64
	Bool() bool
	Duration() time.Duration
	Time() time.Time
	Any() interface{}

	Strings() []string
	Ints() []int
	Int64s() []int64
	Uints() []uint
	Uint64s() []uint64
	Float64s() []float64
	Bools() []bool
	Durations() []time.Duration
	Times() []time.Time
}

type AppendValue interface {
	Value
	Append(string) error
}

func Type(flag Flag) Flag {
	switch {
	case (flag & ValueInt) > 0:
		return ValueInt
	case (flag & ValueInt64) > 0:
		return ValueInt64
	case (flag & ValueUint) > 0:
		return ValueUint
	case (flag & ValueUint64) > 0:
		return ValueUint64
	case (flag & ValueFloat64) > 0:
		return ValueFloat64
	case (flag & ValueBool) > 0:
		return ValueBool
	case (flag & ValueDuration) > 0:
		return ValueDuration
	case (flag & ValueTime) > 0:
		return ValueTime
	case (flag & ValueAny) > 0:
		return ValueAny
	default:
		return ValueString
	}
}
