package input

import (
	"context"

	"gitoa.ru/go-4devs/console/input/value"
)

type Input interface {
	Option(ctx context.Context, name string) value.Value
	Argument(ctx context.Context, name string) value.Value
	Bind(ctx context.Context, def *Definition) error
}
