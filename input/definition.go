package input

import (
	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/input/errs"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/input/variable"
)

func NewDefinition() *Definition {
	return &Definition{
		options: make(map[string]variable.Variable),
		args:    make(map[string]variable.Variable),
		short:   make(map[string]string),
	}
}

type Definition struct {
	options map[string]variable.Variable
	posOpt  []string
	args    map[string]variable.Variable
	posArgs []string
	short   map[string]string
}

func (d *Definition) Options() []string {
	return d.posOpt
}

func (d *Definition) Arguments() []string {
	return d.posArgs
}

func (d *Definition) SetOption(name, description string, opts ...variable.Option) *Definition {
	return d.SetOptions(option.String(name, description, opts...))
}

func (d *Definition) SetOptions(opts ...variable.Variable) *Definition {
	for _, opt := range opts {
		if _, has := d.options[opt.Name]; !has {
			d.posOpt = append([]string{opt.Name}, d.posOpt...)
		}

		d.options[opt.Name] = opt
		if opt.HasShort() {
			d.short[opt.Alias] = opt.Name
		}
	}

	return d
}

func (d *Definition) SetArgument(name, description string, opts ...variable.Option) *Definition {
	return d.SetArguments(argument.String(name, description, opts...))
}

func (d *Definition) SetArguments(args ...variable.Variable) *Definition {
	for _, arg := range args {
		if _, ok := d.args[arg.Name]; !ok {
			d.posArgs = append(d.posArgs, arg.Name)
		}

		d.args[arg.Name] = arg
	}

	return d
}

func (d *Definition) Argument(pos int) (variable.Variable, error) {
	if len(d.posArgs) == 0 {
		return variable.Variable{}, errs.ErrNoArgs
	}

	lastPos := len(d.posArgs) - 1
	if lastPos < pos {
		arg := d.args[d.posArgs[lastPos]]
		if arg.IsArray() {
			return arg, nil
		}

		return variable.Variable{}, errs.ErrToManyArgs
	}

	return d.args[d.posArgs[pos]], nil
}

func (d *Definition) ShortOption(short string) (variable.Variable, error) {
	name, ok := d.short[short]
	if !ok {
		return variable.Variable{}, errs.ErrNotFound
	}

	return d.Option(name)
}

func (d *Definition) Option(name string) (variable.Variable, error) {
	if opt, ok := d.options[name]; ok {
		return opt, nil
	}

	return variable.Variable{}, errs.ErrNotFound
}
