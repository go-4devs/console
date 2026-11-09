package console

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/provider/chain"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/value"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/command/help"
	"github.com/go-4devs/console/errs"
	"github.com/go-4devs/console/internal/registry"
	"github.com/go-4devs/console/output"
	"github.com/go-4devs/console/setting"
)

// Execute the current command with option.
func Execute(ctx context.Context, cmd command.Command, opts ...func(*App)) {
	opts = append([]func(*App){WithSkipArgs(1)}, opts...)
	New(opts...).exec(ctx, command.With(cmd, command.EmptyUsage))
}

// Run current command by input and output.
func Run(ctx context.Context, cmd command.Command, in config.BindProvider, out output.Output) error {
	def := definition.New()

	err := cmd.Configure(ctx, def)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	command.Default(def)

	berr := in.Bind(ctx, config.NewVars(def.Options()...))
	if berr != nil {
		printErr(ctx, in, out, berr)

		if herr := showHelp(ctx, cmd, in, output.Ansi(out)); herr != nil {
			return herr
		}

		return printedError{err: berr}
	}

	out = command.Verbose(ctx, in, out)

	if command.IsShowVersion(ctx, in) {
		out.Println(ctx, "command <comment>", cmd.Name(), "</comment> version: <info>", setting.Version(cmd), "</info>")

		return nil
	}

	if command.IsShowHelp(ctx, in) {
		return showHelp(ctx, cmd, in, out)
	}

	if err := cmd.Execute(ctx, in, out); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// printedError was already written to the output. App must not print it again.
type printedError struct{ err error }

func (e printedError) Error() string { return e.err.Error() }

func (e printedError) Unwrap() error { return e.err }

func showHelp(ctx context.Context, cmd command.Command, in config.Provider, out output.Output) error {
	arr := &memory.Map{}
	arr.SetOption(value.New(cmd.Name()), help.ArgumentCommandName)
	arr.SetOption(value.New(false), command.OptionHelp)

	if _, err := registry.Find(cmd.Name()); errors.Is(err, errs.ErrNotFound) {
		_ = registry.Add(cmd)
	}

	help, err := registry.Find(help.Name)
	if err != nil {
		return fmt.Errorf("%w", err)
	}

	w := chain.New(arr, in)

	return Run(ctx, help, w, out)
}

func ReadValue(ctx context.Context, in config.Provider, path ...string) config.Value {
	val, err := in.Value(ctx, path...)
	if err != nil {
		return value.EmptyValue()
	}

	return val
}
