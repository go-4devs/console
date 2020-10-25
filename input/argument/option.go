package argument

import (
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
)

func Required(a *input.Argument) {
	a.Flag |= input.ValueRequired
}

func Default(v interface{}) func(*input.Argument) {
	return func(a *input.Argument) {
		a.Default = value.New(v)
	}
}

func Flag(flag input.Flag) func(*input.Argument) {
	return func(a *input.Argument) {
		a.Flag = flag
	}
}

func Array(a *input.Argument) {
	a.Flag |= input.ValueArray
}
