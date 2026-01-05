package console

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/errors"
	"gitoa.ru/go-4devs/console/output"
)

type (
	Action    func(ctx context.Context, input config.Provider, output output.Output) error
	Handle    func(ctx context.Context, in config.Provider, out output.Output, n Action) error
	Configure func(ctx context.Context, cfg config.Definition) error
	Prepare   func(ctx context.Context, cfg config.Definition, n Configure) error
	Option    func(*Command)
)

// WithPrepare append middleware for configuration command.
func WithPrepare(prepares ...Prepare) Option {
	return func(c *Command) {
		if c.Prepare != nil {
			prepares = append([]Prepare{c.Prepare}, prepares...)
		}

		c.Prepare = ChainPrepare(prepares...)
	}
}

// WithHandle append middleware for executed command.
func WithHandle(handles ...Handle) Option {
	return func(c *Command) {
		if c.Handle != nil {
			handles = append([]Handle{c.Handle}, handles...)
		}

		c.Handle = ChainHandle(handles...)
	}
}

// WithHidden sets hidden command.
func WithHidden(hidden bool) Option {
	return func(c *Command) {
		c.Hidden = hidden
	}
}

// WithName sets name command.
func WithName(name string) Option {
	return func(c *Command) {
		c.Name = name
	}
}

func Wrap(cmd *Command) command.Command {
	opts := make([]command.Option, 0)
	if cmd.Hidden {
		opts = append(opts, command.Hidden)
	}

	opts = append(opts, command.Configure(cmd.Init))

	return command.New(cmd.Name, cmd.Description, cmd.Run, opts...)
}

// Deprecated: use command.New().
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
func (c *Command) Run(ctx context.Context, in config.Provider, out output.Output) error {
	if c.Execute == nil {
		return fmt.Errorf("%w", errors.ErrExecuteNil)
	}

	if c.Handle != nil {
		return c.Handle(ctx, in, out, c.Execute)
	}

	return c.Execute(ctx, in, out)
}

// Init configures command.
func (c *Command) Init(ctx context.Context, cfg config.Definition) error {
	switch {
	case c.Prepare != nil && c.Configure != nil:
		return c.Prepare(ctx, cfg, c.Configure)
	case c.Prepare != nil:
		return c.Prepare(ctx, cfg, func(_ context.Context, _ config.Definition) error {
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
	num := len(prepare)
	if num == 1 {
		return prepare[0]
	}

	if num > 1 {
		lastI := num - 1

		return func(ctx context.Context, def config.Definition, next Configure) error {
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

	return func(ctx context.Context, cfg config.Definition, next Configure) error {
		return next(ctx, cfg)
	}
}

// ChainHandle creates middleware for executes command.
func ChainHandle(handlers ...Handle) Handle {
	num := len(handlers)
	if num == 1 {
		return handlers[0]
	}

	if num > 1 {
		lastI := num - 1

		return func(ctx context.Context, in config.Provider, out output.Output, next Action) error {
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

	return func(ctx context.Context, in config.Provider, out output.Output, next Action) error {
		return next(ctx, in, out)
	}
}
