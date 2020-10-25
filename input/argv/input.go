package argv

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/wrap"
)

const doubleDash = `--`

var _ input.ReadInput = (*Input)(nil)

func WithErrorHandle(h func(error) error) func(*Input) {
	return func(i *Input) {
		i.errorHandle = h
	}
}

func New(args []string, opts ...func(*Input)) *wrap.Input {
	i := &Input{
		args:      args,
		arguments: make(map[string]input.AppendValue),
		options:   make(map[string]input.AppendValue),
		errorHandle: func(err error) error {
			return err
		},
	}

	for _, opt := range opts {
		opt(i)
	}

	return &wrap.Input{ReadInput: i}
}

type Input struct {
	args        []string
	arguments   map[string]input.AppendValue
	options     map[string]input.AppendValue
	mu          sync.RWMutex
	errorHandle func(error) error
}

func (i *Input) ReadOption(ctx context.Context, name string) (input.Value, error) {
	if v, ok := i.options[name]; ok {
		return v, nil
	}

	return nil, input.ErrNotFound
}

func (i *Input) SetOption(name string, val input.Value) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.options[name] = &value.Read{Value: val}
}

func (i *Input) ReadArgument(ctx context.Context, name string) (input.Value, error) {
	if v, ok := i.arguments[name]; ok {
		return v, nil
	}

	return nil, input.ErrNotFound
}

func (i *Input) SetArgument(name string, val input.Value) {
	i.mu.Lock()
	defer i.mu.Unlock()

	i.arguments[name] = &value.Read{Value: val}
}

func (i *Input) Bind(ctx context.Context, def *input.Definition) error {
	options := true

	for len(i.args) > 0 {
		var err error

		arg := i.args[0]
		i.args = i.args[1:]

		switch {
		case options && arg == doubleDash:
			options = false
		case options && len(arg) > 2 && arg[0:2] == doubleDash:
			err = i.parseLongOption(arg[2:], def)
		case options && arg[0:1] == "-":
			if len(arg) == 1 {
				return fmt.Errorf("%w: option name required given '-'", input.ErrInvalidName)
			}

			err = i.parseShortOption(arg[1:], def)
		default:
			err = i.parseArgument(arg, def)
		}

		if err != nil {
			if herr := i.errorHandle(err); herr != nil {
				return herr
			}
		}
	}

	return nil
}

func (i *Input) parseLongOption(arg string, def *input.Definition) error {
	var value *string

	name := arg

	if strings.Contains(arg, "=") {
		vals := strings.SplitN(arg, "=", 2)
		name = vals[0]
		value = &vals[1]
	}

	opt, err := def.Option(name)
	if err != nil {
		return input.ErrorOption(name, err)
	}

	return i.appendOption(name, value, opt)
}

func (i *Input) appendOption(name string, data *string, opt input.Option) error {
	v, ok := i.options[name]

	if ok && !opt.IsArray() {
		return fmt.Errorf("%w: got: array, expect: %s", input.ErrUnexpectedType, input.Type(opt.Flag))
	}

	var val string

	switch {
	case data != nil:
		val = *data
	case opt.IsBool():
		val = "true"
	case len(i.args) > 0 && len(i.args[0]) > 0 && i.args[0][0:1] != "-":
		val = i.args[0]
		i.args = i.args[1:]
	default:
		return input.ErrorOption(name, input.ErrRequired)
	}

	if !ok {
		v = value.ByFlag(opt.Flag)
		i.options[name] = v
	}

	if err := v.Append(val); err != nil {
		return input.ErrorOption(name, err)
	}

	return nil
}

func (i *Input) parseShortOption(arg string, def *input.Definition) error {
	name := arg

	var value string

	if len(name) > 1 {
		name, value = arg[0:1], arg[1:]
	}

	opt, err := def.ShortOption(name)
	if err != nil {
		return err
	}

	if opt.IsBool() && value != "" {
		if err := i.parseShortOption(value, def); err != nil {
			return err
		}

		value = ""
	}

	if value == "" {
		return i.appendOption(opt.Name, nil, opt)
	}

	return i.appendOption(opt.Name, &value, opt)
}

func (i *Input) parseArgument(arg string, def *input.Definition) error {
	opt, err := def.Argument(len(i.arguments))
	if err != nil {
		return err
	}

	v, ok := i.arguments[opt.Name]
	if !ok {
		v = value.ByFlag(opt.Flag)
		i.arguments[opt.Name] = v
	}

	if err := v.Append(arg); err != nil {
		return input.ErrorArgument(opt.Name, err)
	}

	return nil
}
