package console

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/output"
)

type (
	Action    func(ctx context.Context, input input.Input, output output.Output) error
	Handle    func(ctx context.Context, in input.Input, out output.Output, n Action) error
	Configure func(ctx context.Context, cfg *input.Definition) error
	Prepare   func(ctx context.Context, cfg *input.Definition, n Configure) error
	Option    func(*Command)
)

// WithPrepare append middleware for configuration command.
func WithPrepare(p ...Prepare) Option {
	return func(c *Command) {
		if c.Prepare != nil {
			p = append([]Prepare{c.Prepare}, p...)
		}

		c.Prepare = ChainPrepare(p...)
	}
}

// WithHandle append middleware for executed command.
func WithHandle(h ...Handle) Option {
	return func(c *Command) {
		if c.Handle != nil {
			h = append([]Handle{c.Handle}, h...)
		}

		c.Handle = ChainHandle(h...)
	}
}

// WithHidden sets hidden command.
func WithHidden(v bool) Option {
	return func(c *Command) {
		c.Hidden = v
	}
}

// WithName sets name command.
func WithName(n string) Option {
	return func(c *Command) {
		c.Name = n
	}
}

type Command struct {
	// The name of the command.
	Name string
	// A short description of the usage of this command.
	Description string
	// A longer explanation of how the command works.
	Help string
	// Vervion command.
	Version string
	// Boolean to hide this command from help or completion.
	Hidden bool
	// Configures the current command.
	Configure Configure
	// The middleware for configures current command.
	Prepare Prepare
	// The function to call when this command is invoked.
	Execute Action
	// The middleware for executes current command.
	Handle Handle
}

func (c *Command) String() string {
	return fmt.Sprintf("name: %s, version: %s", c.Name, c.Version)
}

// With creates new command by parent and options.
func (c *Command) With(opts ...Option) *Command {
	cmd := &Command{
		Name:        c.Name,
		Description: c.Description,
		Help:        c.Help,
		Version:     c.Version,
		Hidden:      c.Hidden,
		Configure:   c.Configure,
		Prepare:     c.Prepare,
		Execute:     c.Execute,
		Handle:      c.Handle,
	}

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

// Run run command with input and output.
func (c *Command) Run(ctx context.Context, in input.Input, out output.Output) error {
	if c.Handle != nil {
		return c.Handle(ctx, in, out, c.Execute)
	}

	return c.Execute(ctx, in, out)
}

// Init configures command.
func (c *Command) Init(ctx context.Context, cfg *input.Definition) error {
	switch {
	case c.Prepare != nil && c.Configure != nil:
		return c.Prepare(ctx, cfg, c.Configure)
	case c.Prepare != nil:
		return c.Prepare(ctx, cfg, func(_ context.Context, _ *input.Definition) error {
			return nil
		})
	case c.Configure != nil:
		return c.Configure(ctx, cfg)
	default:
		return nil
	}
}

// ChainPrepare creates middleware for configures command.
func ChainPrepare(prepare ...Prepare) Prepare {
	n := len(prepare)
	if n == 1 {
		return prepare[0]
	}

	if n > 1 {
		lastI := n - 1

		return func(ctx context.Context, def *input.Definition, next Configure) error {
			var (
				chainHandler func(context.Context, *input.Definition) error
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentDef *input.Definition) error {
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

	return func(ctx context.Context, cfg *input.Definition, next Configure) error {
		return next(ctx, cfg)
	}
}

// ChainHandle creates middleware for executes command.
func ChainHandle(handlers ...Handle) Handle {
	n := len(handlers)
	if n == 1 {
		return handlers[0]
	}

	if n > 1 {
		lastI := n - 1

		return func(ctx context.Context, in input.Input, out output.Output, next Action) error {
			var (
				chainHandler func(context.Context, input.Input, output.Output) error
				curI         int
			)

			chainHandler = func(currentCtx context.Context, currentIn input.Input, currentOut output.Output) error {
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

	return func(ctx context.Context, in input.Input, out output.Output, next Action) error {
		return next(ctx, in, out)
	}
}
