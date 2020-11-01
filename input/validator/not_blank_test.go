package validator_test

import (
	"errors"
	"testing"
	"time"

	"gitoa.ru/go-4devs/console/input/validator"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/value/flag"
)

func TestNotBlank(t *testing.T) {
	cases := map[string]struct {
		flag  flag.Flag
		value value.Value
		empty value.Value
	}{
		"any": {flag: flag.Any, value: value.New(float32(1))},
		"array int": {
			flag:  flag.Int | flag.Array,
			value: value.New([]int{1}),
			empty: value.New([]int{10, 20, 0}),
		},
		"array int64": {
			flag:  flag.Int64 | flag.Array,
			value: value.New([]int64{1}),
			empty: value.New([]int64{0}),
		},
		"array uint": {
			flag:  flag.Uint | flag.Array,
			value: value.New([]uint{1}),
			empty: value.New([]uint{1, 0}),
		},
		"array uint64": {
			flag:  flag.Uint64 | flag.Array,
			value: value.New([]uint64{1}),
			empty: value.New([]uint64{0}),
		},
		"array float64": {
			flag:  flag.Float64 | flag.Array,
			value: value.New([]float64{0.2}),
			empty: value.New([]float64{0}),
		},
		"array bool": {
			flag:  flag.Bool | flag.Array,
			value: value.New([]bool{true, false}),
			empty: value.New([]bool{}),
		},
		"array duration": {
			flag:  flag.Duration | flag.Array,
			value: value.New([]time.Duration{time.Second}),
			empty: value.New([]time.Duration{time.Second, 0}),
		},
		"array time": {
			flag:  flag.Time | flag.Array,
			value: value.New([]time.Time{time.Now()}),
			empty: value.New([]time.Time{{}, time.Now()}),
		},
		"array string": {
			flag:  flag.Array,
			value: value.New([]string{"value"}),
			empty: value.New([]string{""}),
		},
		"int": {
			flag:  flag.Int,
			value: value.New(int(1)),
		},
		"int64": {
			flag:  flag.Int64,
			value: value.New(int64(2)),
		},
		"uint": {
			flag:  flag.Uint,
			value: value.New(uint(1)),
			empty: value.New([]uint{1}),
		},
		"uint64": {
			flag:  flag.Uint64,
			value: value.New(uint64(10)),
		},
		"float64": {
			flag:  flag.Float64,
			value: value.New(float64(.00001)),
		},
		"duration": {
			flag:  flag.Duration,
			value: value.New(time.Minute),
			empty: value.New("same string"),
		},
		"time":   {flag: flag.Time, value: value.New(time.Now())},
		"string": {value: value.New("string"), empty: value.New("")},
	}

	for name, ca := range cases {
		valid := validator.NotBlank(ca.flag)
		if err := valid(ca.value); err != nil {
			t.Errorf("case: %s, expected error <nil>, got: %s", name, err)
		}

		if ca.empty == nil {
			ca.empty = value.Empty()
		}

		if err := valid(ca.empty); err == nil || !errors.Is(err, validator.ErrNotBlank) {
			t.Errorf("case: %s, expect: %s, got:%s", name, validator.ErrNotBlank, err)
		}
	}
}
