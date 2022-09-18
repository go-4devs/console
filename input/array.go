package input

import (
	"context"

	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/input/errs"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/input/value"
)

type Array struct {
	Map
	defaults Map
}

func (a *Array) Option(ctx context.Context, name string) value.Value {
	if v := a.Map.Option(ctx, name); !value.IsEmpty(v) {
		return v
	}

	if v := a.defaults.Option(ctx, name); !value.IsEmpty(v) {
		return v
	}

	return value.Empty()
}

func (a *Array) Argument(ctx context.Context, name string) value.Value {
	if v := a.Map.Argument(ctx, name); !value.IsEmpty(v) {
		return v
	}

	if v := a.defaults.Argument(ctx, name); !value.IsEmpty(v) {
		return v
	}

	return value.Empty()
}

func (a *Array) Bind(ctx context.Context, d *Definition) error {
	if err := a.bindArguments(ctx, d); err != nil {
		return err
	}

	return a.bindOption(ctx, d)
}

func (a *Array) bindOption(ctx context.Context, def *Definition) error {
	for _, name := range def.Options() {
		opt, err := def.Option(name)
		if err != nil {
			return err
		}

		if !a.HasOption(name) {
			switch {
			case opt.HasDefault():
				a.defaults.SetOption(name, opt.Default)

				continue
			case opt.IsRequired():
				return option.Err(name, errs.ErrRequired)
			default:
				continue
			}
		}

		v := a.Map.Option(ctx, name)
		if value.IsEmpty(v) {
			continue
		}

		if err := opt.Validate(v); err != nil {
			return option.Err(name, err)
		}
	}

	return nil
}

func (a *Array) bindArguments(ctx context.Context, def *Definition) error {
	for pos, name := range def.Arguments() {
		arg, err := def.Argument(pos)
		if err != nil {
			return err
		}

		if !a.HasArgument(name) {
			switch {
			case arg.HasDefault():
				a.defaults.SetArgument(name, arg.Default)

				continue
			case arg.IsRequired():
				return argument.Err(name, errs.ErrRequired)
			default:
				continue
			}
		}

		if v := a.Map.Argument(ctx, name); !value.IsEmpty(v) {
			if err := arg.Validate(v); err != nil {
				return argument.Err(name, err)
			}
		}
	}

	return nil
}
