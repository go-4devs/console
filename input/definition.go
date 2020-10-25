package input

func NewDefinition() *Definition {
	return &Definition{
		options: make(map[string]Option),
		args:    make(map[string]Argument),
		short:   make(map[string]string),
	}
}

type Definition struct {
	options map[string]Option
	posOpt  []string
	args    map[string]Argument
	posArgs []string
	short   map[string]string
}

func (d *Definition) Options() []string {
	return d.posOpt
}

func (d *Definition) Arguments() []string {
	return d.posArgs
}

func (d *Definition) SetOption(name, description string, opts ...func(*Option)) *Definition {
	return d.SetOptions(NewOption(name, description, opts...))
}

func (d *Definition) SetOptions(opts ...Option) *Definition {
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

func (d *Definition) SetArgument(name, description string, opts ...func(*Argument)) *Definition {
	return d.SetArguments(NewArgument(name, description, opts...))
}

func (d *Definition) SetArguments(args ...Argument) *Definition {
	for _, arg := range args {
		if _, ok := d.args[arg.Name]; !ok {
			d.posArgs = append(d.posArgs, arg.Name)
		}

		d.args[arg.Name] = arg
	}

	return d
}

func (d *Definition) Argument(pos int) (Argument, error) {
	if len(d.posArgs) == 0 {
		return Argument{}, ErrNoArgs
	}

	lastPos := len(d.posArgs) - 1
	if lastPos < pos {
		arg := d.args[d.posArgs[lastPos]]
		if arg.IsArray() {
			return arg, nil
		}

		return Argument{}, ErrToManyArgs
	}

	return d.args[d.posArgs[pos]], nil
}

func (d *Definition) ShortOption(short string) (Option, error) {
	name, ok := d.short[short]
	if !ok {
		return Option{}, ErrNotFound
	}

	return d.Option(name)
}

func (d *Definition) Option(name string) (Option, error) {
	if opt, ok := d.options[name]; ok {
		return opt, nil
	}

	return Option{}, ErrNotFound
}
