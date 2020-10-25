package input

import (
	"context"
)

type ReadInput interface {
	Bind(ctx context.Context, def *Definition) error

	ReadOption(ctx context.Context, name string) (Value, error)
	SetOption(name string, value Value)

	ReadArgument(ctx context.Context, name string) (Value, error)
	SetArgument(name string, value Value)
}

type Input interface {
	Option(ctx context.Context, name string) Value
	Argument(ctx context.Context, name string) Value
	ReadInput
}
