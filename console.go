package console

import (
	"context"
	"errors"
	"log"
	"math"
	"os"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/provider/chain"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/verbosity"
)

const (
	verboseTrace = 3
	verboseDebug = 2
	verboseInfo  = 1
)

const (
	OptionHelp    = "help"
	OptionVersion = "version"
	OptionAnsi    = "ansi"
	OptionNoAnsi  = "no-ansi"
	OptionQuiet   = "quiet"
	OptionVerbose = "verbose"
)

const (
	defaultOptionsPosition = math.MaxUint64 / 2
)

// Execute the current command with option.
func Execute(ctx context.Context, cmd *Command, opts ...func(*App)) {
	opts = append([]func(*App){WithSkipArgs(1)}, opts...)
	New(opts...).exec(ctx, cmd)
}

// Run current command by input and output.
func Run(ctx context.Context, cmd *Command, in config.BindProvider, out output.Output) error {
	def := definition.New()

	err := cmd.Init(ctx, def)
	if err != nil {
		return err
	}

	def.Add(Default()...)

	berr := in.Bind(ctx, config.NewVars(def.Options()...))
	if berr != nil {
		log.Print(berr)

		return showHelp(ctx, cmd, in, output.Ansi(out))
	}

	out = ansi(ctx, in, out)

	out = verbose(ctx, in, out)

	if ReadValue(ctx, in, OptionVersion).Bool() {
		version := cmd.Version
		if version == "" {
			version = "unknown"
		}

		out.Println(ctx, "command <comment>", cmd.Name, "</comment> version: <info>", version, "</info>")

		return nil
	}

	if ReadValue(ctx, in, OptionHelp).Bool() {
		return showHelp(ctx, cmd, in, out)
	}

	return cmd.Run(ctx, in, out)
}

func ansi(ctx context.Context, in config.Provider, out output.Output) output.Output {
	switch {
	case ReadValue(ctx, in, OptionAnsi).Bool():
		out = output.Ansi(out)
	case ReadValue(ctx, in, OptionNoAnsi).Bool():
		out = output.None(out)
	case lookupEnv("NO_COLOR"):
		out = output.None(out)
	default:
		out = output.Ansi(out)
	}

	return out
}

func lookupEnv(name string) bool {
	v, has := os.LookupEnv(name)

	return has && v == "true"
}

func verbose(ctx context.Context, in config.Provider, out output.Output) output.Output {
	switch {
	case ReadValue(ctx, in, OptionQuiet).Bool():
		out = output.Quiet()
	default:
		var verb []bool

		_ = ReadValue(ctx, in, OptionVerbose).Unmarshal(&verb)

		switch {
		case len(verb) == verboseInfo:
			out = output.Verbosity(out, verbosity.Info)
		case len(verb) == verboseDebug:
			out = output.Verbosity(out, verbosity.Debug)
		case len(verb) >= verboseTrace:
			out = output.Verbosity(out, verbosity.Trace)
		default:
			out = output.Verbosity(out, verbosity.Norm)
		}
	}

	return out
}

func showHelp(ctx context.Context, cmd *Command, in config.Provider, out output.Output) error {
	arr := &memory.Map{}
	arr.SetOption(value.New(cmd.Name), ArgumentCommandName)
	arr.SetOption(value.New(false), OptionHelp)

	if _, err := Find(cmd.Name); errors.Is(err, ErrNotFound) {
		register(cmd)
	}

	help, err := Find(CommandHelp)
	if err != nil {
		return err
	}

	w := chain.New(arr, in)

	return Run(ctx, help, w, out)
}

// Default options and argument command.
func Default() []config.Option {
	return []config.Option{
		option.Bool(OptionNoAnsi, "Disable ANSI output", option.Position(defaultOptionsPosition)),
		option.Bool(OptionAnsi, "Do not ask any interactive question", option.Position(defaultOptionsPosition)),
		option.Bool(OptionVersion, "Display this application version", option.Short('V'), option.Position(defaultOptionsPosition)),
		option.Bool(OptionHelp, "Display this help message", option.Short('h'), option.Position(defaultOptionsPosition)),
		option.Bool(OptionVerbose,
			"Increase the verbosity of messages: -v for info output, -vv for debug and -vvv for trace",
			option.Short('v'), option.Slice, option.Position(defaultOptionsPosition)),
		option.Bool(OptionQuiet, "Do not output any message", option.Short('q'), option.Position(defaultOptionsPosition)),
	}
}

func ReadValue(ctx context.Context, in config.Provider, path ...string) config.Value {
	val, err := in.Value(ctx, path...)
	if err != nil {
		return value.EmptyValue()
	}

	return val
}
