package value

import (
	"fmt"
	"strconv"
	"time"
)

var (
	_ ParseValue = (String)("")
	_ SliceValue = (Strings)(nil)
)

func NewStrings(in []string) Slice {
	return Slice{SliceValue: Strings(in)}
}

type Strings []string

func (s Strings) Unmarshal(in interface{}) error {
	val, ok := in.(*[]string)
	if !ok {
		return fmt.Errorf("%w: expect *[]string", ErrWrongType)
	}

	*val = s

	return nil
}

func (s Strings) Any() interface{} {
	return s.Strings()
}

func (s Strings) Strings() []string {
	out := make([]string, len(s))
	copy(out, s)

	return out
}

func (s Strings) Ints() []int {
	return nil
}

func (s Strings) Int64s() []int64 {
	return nil
}

func (s Strings) Uints() []uint {
	return nil
}

func (s Strings) Uint64s() []uint64 {
	return nil
}

func (s Strings) Float64s() []float64 {
	return nil
}

func (s Strings) Bools() []bool {
	return nil
}

func (s Strings) Durations() []time.Duration {
	return nil
}

func (s Strings) Times() []time.Time {
	return nil
}

func NewString(in string) Value {
	return Read{ParseValue: String(in)}
}

type String string

func (s String) ParseString() (string, error) {
	return string(s), nil
}

func (s String) Unmarshal(in interface{}) error {
	v, ok := in.(*string)
	if !ok {
		return fmt.Errorf("%w: expect *string", ErrWrongType)
	}

	*v = string(s)

	return nil
}

func (s String) Any() interface{} {
	return string(s)
}

func (s String) ParseInt() (int, error) {
	v, err := strconv.Atoi(string(s))
	if err != nil {
		return 0, fmt.Errorf("string int:%w", err)
	}

	return v, nil
}

func (s String) Int64() int64 {
	out, _ := s.ParseInt64()

	return out
}

func (s String) ParseInt64() (int64, error) {
	v, err := strconv.ParseInt(string(s), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("string int64:%w", err)
	}

	return v, nil
}

func (s String) ParseUint() (uint, error) {
	uout, err := s.ParseUint64()

	return uint(uout), err
}

func (s String) ParseUint64() (uint64, error) {
	uout, err := strconv.ParseUint(string(s), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("string uint:%w", err)
	}

	return uout, nil
}

func (s String) ParseFloat64() (float64, error) {
	fout, err := strconv.ParseFloat(string(s), 64)
	if err != nil {
		return 0, fmt.Errorf("string float64:%w", err)
	}

	return fout, nil
}

func (s String) ParseBool() (bool, error) {
	v, err := strconv.ParseBool(string(s))
	if err != nil {
		return false, fmt.Errorf("string bool:%w", err)
	}

	return v, nil
}

func (s String) ParseDuration() (time.Duration, error) {
	v, err := time.ParseDuration(string(s))
	if err != nil {
		return 0, fmt.Errorf("string duration:%w", err)
	}

	return v, nil
}

func (s String) ParseTime() (time.Time, error) {
	v, err := time.Parse(time.RFC3339, string(s))
	if err != nil {
		return time.Time{}, fmt.Errorf("string time:%w", err)
	}

	return v, nil
}
