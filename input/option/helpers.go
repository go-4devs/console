package option

import "gitoa.ru/go-4devs/console/input"

func Bool(name, description string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, description, append(opts, Value(input.ValueBool))...)
}

func Duration(name, description string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, description, append(opts, Value(input.ValueDuration))...)
}

func Float64(name, description string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, description, append(opts, Value(input.ValueFloat64))...)
}

func Int(name, description string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, description, append(opts, Value(input.ValueInt))...)
}

func Int64(name, description string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, description, append(opts, Value(input.ValueInt64))...)
}

func Time(name, description string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, description, append(opts, Value(input.ValueTime))...)
}

func Uint(name, description string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, description, append(opts, Value(input.ValueUint))...)
}

func Uint64(name, descriontion string, opts ...func(*input.Option)) input.Option {
	return input.NewOption(name, descriontion, append(opts, Value(input.ValueUint64))...)
}
