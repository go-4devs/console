package value

import (
	"time"

	"gitoa.ru/go-4devs/console/input/value/flag"
)

type Time struct {
	empty
	Val  []time.Time
	Flag flag.Flag
}

func (t *Time) Append(in string) error {
	v, err := time.Parse(time.RFC3339, in)
	if err != nil {
		return err
	}

	t.Val = append(t.Val, v)

	return nil
}

func (t *Time) Time() time.Time {
	if !t.Flag.IsArray() && len(t.Val) == 1 {
		return t.Val[0]
	}

	return time.Time{}
}

func (t *Time) Times() []time.Time {
	return t.Val
}

func (t *Time) Amy() interface{} {
	if t.Flag.IsArray() {
		return t.Times()
	}

	return t.Time()
}
