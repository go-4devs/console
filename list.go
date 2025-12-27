package console

import (
	"context"
	"fmt"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/descriptor"
)

const defaultLenNamespace = 2

//nolint:gochecknoinits
func init() {
	MustRegister(list())
}

const (
	ArgumentNamespace = "namespace"
)

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
		Execute: executeList,
		Configure: func(_ context.Context, cfg config.Definition) error {
			formats := descriptor.Descriptors()
			cfg.
				Add(
					arg.String(ArgumentNamespace, "The namespace name"),
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

//nolint:cyclop
func executeList(ctx context.Context, in config.Provider, out output.Output) error {
	ns := ReadValue(ctx, in, ArgumentNamespace).String()
	format := ReadValue(ctx, in, OptionFormat).String()

	des, err := descriptor.Find(format)
	if err != nil {
		return fmt.Errorf("find descriptor[%v]: %w", format, err)
	}

	cmds := Commands()
	commands := descriptor.Commands{
		Namespace:  ns,
		Definition: descriptor.NewDefinition(config.NewVars(definition.New(Default()...).Options()...).Variables()),
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

		gn := strings.SplitN(name, ":", defaultLenNamespace)
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

	if err := des.Commands(ctx, out, commands); err != nil {
		return fmt.Errorf("descriptor:%w", err)
	}

	return nil
}
