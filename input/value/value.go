package value

import (
	"time"

	"gitoa.ru/go-4devs/console/input/value/flag"
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

type Append interface {
	Value
	Append(string) error
}

//nolint: gocyclo
func New(v interface{}) Append {
	switch val := v.(type) {
	case string:
		return &String{Val: []string{val}, Flag: flag.String}
	case int:
		return &Int{Val: []int{val}, Flag: flag.Int}
	case int64:
		return &Int64{Val: []int64{val}, Flag: flag.Int64}
	case uint:
		return &Uint{Val: []uint{val}, Flag: flag.Uint}
	case uint64:
		return &Uint64{Val: []uint64{val}, Flag: flag.Uint64}
	case float64:
		return &Float64{Val: []float64{val}, Flag: flag.Float64}
	case bool:
		return &Bool{Val: []bool{val}, Flag: flag.Bool}
	case time.Duration:
		return &Duration{Val: []time.Duration{val}, Flag: flag.Duration}
	case time.Time:
		return &Time{Val: []time.Time{val}, Flag: flag.Time}
	case []int64:
		return &Int64{Val: val, Flag: flag.Int64 | flag.Array}
	case []uint:
		return &Uint{Val: val, Flag: flag.Uint | flag.Array}
	case []uint64:
		return &Uint64{Val: val, Flag: flag.Uint64 | flag.Array}
	case []float64:
		return &Float64{Val: val, Flag: flag.Float64 | flag.Array}
	case []bool:
		return &Bool{Val: val, Flag: flag.Bool | flag.Array}
	case []time.Duration:
		return &Duration{Val: val, Flag: flag.Duration | flag.Array}
	case []time.Time:
		return &Time{Val: val, Flag: flag.Time | flag.Array}
	case []string:
		return &String{Val: val, Flag: flag.String | flag.Array}
	case []int:
		return &Int{Val: val, Flag: flag.Int | flag.Array}
	case []interface{}:
		return &Any{Val: val, Flag: flag.Any | flag.Array}
	case Append:
		return val
	case Value:
		return &Read{Value: val}
	default:
		if v != nil {
			return &Any{Val: []interface{}{v}, Flag: flag.Any}
		}

		return &empty{}
	}
}

func ByFlag(f flag.Flag) Append {
	switch {
	case f.IsInt():
		return &Int{Flag: f | flag.Int}
	case f.IsInt64():
		return &Int64{Flag: f | flag.Int64}
	case f.IsUint():
		return &Uint{Flag: f | flag.Uint}
	case f.IsUint64():
		return &Uint64{Flag: f | flag.Uint64}
	case f.IsFloat64():
		return &Float64{Flag: f | flag.Float64}
	case f.IsBool():
		return &Bool{Flag: f | flag.Bool}
	case f.IsDuration():
		return &Duration{Flag: f | flag.Duration}
	case f.IsTime():
		return &Time{Flag: f | flag.Time}
	case f.IsAny():
		return &Any{Flag: f | flag.Any}
	default:
		return &String{}
	}
}
