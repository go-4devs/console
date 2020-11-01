package value

import (
	"strconv"

	"gitoa.ru/go-4devs/console/input/value/flag"
)

type Bool struct {
	empty
	Val  []bool
	Flag flag.Flag
}

func (b *Bool) Append(in string) error {
	v, err := strconv.ParseBool(in)
	if err != nil {
		return err
	}

	b.Val = append(b.Val, v)

	return nil
}

func (b *Bool) Bool() bool {
	if !b.Flag.IsArray() && len(b.Val) == 1 {
		return b.Val[0]
	}

	return false
}

func (b *Bool) Bools() []bool {
	return b.Val
}

func (b *Bool) Any() interface{} {
	if b.Flag.IsArray() {
		return b.Bools()
	}

	return b.Bool()
}
