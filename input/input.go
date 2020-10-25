package input

import (
	"context"

	"gitoa.ru/go-4devs/console/input/value"
)

type ReadInput interface {
	Bind(ctx context.Context, def *Definition) error

	ReadOption(ctx context.Context, name string) (value.Value, error)
	SetOption(name string, v value.Value)

	ReadArgument(ctx context.Context, name string) (value.Value, error)
	SetArgument(name string, v value.Value)
}

type Input interface {
	Option(ctx context.Context, name string) value.Value
	Argument(ctx context.Context, name string) value.Value
	ReadInput
}
