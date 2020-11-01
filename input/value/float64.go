package value

import (
	"strconv"

	"gitoa.ru/go-4devs/console/input/value/flag"
)

type Float64 struct {
	empty
	Val  []float64
	Flag flag.Flag
}

func (f *Float64) Append(in string) error {
	v, err := strconv.ParseFloat(in, 64)
	if err != nil {
		return err
	}

	f.Val = append(f.Val, v)

	return nil
}

func (f *Float64) Float64() float64 {
	if !f.Flag.IsArray() && len(f.Val) == 1 {
		return f.Val[0]
	}

	return 0
}

func (f *Float64) Float64s() []float64 {
	return f.Val
}

func (f *Float64) Any() interface{} {
	if f.Flag.IsArray() {
		return f.Float64s()
	}

	return f.Float64()
}
