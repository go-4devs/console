// Code generated gitoa.ru/go-4devs/config DO NOT EDIT.
package dump

import (
	"context"
	"fmt"
	"gitoa.ru/go-4devs/config"
)

func WithRConfigureConfigHandle(fn func(context.Context, error)) func(*RConfigureConfig) {
	return func(ci *RConfigureConfig) {
		ci.handle = fn
	}
}

func NewRConfigureConfig(prov config.Provider, opts ...func(*RConfigureConfig)) RConfigureConfig {
	i := RConfigureConfig{
		Provider: prov,
		handle: func(_ context.Context, err error) {
			fmt.Printf("RConfigureConfig:%v", err)
		},
	}

	for _, opt := range opts {
		opt(&i)
	}

	return i
}

type RConfigureConfig struct {
	config.Provider
	handle func(context.Context, error)
}

// readCommandName command name.
func (i RConfigureConfig) readCommandName(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "command-name")
	if err != nil {
		return v, fmt.Errorf("read [%v]:%w", []string{"command-name"}, err)

	}

	return val.ParseString()

}

// ReadCommandName command name.
func (i RConfigureConfig) ReadCommandName(ctx context.Context) (string, error) {
	return i.readCommandName(ctx)
}

// CommandName command name.
func (i RConfigureConfig) CommandName(ctx context.Context) string {
	val, err := i.readCommandName(ctx)
	if err != nil {
		i.handle(ctx, err)
	}

	return val
}

// readFormat format.
func (i RConfigureConfig) readFormat(ctx context.Context) (v string, e error) {
	val, err := i.Value(ctx, "format")
	if err != nil {
		i.handle(ctx, err)

		return "arg", nil
	}

	return val.ParseString()

}

// ReadFormat format.
func (i RConfigureConfig) ReadFormat(ctx context.Context) (string, error) {
	return i.readFormat(ctx)
}

// Format format.
func (i RConfigureConfig) Format(ctx context.Context) string {
	val, err := i.readFormat(ctx)
	if err != nil {
		i.handle(ctx, err)
	}

	return val
}
