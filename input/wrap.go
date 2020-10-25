package input

import (
	"context"

	"gitoa.ru/go-4devs/console/input/value"
)

type Wrap struct {
	Input
	Array
}

func (w *Wrap) Option(ctx context.Context, name string) value.Value {
	if v, ok := w.Array.GetOption(name); ok {
		return v
	}

	return w.Input.Option(ctx, name)
}

func (w *Wrap) Argument(ctx context.Context, name string) value.Value {
	if v, ok := w.Array.GetArgument(name); ok {
		return v
	}

	return w.Input.Argument(ctx, name)
}

func (w *Wrap) Bind(ctx context.Context, def *Definition) error {
	if err := w.Input.Bind(ctx, def); err != nil {
		return err
	}

	return w.Array.Bind(ctx, def)
}
