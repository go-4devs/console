package value

import (
	"strconv"

	"gitoa.ru/go-4devs/console/input"
)

type Int struct {
	Empty
	Val  []int
	Flag input.Flag
}

func (i *Int) Append(in string) error {
	v, err := strconv.Atoi(in)
	if err != nil {
		return err
	}

	i.Val = append(i.Val, v)

	return nil
}

func (i *Int) Int() int {
	if !i.Flag.IsArray() && len(i.Val) == 1 {
		return i.Val[0]
	}

	return 0
}

func (i *Int) Ints() []int {
	return i.Val
}

func (i *Int) Any() interface{} {
	if i.Flag&input.ValueArray > 0 {
		return i.Ints()
	}

	return i.Int()
}
