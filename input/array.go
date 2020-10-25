package input

import (
	"context"
	"sync"

	"gitoa.ru/go-4devs/console/input/errs"
	"gitoa.ru/go-4devs/console/input/value"
)

type Array struct {
	defaults array
	value    array
}

func (a *Array) GetOption(name string) (value.Append, bool) {
	return a.value.GetOption(name)
}

func (a *Array) SetOption(name string, v interface{}) {
	a.value.SetOption(name, v)
}

func (a *Array) LenArguments() int {
	return a.value.LenArguments()
}

func (a *Array) GetArgument(name string) (value.Append, bool) {
	return a.value.GetArgument(name)
}

func (a *Array) SetArgument(name string, v interface{}) {
	a.value.SetArgument(name, v)
}

func (a *Array) Option(_ context.Context, name string) value.Value {
	if v, ok := a.value.GetOption(name); ok {
		return v
	}

	if v, ok := a.defaults.GetOption(name); ok {
		return v
	}

	return value.Empty
}

func (a *Array) Argument(_ context.Context, name string) value.Value {
	if v, ok := a.value.GetArgument(name); ok {
		return v
	}

	if v, ok := a.defaults.GetArgument(name); ok {
		return v
	}

	return value.Empty
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

		v, ok := a.value.GetOption(name)
		if !ok {
			switch {
			case opt.HasDefault():
				a.defaults.SetOption(name, opt.Default)
				continue
			case opt.IsRequired():
				return errs.Option(name, errs.ErrRequired)
			default:
				continue
			}
		}

		if err := opt.Validate(v); err != nil {
			return errs.Option(name, err)
		}

		a.SetOption(name, v)
	}

	return nil
}

func (a *Array) bindArguments(ctx context.Context, def *Definition) error {
	for pos, name := range def.Arguments() {
		arg, err := def.Argument(pos)
		if err != nil {
			return err
		}

		v, ok := a.value.GetArgument(name)
		if !ok {
			switch {
			case arg.HasDefault():
				a.defaults.SetArgument(name, arg.Default)
				continue
			case arg.IsRequired():
				return errs.Argument(name, errs.ErrRequired)
			default:
				continue
			}
		}

		if err := arg.Validate(v); err != nil {
			return errs.Argument(name, err)
		}

		a.SetArgument(name, v)
	}

	return nil
}

type array struct {
	opts map[string]value.Append
	args map[string]value.Append
	mu   sync.Mutex
}

func (a *array) GetOption(name string) (value.Append, bool) {
	v, ok := a.opts[name]

	return v, ok
}

func (a *array) SetOption(name string, v interface{}) {
	if a.opts == nil {
		a.opts = make(map[string]value.Append)
	}

	a.mu.Lock()
	a.opts[name] = value.New(v)
	a.mu.Unlock()
}

func (a *array) LenArguments() int {
	return len(a.args)
}

func (a *array) GetArgument(name string) (value.Append, bool) {
	v, ok := a.args[name]

	return v, ok
}

func (a *array) SetArgument(name string, v interface{}) {
	if a.args == nil {
		a.args = make(map[string]value.Append)
	}

	a.mu.Lock()
	a.args[name] = value.New(v)
	a.mu.Unlock()
}
