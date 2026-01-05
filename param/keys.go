package param

import (
	"fmt"

	cerr "gitoa.ru/go-4devs/console/errors"
)

type key uint8

const (
	paramHidden key = iota + 1
	paramDescription
	paramVerssion
	paramHelp
)

const (
	defaultVersion = "undefined"
)

func IsHidden(in Params) bool {
	data, ok := Bool(in, paramHidden)

	return ok && data
}

func Hidden(in Params) Params {
	return in.With(paramHidden, true)
}

func Description(in Params) string {
	data, _ := String(in, paramDescription)

	return data
}

func WithDescription(desc string) Option {
	return func(p Params) Params {
		return p.With(paramDescription, desc)
	}
}

func Version(in Params) string {
	if data, ok := String(in, paramVerssion); ok {
		return data
	}

	return defaultVersion
}

func WithVersion(in string) Option {
	return func(p Params) Params {
		return p.With(paramVerssion, in)
	}
}

func HelpData(bin, name string) HData {
	return HData{
		Bin:  bin,
		Name: name,
	}
}

type HData struct {
	Bin  string
	Name string
}

type HelpFn func(data HData) (string, error)

func WithHelp(fn HelpFn) Option {
	return func(p Params) Params {
		return p.With(paramHelp, fn)
	}
}

func Help(in Params, data HData) (string, error) {
	fn, ok := in.Param(paramHelp)
	if !ok {
		return "", nil
	}

	hfn, fok := fn.(HelpFn)
	if !fok {
		return "", fmt.Errorf("%w: expect:%T, got:%T", cerr.ErrWrongType, (HelpFn)(nil), fn)
	}

	return hfn(data)
}
