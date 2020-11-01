package argument

import (
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/value/flag"
)

func Required(a *Argument) {
	a.Flag |= flag.Required
}

func Default(v interface{}) func(*Argument) {
	return func(a *Argument) {
		a.Default = value.New(v)
	}
}

func Flag(flag flag.Flag) func(*Argument) {
	return func(a *Argument) {
		a.Flag = flag
	}
}

func Array(a *Argument) {
	a.Flag |= flag.Array
}
