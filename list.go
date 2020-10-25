package console

import (
	"context"
	"fmt"
	"strings"

	"gitoa.ru/go-4devs/console/descriptor"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/validator"
)

const defaultLenNamespace = 2

//nolint: gochecknoinits
func init() {
	MustRegister(list())
}

func list() *Command {
	return &Command{
		Name:        CommandList,
		Description: "Lists commands",
		Help: `
The <info>{{ .Name }}</info> command lists all commands:
  <info>{{ .Bin }} {{ .Name }}</info>
You can also display the commands for a specific namespace:
  <info>{{ .Bin }} {{ .Name }} test</info>
You can also output the information in other formats by using the <comment>--format</comment> option:
  <info>{{ .Bin }} {{ .Name }} --format=xml</info>
`,
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			ns := in.Argument(ctx, "namespace").String()
			format := in.Option(ctx, helpOptFormat).String()

			des, err := descriptor.Find(format)
			if err != nil {
				return err
			}
			cmds := Commands()
			commands := descriptor.Commands{
				Namespace:  ns,
				Definition: Default(input.NewDefinition()),
			}
			groups := make(map[string]*descriptor.NSCommand)
			namespaces := make([]string, 0, len(cmds))
			empty := descriptor.NSCommand{}

			for _, name := range cmds {
				if ns != "" && !strings.HasPrefix(name, ns+":") {
					continue
				}

				cmd, _ := Find(name)
				if cmd.Hidden {
					continue
				}

				gn := strings.SplitN(name, ":", 2)
				if len(gn) != defaultLenNamespace {
					empty.Append(cmd.Name, cmd.Description)

					continue
				}

				if _, ok := groups[gn[0]]; !ok {
					groups[gn[0]] = &descriptor.NSCommand{
						Name: gn[0],
					}
					namespaces = append(namespaces, gn[0])
				}

				groups[gn[0]].Append(name, cmd.Description)
			}
			if len(empty.Commands) > 0 {
				commands.Commands = append(commands.Commands, empty)
			}

			for _, name := range namespaces {
				commands.Commands = append(commands.Commands, *groups[name])
			}

			if ns != "" && len(commands.Commands) == 0 {
				return fmt.Errorf("%w: namespace %s", ErrNotFound, ns)
			}

			return des.Commands(ctx, out, commands)
		},
		Configure: func(ctx context.Context, config *input.Definition) error {
			formats := descriptor.Descriptors()
			config.
				SetArguments(
					input.NewArgument("namespace", "The namespace name"),
				).
				SetOptions(
					input.NewOption(helpOptFormat, fmt.Sprintf("The output format (%s)", strings.Join(formats, ", ")),
						option.Required,
						option.Default(formats[0]),
						option.Valid(
							validator.NotBlank(0),
							validator.Enum(formats...),
						),
					),
				)

			return nil
		},
	}
}
