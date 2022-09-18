package value

import (
	"time"
)

type Value interface {
	ReadValue
	ParseValue
	ArrValue
}

type UnmarshalValue interface {
	Unmarshal(val interface{}) error
}

type ReadValue interface {
	String() string
	Int() int
	Int64() int64
	Uint() uint
	Uint64() uint64
	Float64() float64
	Bool() bool
	Duration() time.Duration
	Time() time.Time
}

type AnyValue interface {
	Any() interface{}
}

type SliceValue interface {
	AnyValue
	UnmarshalValue
	ArrValue
}

type ArrValue interface {
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

//nolint:interfacebloat
type ParseValue interface {
	ParseString() (string, error)
	ParseInt() (int, error)
	ParseInt64() (int64, error)
	ParseUint() (uint, error)
	ParseUint64() (uint64, error)
	ParseFloat64() (float64, error)
	ParseBool() (bool, error)
	ParseDuration() (time.Duration, error)
	ParseTime() (time.Time, error)
	UnmarshalValue
	AnyValue
}

type Append interface {
	Value
	Append(string) (Value, error)
}

//nolint:gocyclo,cyclop
func New(in interface{}) Value {
	switch val := in.(type) {
	case bool:
		return Read{Bool(val)}
	case []bool:
		return NewBools(val)
	case string:
		return Read{String(val)}
	case int:
		return Read{Int(val)}
	case int64:
		return Read{Int64(val)}
	case uint:
		return Read{Uint(val)}
	case uint64:
		return Read{Uint64(val)}
	case float64:
		return Read{Float64(val)}
	case time.Duration:
		return Read{Duration(val)}
	case time.Time:
		return Read{Time{val}}
	case []int64:
		return Slice{Int64s(val)}
	case []uint:
		return Slice{Uints(val)}
	case []uint64:
		return Slice{Uint64s(val)}
	case []float64:
		return Slice{Float64s(val)}
	case []time.Duration:
		return Slice{Durations(val)}
	case []time.Time:
		return Slice{Times(val)}
	case []string:
		return Slice{Strings(val)}
	case []int:
		return Slice{Ints(val)}
	case []interface{}:
		return Read{Any{v: val}}
	case Value:
		return val
	default:
		if in != nil {
			return Read{Any{v: in}}
		}

		return Empty()
	}
}
