package option

import (
	"gitoa.ru/go-4devs/console/input/value/flag"
)

func Bool(name, description string, opts ...func(*Option)) Option {
	return New(name, description, append(opts, Value(flag.Bool))...)
}

func Duration(name, description string, opts ...func(*Option)) Option {
	return New(name, description, append(opts, Value(flag.Duration))...)
}

func Float64(name, description string, opts ...func(*Option)) Option {
	return New(name, description, append(opts, Value(flag.Float64))...)
}

func Int(name, description string, opts ...func(*Option)) Option {
	return New(name, description, append(opts, Value(flag.Int))...)
}

func Int64(name, description string, opts ...func(*Option)) Option {
	return New(name, description, append(opts, Value(flag.Int64))...)
}

func Time(name, description string, opts ...func(*Option)) Option {
	return New(name, description, append(opts, Value(flag.Time))...)
}

func Uint(name, description string, opts ...func(*Option)) Option {
	return New(name, description, append(opts, Value(flag.Uint))...)
}

func Uint64(name, descriontion string, opts ...func(*Option)) Option {
	return New(name, descriontion, append(opts, Value(flag.Uint64))...)
}
