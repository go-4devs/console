package command

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/setting"
)

type (
	ExecuteFn   func(ctx context.Context, input config.Provider, output output.Output) error
	HandleFn    func(ctx context.Context, in config.Provider, out output.Output, n ExecuteFn) error
	ConfigureFn func(ctx context.Context, cfg config.Definition) error
	PrepareFn   func(ctx context.Context, cfg config.Definition, n ConfigureFn) error

	Option func(*Command)
)

func Configure(fn ConfigureFn) Option {
	return func(c *Command) {
		c.configure = fn
	}
}

func Version(in string) Option {
	return func(c *Command) {
		c.Setting = setting.WithVersion(in)(c.Setting)
	}
}

func Hidden(c *Command) {
	c.Setting = setting.Hidden(c.Setting)
}

func Help(fn setting.HelpFn) Option {
	return func(c *Command) {
		c.Setting = setting.WithHelp(fn)(c.Setting)
	}
}

func WithName(name string) Option {
	return func(c *Command) {
		c.name = name
	}
}

func Handle(fn HandleFn) Option {
	return func(c *Command) {
		handle := c.handle
		c.handle = ChainHandle(fn, handle)
	}
}

func Prepare(fn PrepareFn) Option {
	return func(c *Command) {
		prepare := c.prepare
		c.prepare = ChainPrepare(fn, prepare)
	}
}

func Usage(fn setting.UsageFn) Option {
	return func(c *Command) {
		c.Setting = setting.WithUsage(fn)(c.Setting)
	}
}

func EmptyUsage(cmd *Command) {
	cmd.Setting = setting.WithUsage(func(setting.UData) (string, error) {
		return "", nil
	})(cmd.Setting)
}

func New(name, desc string, execute ExecuteFn, opts ...Option) Command {
	cmd := Command{
		name:      name,
		execute:   execute,
		configure: emptyConfigure,
		handle:    emptyHandle,
		prepare:   emptyPrepare,
		Setting:   setting.New(setting.WithDescription(desc)),
	}

	for _, opt := range opts {
		opt(&cmd)
	}

	return cmd
}

type Command struct {
	setting.Setting

	name      string
	execute   ExecuteFn
	configure ConfigureFn
	prepare   PrepareFn
	handle    HandleFn
}

func (c Command) Name() string {
	return c.name
}

func (c Command) Execute(ctx context.Context, input config.Provider, output output.Output) error {
	return c.handle(ctx, input, output, c.execute)
}

func (c Command) Configure(ctx context.Context, cfg config.Definition) error {
	return c.prepare(ctx, cfg, c.configure)
}

func (c Command) IsZero() bool {
	return c.name == "" ||
		c.execute == nil ||
		c.configure == nil ||
		c.handle == nil ||
		c.prepare == nil
}

func (c Command) String() string {
	return fmt.Sprintf("command:%v, version:%v", c.Name(), setting.Version(c))
}

func With(parent Command, opts ...Option) Command {
	cmd := Command{
		Setting:   parent.Setting,
		name:      parent.Name(),
		execute:   parent.Execute,
		configure: parent.Configure,
		handle:    emptyHandle,
		prepare:   emptyPrepare,
	}

	for _, opt := range opts {
		opt(&cmd)
	}

	return cmd
}

func emptyPrepare(ctx context.Context, cfg config.Definition, n ConfigureFn) error {
	return n(ctx, cfg)
}

func emptyHandle(ctx context.Context, in config.Provider, out output.Output, n ExecuteFn) error {
	return n(ctx, in, out)
}

func emptyConfigure(context.Context, config.Definition) error {
	return nil
}
