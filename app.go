package console

import (
	"context"
	"os"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/output"
)

// WithOutput sets outpu,^ by default output os.Stdout.
func WithOutput(out output.Output) func(*App) {
	return func(a *App) {
		a.out = out
	}
}

// WithInput sets input, by default creates inpur by os.Args.
func WithInput(in input.Input) func(*App) {
	return func(a *App) {
		a.in = in
	}
}

// WithSkipArgs sets how many arguments are passed. For example, you don't need to pass the name of a single command.
func WithSkipArgs(l int) func(*App) {
	return func(a *App) {
		a.skipArgv = l
	}
}

// WithExit sets exit callback by default os.Exit.
func WithExit(f func(int)) func(*App) {
	return func(a *App) {
		a.exit = f
	}
}

// New creates and configure new console app.
func New(opts ...func(*App)) *App {
	a := &App{
		out:  output.Stdout(),
		exit: os.Exit,
	}

	for _, opt := range opts {
		opt(a)
	}

	if a.in == nil {
		skip := 2

		switch {
		case a.skipArgv > 0 && len(os.Args) > a.skipArgv:
			skip = a.skipArgv
		case a.skipArgv > 0:
			skip = len(os.Args)
		case len(os.Args) == 1:
			skip = 1
		case len(os.Args) > 1 && os.Args[1][0] == '-':
			skip = 1
		}

		a.in = &input.Argv{
			Args: os.Args[skip:],
		}
	}

	return a
}

// App is collection of command and configure env.
type App struct {
	cmds     []*Command
	out      output.Output
	in       input.Input
	skipArgv int
	exit     func(int)
}

// Add add or replace command.
func (a *App) Add(cmds ...*Command) *App {
	a.cmds = append(a.cmds, cmds...)

	return a
}

// Execute run the command by name and arguments.
func (a *App) Execute(ctx context.Context) {
	for _, cmd := range a.cmds {
		register(cmd)
	}

	cmd, err := a.find(ctx)
	if err != nil {
		a.printError(ctx, err)

		if err := a.list(ctx); err != nil {
			a.printError(ctx, err)
		}

		a.exit(1)
	}

	a.exec(ctx, cmd)
}

func (a *App) exec(ctx context.Context, cmd *Command) {
	if err := Run(ctx, cmd, a.in, a.out); err != nil {
		a.printError(ctx, err)
		a.exit(1)
	}

	a.exit(0)
}

func (a *App) find(_ context.Context) (*Command, error) {
	if len(os.Args) < 2 || os.Args[1][1] == '-' {
		return Find(CommandList)
	}

	name := os.Args[1]

	return Find(name)
}

func (a *App) list(ctx context.Context) error {
	cmd, err := Find(CommandHelp)
	if err != nil {
		return err
	}
	in := &input.Wrap{
		Input: a.in,
	}
	in.SetArgument("command_name", value.New(CommandList))

	return Run(ctx, cmd, in, a.out)
}

func (a *App) printError(ctx context.Context, err error) {
	a.out.Println(ctx, "<error>\n\n  ", err, "\n</error>")
}
