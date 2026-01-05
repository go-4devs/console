package command

import (
	"context"
	"math"
	"os"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/verbosity"
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
	verboseTrace = 3
	verboseDebug = 2
	verboseInfo  = 1
)

const (
	defaultOptionsPosition = math.MaxUint64 / 2
)

// Default options and argument command.
func Default(def config.Definition) {
	def.Add(
		option.Bool(OptionNoAnsi, "Disable ANSI output", option.Position(defaultOptionsPosition)),
		option.Bool(OptionAnsi, "Do not ask any interactive question", option.Position(defaultOptionsPosition)),
		option.Bool(OptionVersion, "Display this application version", option.Short('V'), option.Position(defaultOptionsPosition)),
		option.Bool(OptionHelp, "Display this help message", option.Short('h'), option.Position(defaultOptionsPosition)),
		option.Bool(OptionVerbose,
			"Increase the verbosity of messages: -v for info output, -vv for debug and -vvv for trace",
			option.Short('v'), option.Slice, option.Position(defaultOptionsPosition)),
		option.Bool(OptionQuiet, "Do not output any message", option.Short('q'), option.Position(defaultOptionsPosition)),
	)
}

func IsShowVersion(ctx context.Context, in config.Provider) bool {
	v, err := in.Value(ctx, OptionVersion)
	if err != nil {
		return false
	}

	return v.Bool()
}

func IsShowHelp(ctx context.Context, in config.Provider) bool {
	v, err := in.Value(ctx, OptionHelp)
	if err != nil {
		return false
	}

	return v.Bool()
}

func Ansi(ctx context.Context, in config.Provider, out output.Output) output.Output {
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

func Verbose(ctx context.Context, in config.Provider, out output.Output) output.Output {
	out = Ansi(ctx, in, out)

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

func ReadValue(ctx context.Context, in config.Provider, path ...string) config.Value {
	val, err := in.Value(ctx, path...)
	if err != nil {
		return value.EmptyValue()
	}

	return val
}
