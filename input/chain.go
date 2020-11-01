package input

import (
	"context"
	"log"

	"gitoa.ru/go-4devs/console/input/value"
)

func Chain(c ...Input) Input {
	return chain(c)
}

type chain []Input

func (c chain) Option(ctx context.Context, name string) value.Value {
	log.Println(name, len(c))
	for _, in := range c {
		log.Printf("%T\n", in)
		if val := in.Option(ctx, name); !value.IsEmpty(val) {
			return val
		}
	}

	return value.Empty
}

func (c chain) Argument(ctx context.Context, name string) value.Value {
	for _, in := range c {
		if val := in.Argument(ctx, name); !value.IsEmpty(val) {
			return val
		}
	}

	return value.Empty
}

func (c chain) Bind(ctx context.Context, def *Definition) error {
	for _, input := range c {
		if err := input.Bind(ctx, def); err != nil {
			return err
		}
	}

	return nil
}
