package setting

import (
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/errs"
)

type key uint8

const (
	paramHidden key = iota + 1
	paramDescription
	paramVerssion
	paramHelp
	paramUsage
)

const (
	defaultVersion = "undefined"
)

func IsHidden(in Setting) bool {
	data, ok := Bool(in, paramHidden)

	return ok && data
}

func Hidden(in Setting) Setting {
	return in.With(paramHidden, true)
}

func Description(in Setting) string {
	data, _ := String(in, paramDescription)

	return data
}

func WithDescription(desc string) Option {
	return func(p Setting) Setting {
		return p.With(paramDescription, desc)
	}
}

func Version(in Setting) string {
	if data, ok := String(in, paramVerssion); ok {
		return data
	}

	return defaultVersion
}

func WithVersion(in string) Option {
	return func(p Setting) Setting {
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
	return func(p Setting) Setting {
		return p.With(paramHelp, fn)
	}
}

func Help(in Setting, data HData) (string, error) {
	fn, ok := in.Param(paramHelp)
	if !ok {
		return "", nil
	}

	hfn, fok := fn.(HelpFn)
	if !fok {
		return "", fmt.Errorf("%w: expect:func(data HData) (string, error), got:%T", errs.ErrWrongType, fn)
	}

	return hfn(data)
}

func UsageData(name string, opts config.Options) UData {
	return UData{
		Options: opts,
		Name:    name,
	}
}

type UData struct {
	config.Options

	Name string
}

type UsageFn func(data UData) (string, error)

func WithUsage(fn UsageFn) Option {
	return func(p Setting) Setting {
		return p.With(paramUsage, fn)
	}
}

func Usage(in Setting, data UData) (string, error) {
	fn, ok := in.Param(paramUsage)
	if !ok {
		return "", fmt.Errorf("%w", errs.ErrNotFound)
	}

	ufn, ok := fn.(UsageFn)
	if !ok {
		return "", fmt.Errorf("%w: expect: func(data Udata) (string, error), got:%T", errs.ErrWrongType, fn)
	}

	return ufn(data)
}
