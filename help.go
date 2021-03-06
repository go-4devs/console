package console

import (
	"context"
	"fmt"
	"os"
	"strings"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/input/validator"
	"gitoa.ru/go-4devs/console/input/value/flag"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/descriptor"
)

//nolint: gochecknoinits
func init() {
	MustRegister(help())
}

const (
	HelpArgumentCommandName = "command_name"
	helpOptFormat           = "format"
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
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			var err error
			name := in.Argument(ctx, HelpArgumentCommandName).String()
			format := in.Option(ctx, helpOptFormat).String()

			des, err := descriptor.Find(format)
			if err != nil {
				return err
			}

			cmd, err := Find(name)
			if err != nil {
				return err
			}

			def := input.NewDefinition()

			if err := cmd.Init(ctx, Default(def)); err != nil {
				return err
			}

			var bin string
			if len(os.Args) > 0 {
				bin = os.Args[0]
			}

			return des.Command(ctx, out, descriptor.Command{
				Bin:         bin,
				Name:        cmd.Name,
				Description: cmd.Description,
				Help:        cmd.Help,
				Definition:  def,
			})
		},
		Configure: func(ctx context.Context, config *input.Definition) error {
			formats := descriptor.Descriptors()
			config.
				SetArguments(
					argument.New(HelpArgumentCommandName, "The command name", argument.Default("help")),
				).
				SetOptions(
					option.New(helpOptFormat, fmt.Sprintf("The output format (%s)", strings.Join(formats, ", ")),
						option.Required,
						option.Default(formats[0]),
						option.Valid(
							validator.NotBlank(flag.String),
							validator.Enum(formats...),
						),
					),
				)

			return nil
		},
	}
}
