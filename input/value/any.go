package value

import "gitoa.ru/go-4devs/console/input"

type Any struct {
	Empty
	Val  []interface{}
	Flag input.Flag
}

func (a *Any) Any() interface{} {
	if a.Flag.IsArray() {
		return a.Val
	}

	if len(a.Val) > 0 {
		return a.Val[0]
	}

	return nil
}
