package value

import (
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
)

type Int64 struct {
	Empty
	Val  []int64
	Flag flag.Flag
}

func (i *Int64) Int64() int64 {
	if !i.Flag.IsArray() && len(i.Val) == 1 {
		return i.Val[0]
	}

	return 0
}

func (i *Int64) Int64s() []int64 {
	return i.Val
}

func (i *Int64) Any() interface{} {
	if i.Flag.IsArray() {
		return i.Int64s()
	}

	return i.Int64()
}

func (i *Int64) Append(in string) error {
	v, err := strconv.ParseInt(in, 10, 64)
	if err != nil {
		return err
	}

	i.Val = append(i.Val, v)

	return nil
}
