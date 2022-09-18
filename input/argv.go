package input

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/input/errs"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/input/variable"
)

const doubleDash = `--`

func NewArgs(skip int) *Argv {
	res := 2

	switch {
	case skip > 0 && len(os.Args) > skip:
		res = skip
	case skip > 0:
		res = len(os.Args)
	case len(os.Args) == 1:
		res = 1
	case len(os.Args) > 1 && os.Args[1][0] == '-':
		res = 1
	}

	return &Argv{Args: os.Args[res:]}
}

type Argv struct {
	Array
	Args      []string
	ErrHandle func(error) error
}

//nolint:cyclop
func (i *Argv) Bind(ctx context.Context, def *Definition) error {
	options := true

	for len(i.Args) > 0 {
		var err error

		arg := i.Args[0]
		i.Args = i.Args[1:]

		switch {
		case options && arg == doubleDash:
			options = false
		case options && len(arg) > 2 && arg[0:2] == doubleDash:
			err = i.parseLongOption(arg[2:], def)
		case options && arg[0:1] == "-":
			if len(arg) == 1 {
				return fmt.Errorf("%w: option name required given '-'", errs.ErrInvalidName)
			}

			err = i.parseShortOption(arg[1:], def)
		default:
			err = i.parseArgument(arg, def)
		}

		if err != nil && i.ErrHandle != nil {
			if herr := i.ErrHandle(err); herr != nil {
				return herr
			}
		}
	}

	return i.Array.Bind(ctx, def)
}

func (i *Argv) parseLongOption(arg string, def *Definition) error {
	var value *string

	name := arg

	if strings.Contains(arg, "=") {
		vals := strings.SplitN(arg, "=", 2)
		name = vals[0]
		value = &vals[1]
	}

	opt, err := def.Option(name)
	if err != nil {
		return option.Err(name, err)
	}

	return i.appendOption(name, value, opt)
}

func (i *Argv) appendOption(name string, data *string, opt variable.Variable) error {
	if i.HasOption(name) && !opt.IsArray() {
		return fmt.Errorf("%w: got: array, expect: %s", errs.ErrUnexpectedType, opt.Flag.Type())
	}

	var val string

	switch {
	case data != nil:
		val = *data
	case opt.IsBool():
		val = "true"
	case len(i.Args) > 0 && len(i.Args[0]) > 0 && i.Args[0][0:1] != "-":
		val = i.Args[0]
		i.Args = i.Args[1:]
	default:
		return option.Err(name, errs.ErrRequired)
	}

	if err := i.AppendOption(opt, val); err != nil {
		return option.Err(name, err)
	}

	return nil
}

func (i *Argv) parseShortOption(arg string, def *Definition) error {
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

func (i *Argv) parseArgument(arg string, def *Definition) error {
	opt, err := def.Argument(i.LenArguments())
	if err != nil {
		return err
	}

	if err := i.AppendArgument(opt, arg); err != nil {
		return argument.Err(opt.Name, err)
	}

	return nil
}
