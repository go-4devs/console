package option

import (
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
)

func Required(o *input.Option) {
	o.Flag |= input.ValueRequired
}

func Default(in interface{}) func(*input.Option) {
	return func(o *input.Option) {
		o.Default = value.New(in)
	}
}

func Short(s string) func(*input.Option) {
	return func(o *input.Option) {
		o.Short = s
	}
}

func Array(o *input.Option) {
	o.Flag |= input.ValueArray
}

func Value(flag input.Flag) func(*input.Option) {
	return func(o *input.Option) {
		o.Flag |= flag
	}
}

func Flag(in input.Flag) func(*input.Option) {
	return func(o *input.Option) {
		o.Flag = in
	}
}

func Valid(f ...func(input.Value) error) func(*input.Option) {
	return func(o *input.Option) {
		o.Valid = f
	}
}
