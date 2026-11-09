package help

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/param"
	"github.com/go-4devs/config/provider/arg"
	"github.com/go-4devs/config/validator"
	"github.com/go-4devs/config/value"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/errs"
	"github.com/go-4devs/console/internal/registry"
	"github.com/go-4devs/console/output"
	"github.com/go-4devs/console/output/descriptor"
	"github.com/go-4devs/console/setting"
)

//nolint:gochecknoinits
func init() {
	err := registry.Add(Command())
	if err != nil {
		panic(err)
	}
}

const (
	ArgumentCommandName = "command_name"
	OptionFormat        = "format"
	Name                = "help"
)

func Command() command.Command {
	return command.New(
		Name,
		"Displays help for a command",
		Execute,
		command.Configure(Configure),
		command.Help(Help),
	)
}

func Configure(_ context.Context, config config.Definition) error {
	formats := descriptor.Descriptors()
	config.
		Add(
			arg.String(ArgumentCommandName, "The command name", arg.Default(value.New("help"))),
			option.String(OptionFormat, fmt.Sprintf("The output format (%s)", strings.Join(formats, ", ")),
				option.Required,
				option.Default(value.New(descriptor.FormatTxt)),
				validator.Valid(
					validator.NotBlank,
					validator.Enum(formats...),
				),
			),
		)

	return nil
}

func Execute(ctx context.Context, in config.Provider, out output.Output) error {
	var err error

	cfg := read{Provider: in}
	name := cfg.Value(ctx, ArgumentCommandName).String()
	format := cfg.Value(ctx, OptionFormat).String()

	des, err := descriptor.Find(format)
	if err != nil {
		return fmt.Errorf("find descriptor[%v]: %w", format, err)
	}

	cmd, err := registry.Current(ctx).Find(name)
	if err != nil {
		return fmt.Errorf("find cmd: %w", err)
	}

	def := definition.New()
	command.Default(def)

	if err := cmd.Configure(ctx, def); err != nil {
		return fmt.Errorf("init cmd: %w", err)
	}

	var bin string
	if len(os.Args) > 0 {
		bin = os.Args[0]
	}

	help, err := setting.Help(cmd, setting.HelpData(bin, cmd.Name()))
	if err != nil {
		return fmt.Errorf("create help:%w", err)
	}

	hasUsage := true

	usage, err := setting.Usage(cmd, setting.UsageData(cmd.Name(), def))
	if err != nil {
		if !errors.Is(err, errs.ErrNotFound) {
			return fmt.Errorf("create usage:%w", err)
		}

		hasUsage = false
	}

	derr := des.Command(ctx, out, descriptor.Command{
		Bin:         bin,
		Name:        cmd.Name(),
		Description: setting.Description(cmd),
		Help:        help,
		Usage: func() (string, bool) {
			return usage, hasUsage
		},
		Options: def.With(param.New(descriptor.TxtStyle())),
	})
	if derr != nil {
		return fmt.Errorf("descriptor help:%w", derr)
	}

	return nil
}

const tpl = `The <info>%[2]s</info> command displays help for a given command:
  <info>%[1]s %[2]s list</info>
You can also output the help in other formats by using the <comment>--format</comment> option:
  <info>%[1]s %[2]s --format=txt list</info>
To display the list of available commands, please use the <info>list</info> command.
`

func Help(data setting.HData) (string, error) {
	return fmt.Sprintf(tpl, data.Bin, data.Name), nil
}

type read struct {
	config.Provider
}

func (r read) Value(ctx context.Context, key ...string) config.Value {
	val, err := r.Provider.Value(ctx, key...)
	if err != nil {
		return value.Empty{Err: err}
	}

	return val
}
