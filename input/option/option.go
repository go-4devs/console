package option

import (
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/variable"
)

func Short(in rune) variable.Option {
	return func(v *variable.Variable) {
		v.Alias = string(in)
	}
}

func Default(in interface{}) variable.Option {
	return variable.Default(value.New(in))
}

func Hidden(in *variable.Variable) {
	variable.Hidden(in)
}

func Required(v *variable.Variable) {
	variable.Required(v)
}

func Valid(f ...func(value.Value) error) variable.Option {
	return variable.Valid(f...)
}

func Array(v *variable.Variable) {
	variable.Array(v)
}

func String(name, description string, opts ...variable.Option) variable.Variable {
	return variable.String(name, description, append(opts, variable.ArgOption)...)
}

func Bool(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Bool(name, description, append(opts, variable.ArgOption)...)
}

func Duration(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Duration(name, description, append(opts, variable.ArgOption)...)
}

func Float64(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Float64(name, description, append(opts, variable.ArgOption)...)
}

func Int(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Int(name, description, append(opts, variable.ArgOption)...)
}

func Int64(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Int64(name, description, append(opts, variable.ArgOption)...)
}

func Time(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Time(name, description, append(opts, variable.ArgOption)...)
}

func Uint(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Uint(name, description, append(opts, variable.ArgOption)...)
}

func Uint64(name, descriontion string, opts ...variable.Option) variable.Variable {
	return variable.Uint64(name, descriontion, append(opts, variable.ArgOption)...)
}

func Err(name string, err error) variable.Error {
	return variable.Err(name, variable.TypeOption, err)
}
