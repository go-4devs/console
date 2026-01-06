package console

import (
	"context"
	"fmt"
	"os"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/provider/chain"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/command/help"
	"gitoa.ru/go-4devs/console/command/list"
	"gitoa.ru/go-4devs/console/internal/registry"
	"gitoa.ru/go-4devs/console/output"
)

// WithOutput sets outpu,^ by default output os.Stdout.
func WithOutput(out output.Output) func(*App) {
	return func(a *App) {
		a.out = out
	}
}

// WithInput sets input, by default creates inpur by os.Args.
func WithInput(in config.BindProvider) func(*App) {
	return func(a *App) {
		a.in = in
	}
}

// WithSkipArgs sets how many arguments are passed. For example, you don't need to pass the name of a single command.
func WithSkipArgs(l int) func(*App) {
	return WithInput(chain.New(
		arg.New(arg.WithArgs(os.Args[ResolveSkip(l):])),
		&memory.Default{}),
	)
}

// WithExit sets exit callback by default os.Exit.
func WithExit(f func(int)) func(*App) {
	return func(a *App) {
		a.exit = f
	}
}

func WithReplaceCommand(a *App) {
	a.registry = registry.Set
}

// New creates and configure new console app.
func New(opts ...func(*App)) *App {
	app := &App{
		out:  output.Stdout(),
		exit: os.Exit,
		in: chain.New(
			arg.New(arg.WithArgs(os.Args[ResolveSkip(0):])),
			&memory.Default{},
		),
		registry: registry.Add,
	}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

// App is collection of command and configure env.
type App struct {
	registry func(...command.Command) error
	out      output.Output
	in       config.BindProvider
	exit     func(int)
}

// Add add or replace command.
func (a *App) Add(cmds ...command.Command) *App {
	if err := a.registry(cmds...); err != nil {
		a.printError(context.Background(), err)
		a.exit(1)
	}

	return a
}

// Execute run the command by name and arguments.
func (a *App) Execute(ctx context.Context) {
	cmd, err := registry.Find(a.commandName())
	if err != nil {
		a.printError(ctx, err)

		err := a.list(ctx)
		if err != nil {
			a.printError(ctx, err)
		}

		a.exit(1)
	}

	a.exec(ctx, cmd)
}

func (a *App) exec(ctx context.Context, cmd command.Command) {
	err := Run(ctx, cmd, a.in, a.out)
	if err != nil {
		a.printError(ctx, err)
		a.exit(1)
	}

	a.exit(0)
}

func (a *App) commandName() string {
	name := list.Name
	if len(os.Args) > 1 && len(os.Args[1]) > 1 && os.Args[1][1] != '-' {
		name = os.Args[1]
	}

	return name
}

func (a *App) list(ctx context.Context) error {
	cmd, err := registry.Find(help.Name)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	arr := &memory.Map{}
	arr.SetOption(value.New(list.Name), help.ArgumentCommandName)
	in := chain.New(arr, a.in)

	return Run(ctx, cmd, in, a.out)
}

func (a *App) printError(ctx context.Context, err error) {
	printErr(ctx, a.in, a.out, err)
}

func printErr(ctx context.Context, in config.Provider, out output.Output, err error) {
	command.Ansi(ctx, in, out).Printf(ctx, "<error>\n\n  %v\n</error>\n", err)
}

func ResolveSkip(in int) int {
	res := 2

	switch {
	case in > 0 && len(os.Args) > in:
		res = in
	case in > 0:
		res = len(os.Args)
	case len(os.Args) == 1:
		res = 1
	case len(os.Args) > 1 && os.Args[1][0] == '-':
		res = 1
	}

	return res
}
