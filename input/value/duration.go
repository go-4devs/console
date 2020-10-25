package value

import (
	"time"

	"gitoa.ru/go-4devs/console/input"
)

type Duration struct {
	Empty
	Val  []time.Duration
	Flag input.Flag
}

func (d *Duration) Append(in string) error {
	v, err := time.ParseDuration(in)
	if err != nil {
		return err
	}

	d.Val = append(d.Val, v)

	return nil
}

func (d *Duration) Duration() time.Duration {
	if !d.Flag.IsArray() && len(d.Val) == 1 {
		return d.Val[0]
	}

	return 0
}

func (d *Duration) Durations() []time.Duration {
	return d.Val
}

func (d *Duration) Any() interface{} {
	if d.Flag&input.ValueArray > 0 {
		return d.Durations()
	}

	return d.Duration()
}
