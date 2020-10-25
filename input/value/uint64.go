package value

import (
	"strconv"

	"gitoa.ru/go-4devs/console/input/flag"
)

type Uint64 struct {
	empty
	Val  []uint64
	Flag flag.Flag
}

func (u *Uint64) Append(in string) error {
	v, err := strconv.ParseUint(in, 10, 64)
	if err != nil {
		return err
	}

	u.Val = append(u.Val, v)

	return nil
}

func (u *Uint64) Uint64() uint64 {
	if !u.Flag.IsArray() && len(u.Val) == 1 {
		return u.Val[0]
	}

	return 0
}

func (u *Uint64) Uint64s() []uint64 {
	return u.Val
}

func (u *Uint64) Any() interface{} {
	if u.Flag.IsArray() {
		return u.Uint64s()
	}

	return u.Uint64()
}
