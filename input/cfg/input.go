package cfg

import (
	"context"
	"errors"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
)

type Resolver func(ctx context.Context, name string) (config.Value, error)

func New(resolver func(ctx context.Context, name string) (config.Value, error)) *Input {
	return &Input{
		resolver: resolver,
	}
}

type Input struct {
	resolver Resolver
}

type Value struct {
	config.Value
}

func (v Value) Any() interface{} {
	var out interface{}
	_ = v.Value.Unmarshal(&out)

	return out
}

func (i *Input) Option(ctx context.Context, name string) value.Value {
	val, err := i.resolver(ctx, name)
	if errors.Is(err, config.ErrVariableNotFound) {
		return value.Empty()
	}

	return value.Read{ParseValue: Value{Value: val}}
}

func (i *Input) Argument(ctx context.Context, name string) value.Value {
	return value.Empty()
}

func (i *Input) Bind(ctx context.Context, def *input.Definition) error {
	return nil
}
