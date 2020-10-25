package validator_test

import (
	"errors"
	"testing"
	"time"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/validator"
)

func TestNotBlank(t *testing.T) {
	cases := map[string]struct {
		flag  input.Flag
		value input.Value
		empty input.Value
	}{
		"any": {flag: input.ValueAny, value: value.New(float32(1))},
		"array int": {
			flag:  input.ValueInt | input.ValueArray,
			value: value.New([]int{1}),
			empty: value.New([]int{10, 20, 0}),
		},
		"array int64": {
			flag:  input.ValueInt64 | input.ValueArray,
			value: value.New([]int64{1}),
			empty: value.New([]int64{0}),
		},
		"array uint": {
			flag:  input.ValueUint | input.ValueArray,
			value: value.New([]uint{1}),
			empty: value.New([]uint{1, 0}),
		},
		"array uint64": {
			flag:  input.ValueUint64 | input.ValueArray,
			value: value.New([]uint64{1}),
			empty: value.New([]uint64{0}),
		},
		"array float64": {
			flag:  input.ValueFloat64 | input.ValueArray,
			value: value.New([]float64{0.2}),
			empty: value.New([]float64{0}),
		},
		"array bool": {
			flag:  input.ValueBool | input.ValueArray,
			value: value.New([]bool{true, false}),
			empty: value.New([]bool{}),
		},
		"array duration": {
			flag:  input.ValueDuration | input.ValueArray,
			value: value.New([]time.Duration{time.Second}),
			empty: value.New([]time.Duration{time.Second, 0}),
		},
		"array time": {
			flag:  input.ValueTime | input.ValueArray,
			value: value.New([]time.Time{time.Now()}),
			empty: value.New([]time.Time{{}, time.Now()}),
		},
		"array string": {
			flag:  input.ValueArray,
			value: value.New([]string{"value"}),
			empty: value.New([]string{""}),
		},
		"int": {
			flag:  input.ValueInt,
			value: value.New(int(1)),
		},
		"int64": {
			flag:  input.ValueInt64,
			value: value.New(int64(2)),
		},
		"uint": {
			flag:  input.ValueUint,
			value: value.New(uint(1)),
			empty: value.New([]uint{1}),
		},
		"uint64": {
			flag:  input.ValueUint64,
			value: value.New(uint64(10)),
		},
		"float64": {
			flag:  input.ValueFloat64,
			value: value.New(float64(.00001)),
		},
		"duration": {
			flag:  input.ValueDuration,
			value: value.New(time.Minute),
			empty: value.New("same string"),
		},
		"time":   {flag: input.ValueTime, value: value.New(time.Now())},
		"string": {value: value.New("string"), empty: value.New("")},
	}

	for name, ca := range cases {
		valid := validator.NotBlank(ca.flag)
		if err := valid(ca.value); err != nil {
			t.Errorf("case: %s, expected error <nil>, got: %s", name, err)
		}

		if ca.empty == nil {
			ca.empty = &value.Empty{}
		}

		if err := valid(ca.empty); err == nil || !errors.Is(err, validator.ErrNotBlank) {
			t.Errorf("case: %s, expect: %s, got:%s", name, validator.ErrNotBlank, err)
		}
	}
}
