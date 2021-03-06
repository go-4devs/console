package value

import (
	"strconv"

	"gitoa.ru/go-4devs/console/input/value/flag"
)

type Uint struct {
	empty
	Val  []uint
	Flag flag.Flag
}

func (u *Uint) Append(in string) error {
	v, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return err
	}

	u.Val = append(u.Val, uint(v))

	return nil
}

func (u *Uint) Uint() uint {
	if !u.Flag.IsArray() && len(u.Val) == 1 {
		return u.Val[0]
	}

	return 0
}

func (u *Uint) Uints() []uint {
	return u.Val
}

func (u *Uint) Any() interface{} {
	if u.Flag.IsArray() {
		return u.Uints()
	}

	return u.Uint()
}
