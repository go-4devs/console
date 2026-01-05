package console

import (
	"context"
	"errors"
	"fmt"
	"log"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/provider/chain"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/command/help"
	cerr "gitoa.ru/go-4devs/console/errors"
	"gitoa.ru/go-4devs/console/internal/registry"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/param"
)

// Execute the current command with option.
func Execute(ctx context.Context, cmd command.Command, opts ...func(*App)) {
	opts = append([]func(*App){WithSkipArgs(1)}, opts...)
	New(opts...).exec(ctx, cmd)
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
		log.Print(berr)

		return showHelp(ctx, cmd, in, output.Ansi(out))
	}

	out = command.Verbose(ctx, in, out)

	if command.IsShowVersion(ctx, in) {
		out.Println(ctx, "command <comment>", cmd.Name(), "</comment> version: <info>", param.Version(cmd), "</info>")

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

func showHelp(ctx context.Context, cmd command.Command, in config.Provider, out output.Output) error {
	arr := &memory.Map{}
	arr.SetOption(value.New(cmd.Name()), help.ArgumentCommandName)
	arr.SetOption(value.New(false), command.OptionHelp)

	if _, err := registry.Find(cmd.Name()); errors.Is(err, cerr.ErrNotFound) {
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
