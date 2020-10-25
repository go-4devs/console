package input

import (
	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/input/option"
)

func NewDefinition() *Definition {
	return &Definition{
		options: make(map[string]option.Option),
		args:    make(map[string]argument.Argument),
		short:   make(map[string]string),
	}
}

type Definition struct {
	options map[string]option.Option
	posOpt  []string
	args    map[string]argument.Argument
	posArgs []string
	short   map[string]string
}

func (d *Definition) Options() []string {
	return d.posOpt
}

func (d *Definition) Arguments() []string {
	return d.posArgs
}

func (d *Definition) SetOption(name, description string, opts ...func(*option.Option)) *Definition {
	return d.SetOptions(option.New(name, description, opts...))
}

func (d *Definition) SetOptions(opts ...option.Option) *Definition {
	for _, opt := range opts {
		if _, has := d.options[opt.Name]; !has {
			d.posOpt = append([]string{opt.Name}, d.posOpt...)
		}

		d.options[opt.Name] = opt
		if opt.HasShort() {
			d.short[opt.Short] = opt.Name
		}
	}

	return d
}

func (d *Definition) SetArgument(name, description string, opts ...func(*argument.Argument)) *Definition {
	return d.SetArguments(argument.New(name, description, opts...))
}

func (d *Definition) SetArguments(args ...argument.Argument) *Definition {
	for _, arg := range args {
		if _, ok := d.args[arg.Name]; !ok {
			d.posArgs = append(d.posArgs, arg.Name)
		}

		d.args[arg.Name] = arg
	}

	return d
}

func (d *Definition) Argument(pos int) (argument.Argument, error) {
	if len(d.posArgs) == 0 {
		return argument.Argument{}, ErrNoArgs
	}

	lastPos := len(d.posArgs) - 1
	if lastPos < pos {
		arg := d.args[d.posArgs[lastPos]]
		if arg.IsArray() {
			return arg, nil
		}

		return argument.Argument{}, ErrToManyArgs
	}

	return d.args[d.posArgs[pos]], nil
}

func (d *Definition) ShortOption(short string) (option.Option, error) {
	name, ok := d.short[short]
	if !ok {
		return option.Option{}, ErrNotFound
	}

	return d.Option(name)
}

func (d *Definition) Option(name string) (option.Option, error) {
	if opt, ok := d.options[name]; ok {
		return opt, nil
	}

	return option.Option{}, ErrNotFound
}
