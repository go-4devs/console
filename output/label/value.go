package label

import "fmt"

type Type int

const (
	TypeAny Type = iota
	TypeBool
	TypeInt
	TypeInt64
	TypeUint
	TypeUint64
	TypeFloat64
	TypeString
)

type Value struct {
	vtype Type
	value any
}

func (v Value) String() string {
	return fmt.Sprint(v.value)
}

func AnyValue(v any) Value {
	return Value{vtype: TypeAny, value: v}
}

func BoolValue(v bool) Value {
	return Value{vtype: TypeBool, value: v}
}

func IntValue(v int) Value {
	return Value{vtype: TypeInt, value: v}
}

func Int64Value(v int64) Value {
	return Value{vtype: TypeInt64, value: v}
}

func UintValue(v uint) Value {
	return Value{vtype: TypeUint, value: v}
}

func Uint64Value(v uint64) Value {
	return Value{vtype: TypeUint64, value: v}
}

func Float64Value(v float64) Value {
	return Value{vtype: TypeFloat64, value: v}
}

func StringValue(v string) Value {
	return Value{vtype: TypeString, value: v}
}
