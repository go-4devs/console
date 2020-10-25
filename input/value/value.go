package value

import (
	"time"

	"gitoa.ru/go-4devs/console/input"
)

//nolint: gocyclo
func New(v interface{}) input.Value {
	switch val := v.(type) {
	case string:
		return &String{Val: []string{val}, Flag: input.ValueString}
	case int:
		return &Int{Val: []int{val}, Flag: input.ValueInt}
	case int64:
		return &Int64{Val: []int64{val}, Flag: input.ValueInt64}
	case uint:
		return &Uint{Val: []uint{val}, Flag: input.ValueUint}
	case uint64:
		return &Uint64{Val: []uint64{val}, Flag: input.ValueUint64}
	case float64:
		return &Float64{Val: []float64{val}, Flag: input.ValueFloat64}
	case bool:
		return &Bool{Val: []bool{val}, Flag: input.ValueBool}
	case time.Duration:
		return &Duration{Val: []time.Duration{val}, Flag: input.ValueDuration}
	case time.Time:
		return &Time{Val: []time.Time{val}, Flag: input.ValueTime}
	case []int64:
		return &Int64{Val: val, Flag: input.ValueInt64 | input.ValueArray}
	case []uint:
		return &Uint{Val: val, Flag: input.ValueUint | input.ValueArray}
	case []uint64:
		return &Uint64{Val: val, Flag: input.ValueUint64 | input.ValueArray}
	case []float64:
		return &Float64{Val: val, Flag: input.ValueFloat64 | input.ValueArray}
	case []bool:
		return &Bool{Val: val, Flag: input.ValueBool | input.ValueArray}
	case []time.Duration:
		return &Duration{Val: val, Flag: input.ValueDuration | input.ValueArray}
	case []time.Time:
		return &Time{Val: val, Flag: input.ValueTime | input.ValueArray}
	case []string:
		return &String{Val: val, Flag: input.ValueString | input.ValueArray}
	case []int:
		return &Int{Val: val, Flag: input.ValueInt | input.ValueArray}
	case []interface{}:
		return &Any{Val: val, Flag: input.ValueAny | input.ValueArray}
	case input.Value:
		return val
	default:
		if v != nil {
			return &Any{Val: []interface{}{v}, Flag: input.ValueAny}
		}

		return &Empty{}
	}
}

func ByFlag(flag input.Flag) input.AppendValue {
	switch {
	case flag.IsInt():
		return &Int{Flag: flag | input.ValueInt}
	case flag.IsInt64():
		return &Int64{Flag: flag | input.ValueInt64}
	case flag.IsUint():
		return &Uint{Flag: flag | input.ValueUint}
	case flag.IsUint64():
		return &Uint64{Flag: flag | input.ValueUint64}
	case flag.IsFloat64():
		return &Float64{Flag: flag | input.ValueFloat64}
	case flag.IsBool():
		return &Bool{Flag: flag | input.ValueBool}
	case flag.IsDuration():
		return &Duration{Flag: flag | input.ValueDuration}
	case flag.IsTime():
		return &Time{Flag: flag | input.ValueTime}
	case flag.IsAny():
		return &Any{Flag: flag | input.ValueAny}
	default:
		return &String{}
	}
}
