package wrap

import (
	"context"
	"errors"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
)

type Input struct {
	input.ReadInput
}

func (i *Input) Option(ctx context.Context, name string) value.Value {
	if v, err := i.ReadOption(ctx, name); err == nil {
		return v
	}

	return &value.Empty{}
}

func (i *Input) Argument(ctx context.Context, name string) value.Value {
	if v, err := i.ReadArgument(ctx, name); err == nil {
		return v
	}

	return &value.Empty{}
}

func (i *Input) Bind(ctx context.Context, def *input.Definition) error {
	if err := i.ReadInput.Bind(ctx, def); err != nil {
		return err
	}

	if err := i.bindArguments(ctx, def); err != nil {
		return err
	}

	return i.bindOptions(ctx, def)
}

func (i *Input) bindOptions(ctx context.Context, def *input.Definition) error {
	for _, name := range def.Options() {
		opt, err := def.Option(name)
		if err != nil {
			return err
		}

		v, err := i.ReadOption(ctx, name)
		if err != nil && !errors.Is(err, input.ErrNotFound) {
			return input.ErrorOption(name, err)
		}

		if err == nil {
			if err := opt.Validate(v); err != nil {
				return input.ErrorOption(name, err)
			}

			continue
		}

		if opt.IsRequired() && !opt.HasDefault() {
			return input.ErrorOption(name, input.ErrRequired)
		}

		if opt.HasDefault() {
			i.SetOption(name, opt.Default)
		}
	}

	return nil
}

func (i *Input) bindArguments(ctx context.Context, def *input.Definition) error {
	for pos, name := range def.Arguments() {
		arg, err := def.Argument(pos)
		if err != nil {
			return err
		}

		v, err := i.ReadArgument(ctx, name)
		if err != nil && !errors.Is(err, input.ErrNotFound) {
			return input.ErrorArgument(name, err)
		}

		if err == nil {
			if err := arg.Validate(v); err != nil {
				return input.ErrorArgument(name, err)
			}

			continue
		}

		if arg.IsRequired() && !arg.HasDefault() {
			return input.ErrorArgument(name, input.ErrRequired)
		}

		if arg.HasDefault() {
			i.SetArgument(name, arg.Default)
		}
	}

	return nil
}
