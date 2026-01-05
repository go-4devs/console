package console

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/descriptor"
)

//nolint:gochecknoinits
func init() {
	MustRegister(help())
}

const (
	ArgumentCommandName = "command_name"
	OptionFormat        = "format"
)

func help() *Command {
	return &Command{
		Name:        CommandHelp,
		Description: `Displays help for a command`,
		Help: `
The <info>{{ .Name }}</info> command displays help for a given command:
  <info>{{ .Bin }} {{ .Name }} list</info>
You can also output the help in other formats by using the <comment>--format</comment> option:
  <info>{{ .Bin }} {{ .Name }} --format=xml list</info>
To display the list of available commands, please use the <info>list</info> command.
`,
		Execute: func(ctx context.Context, in config.Provider, out output.Output) error {
			var err error

			name := ReadValue(ctx, in, ArgumentCommandName).String()
			format := ReadValue(ctx, in, OptionFormat).String()

			des, err := descriptor.Find(format)
			if err != nil {
				return fmt.Errorf("find descriptor[%v]: %w", format, err)
			}

			cmd, err := Find(name)
			if err != nil {
				return fmt.Errorf("find cmd: %w", err)
			}

			def := definition.New()
			def.Add(Default()...)

			if err := cmd.Init(ctx, def); err != nil {
				return fmt.Errorf("init cmd: %w", err)
			}

			var bin string
			if len(os.Args) > 0 {
				bin = os.Args[0]
			}

			derr := des.Command(ctx, out, descriptor.Command{
				Bin:         bin,
				Name:        cmd.Name,
				Description: cmd.Description,
				Help:        cmd.Help,
				Options:     def.With(param.New(descriptor.TxtStyle())),
			})
			if derr != nil {
				return fmt.Errorf("descriptor help:%w", derr)
			}

			return nil
		},
		Configure: func(_ context.Context, config config.Definition) error {
			formats := descriptor.Descriptors()
			config.
				Add(
					arg.String(ArgumentCommandName, "The command name", arg.Default(value.New("help"))),
					option.String(OptionFormat, fmt.Sprintf("The output format (%s)", strings.Join(formats, ", ")),
						option.Required,
						option.Default(value.New(formats[0])),
						validator.Valid(
							validator.NotBlank,
							validator.Enum(formats...),
						),
					),
				)

			return nil
		},
	}
}
