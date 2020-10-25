package console

import (
	"context"
	"errors"
	"os"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/verbosity"
	"gitoa.ru/go-4devs/console/output/wrap"
)

const (
	verboseTrace = 3
	verboseDebug = 2
	verboseInfo  = 3
)

// Execute the current command with option.
func Execute(ctx context.Context, cmd *Command, opts ...func(*App)) {
	opts = append([]func(*App){WithSkipArgs(1)}, opts...)
	New(opts...).exec(ctx, cmd)
}

// Run current command by input and output.
func Run(ctx context.Context, cmd *Command, in input.Input, out output.Output) error {
	def := input.NewDefinition()

	if err := cmd.Init(ctx, def); err != nil {
		return err
	}

	if err := in.Bind(ctx, Default(def)); err != nil {
		ansi(ctx, in, out).Print(ctx, "<error>\n\n   ", err, "\n</error>\n")

		return showHelp(ctx, cmd, in, wrap.Ansi(out))
	}

	out = ansi(ctx, in, out)

	out = verbose(ctx, in, out)

	if in.Option(ctx, "version").Bool() {
		version := cmd.Version
		if version == "" {
			version = "unknown"
		}

		out.Println(ctx, "command <comment>", cmd.Name, "</comment> version: <info>", version, "</info>")

		return nil
	}

	if in.Option(ctx, "help").Bool() {
		return showHelp(ctx, cmd, in, out)
	}

	return cmd.Run(ctx, in, out)
}

func ansi(ctx context.Context, in input.Input, out output.Output) output.Output {
	switch {
	case in.Option(ctx, "ansi").Bool():
		out = wrap.Ansi(out)
	case in.Option(ctx, "no-ansi").Bool():
		out = wrap.None(out)
	case lookupEnv("NO_COLOR"):
		out = wrap.None(out)
	default:
		out = wrap.Ansi(out)
	}

	return out
}

func lookupEnv(name string) bool {
	v, has := os.LookupEnv(name)

	return has && v == "true"
}

func verbose(ctx context.Context, in input.Input, out output.Output) output.Output {
	switch {
	case in.Option(ctx, "quiet").Bool():
		out = verbosity.Quiet()
	default:
		v := in.Option(ctx, "verbose").Bools()

		switch {
		case len(v) == verboseInfo:
			out = verbosity.Verb(out, output.VerbosityInfo)
		case len(v) == verboseDebug:
			out = verbosity.Verb(out, output.VerbosityDebug)
		case len(v) >= verboseTrace:
			out = verbosity.Verb(out, output.VerbosityTrace)
		default:
			out = verbosity.Verb(out, output.VerbosityNorm)
		}
	}

	return out
}

func showHelp(ctx context.Context, cmd *Command, in input.Input, out output.Output) error {
	in.SetArgument(HelpArgumentCommandName, value.New(cmd.Name))
	in.SetOption("help", value.New(false))

	if _, err := Find(cmd.Name); errors.Is(err, ErrNotFound) {
		register(cmd)
	}

	help, err := Find(CommandHelp)
	if err != nil {
		return err
	}

	return Run(ctx, help, in, out)
}

// Default options and argument command.
func Default(d *input.Definition) *input.Definition {
	return d.SetOptions(
		option.Bool("no-ansi", "Disable ANSI output"),
		option.Bool("ansi", "Do not ask any interactive question"),
		option.Bool("version", "Display this application version", option.Short("V")),
		option.Bool("help", "Display this help message", option.Short("h")),
		option.Bool("verbose",
			"Increase the verbosity of messages: -v for info output, -vv for debug and -vvv for trace",
			option.Short("v"), option.Array),
		option.Bool("quiet", "Do not output any message", option.Short("q")),
	)
}
