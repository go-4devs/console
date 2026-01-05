package command

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/output"
)

// ChainPrepare creates middleware for configures command.
func ChainPrepare(prepare ...PrepareFn) PrepareFn {
	num := len(prepare)
	if num == 1 {
		return prepare[0]
	}

	if num > 1 {
		lastI := num - 1

		return func(ctx context.Context, def config.Definition, next ConfigureFn) error {
			var (
				chainHandler func(context.Context, config.Definition) error
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentDef config.Definition) error {
				if curI == lastI {
					return next(currentCtx, currentDef)
				}

				curI++
				err := prepare[curI](currentCtx, currentDef, chainHandler)
				curI--

				return err
			}

			return prepare[0](ctx, def, chainHandler)
		}
	}

	return func(ctx context.Context, cfg config.Definition, next ConfigureFn) error {
		return next(ctx, cfg)
	}
}

// ChainHandle creates middleware for executes command.
func ChainHandle(handlers ...HandleFn) HandleFn {
	num := len(handlers)
	if num == 1 {
		return handlers[0]
	}

	if num > 1 {
		lastI := num - 1

		return func(ctx context.Context, in config.Provider, out output.Output, next ExecuteFn) error {
			var (
				chainHandler func(context.Context, config.Provider, output.Output) error
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentIn config.Provider, currentOut output.Output) error {
				if curI == lastI {
					return next(currentCtx, currentIn, currentOut)
				}

				curI++
				err := handlers[curI](currentCtx, currentIn, currentOut, chainHandler)
				curI--

				return err
			}

			return handlers[0](ctx, in, out, chainHandler)
		}
	}

	return func(ctx context.Context, in config.Provider, out output.Output, next ExecuteFn) error {
		return next(ctx, in, out)
	}
}
