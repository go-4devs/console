package value

import (
	"encoding/json"
	"fmt"
	"time"
)

var _ Value = NewAny(nil)

//nolint:gochecknoglobals
var (
	emptyValue = NewAny(nil)
)

func Empty() Value {
	return emptyValue
}

func IsEmpty(v Value) bool {
	return v == nil || v == emptyValue
}

func NewAny(in interface{}) Value {
	return Read{Any{v: in}}
}

type Any struct {
	v interface{}
}

func (a Any) Any() interface{} {
	return a.v
}

func (a Any) Unmarshal(val interface{}) error {
	out, err := a.ParseString()
	if err != nil {
		return fmt.Errorf("any parse string:%w", err)
	}

	uerr := json.Unmarshal([]byte(out), val)
	if uerr != nil {
		return fmt.Errorf("any unmarshal: %w", uerr)
	}

	return nil
}

func (a Any) ParseString() (string, error) {
	if a.v == nil {
		return "", nil
	}

	bout, err := json.Marshal(a.v)
	if err != nil {
		return "", fmt.Errorf("any string:%w", err)
	}

	return string(bout), err
}

func (a Any) ParseInt() (int, error) {
	out, ok := a.v.(int)
	if !ok {
		return 0, a.wrongType("int")
	}

	return out, nil
}

func (a Any) ParseInt64() (int64, error) {
	out, ok := a.v.(int64)
	if !ok {
		return 0, a.wrongType("int64")
	}

	return out, nil
}

func (a Any) ParseUint() (uint, error) {
	out, ok := a.v.(uint)
	if !ok {
		return 0, a.wrongType("uint")
	}

	return out, nil
}

func (a Any) ParseUint64() (uint64, error) {
	out, ok := a.v.(uint64)
	if !ok {
		return 0, a.wrongType("uint64")
	}

	return out, nil
}

func (a Any) ParseFloat64() (float64, error) {
	out, ok := a.v.(float64)
	if !ok {
		return 0, a.wrongType("float64")
	}

	return out, nil
}

func (a Any) ParseBool() (bool, error) {
	out, ok := a.v.(bool)
	if !ok {
		return false, a.wrongType("bool")
	}

	return out, nil
}

func (a Any) ParseDuration() (time.Duration, error) {
	out, ok := a.v.(time.Duration)
	if !ok {
		return 0, a.wrongType("time.Duration")
	}

	return out, nil
}

func (a Any) ParseTime() (time.Time, error) {
	out, ok := a.v.(time.Time)
	if !ok {
		return time.Time{}, a.wrongType("time.Time")
	}

	return out, nil
}

func (a Any) wrongType(ex String) error {
	return fmt.Errorf("%w any: got: %T expect: %s", ErrWrongType, a.v, ex)
}
