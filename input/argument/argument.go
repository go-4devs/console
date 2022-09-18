package argument

import (
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/variable"
)

func Default(in interface{}) variable.Option {
	return variable.Default(value.New(in))
}

func Required(v *variable.Variable) {
	variable.Required(v)
}

func Array(v *variable.Variable) {
	variable.Array(v)
}

func String(name, description string, opts ...variable.Option) variable.Variable {
	return variable.String(name, description, append(opts, variable.ArgArgument)...)
}

func Bool(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Bool(name, description, append(opts, variable.ArgArgument)...)
}

func Duration(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Duration(name, description, append(opts, variable.ArgArgument)...)
}

func Float64(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Float64(name, description, append(opts, variable.ArgArgument)...)
}

func Int(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Int(name, description, append(opts, variable.ArgArgument)...)
}

func Int64(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Int64(name, description, append(opts, variable.ArgArgument)...)
}

func Time(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Time(name, description, append(opts, variable.ArgArgument)...)
}

func Uint(name, description string, opts ...variable.Option) variable.Variable {
	return variable.Uint(name, description, append(opts, variable.ArgArgument)...)
}

func Uint64(name, descriontion string, opts ...variable.Option) variable.Variable {
	return variable.Uint64(name, descriontion, append(opts, variable.ArgArgument)...)
}

func Err(name string, err error) variable.Error {
	return variable.Err(name, variable.TypeArgument, err)
}
