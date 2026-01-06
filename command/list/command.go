package list

import (
	"context"
	"fmt"
	"strings"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/errs"
	"gitoa.ru/go-4devs/console/internal/registry"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/descriptor"
	"gitoa.ru/go-4devs/console/setting"
)

//nolint:gochecknoinits
func init() {
	err := registry.Add(Command())
	if err != nil {
		panic(err)
	}
}

const (
	Name                = "list"
	ArgumentNamespace   = "namespace"
	OptionFormat        = "format"
	defaultLenNamespace = 2
)

func Command() command.Command {
	return command.New(
		Name,
		"Lists commands",
		Execite,
		command.Configure(Configure),
		command.Help(Help),
	)
}

func Configure(_ context.Context, cfg config.Definition) error {
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
}

//nolint:cyclop
func Execite(ctx context.Context, in config.Provider, out output.Output) error {
	opt := read{Provider: in}
	ns := opt.Value(ctx, ArgumentNamespace).String()
	format := opt.Value(ctx, OptionFormat).String()

	des, err := descriptor.Find(format)
	if err != nil {
		return fmt.Errorf("find descriptor[%v]: %w", format, err)
	}

	def := definition.New()
	command.Default(def)

	cmds := registry.Commands()
	commands := descriptor.Commands{
		Namespace: ns,
		Options:   def.With(param.New(descriptor.TxtStyle())),
	}
	groups := make(map[string]*descriptor.NSCommand)
	namespaces := make([]string, 0, len(cmds))
	empty := descriptor.NSCommand{}

	for _, name := range cmds {
		if ns != "" && !strings.HasPrefix(name, ns+":") {
			continue
		}

		cmd, _ := registry.Find(name)
		if setting.IsHidden(cmd) {
			continue
		}

		gn := strings.SplitN(name, ":", defaultLenNamespace)
		if len(gn) != defaultLenNamespace {
			empty.Append(cmd.Name(), setting.Description(cmd))

			continue
		}

		if _, ok := groups[gn[0]]; !ok {
			groups[gn[0]] = &descriptor.NSCommand{
				Name: gn[0],
			}

			namespaces = append(namespaces, gn[0])
		}

		groups[gn[0]].Append(name, setting.Description(cmd))
	}

	if len(empty.Commands) > 0 {
		commands.Commands = append(commands.Commands, empty)
	}

	for _, name := range namespaces {
		commands.Commands = append(commands.Commands, *groups[name])
	}

	if ns != "" && len(commands.Commands) == 0 {
		return fmt.Errorf("%w: namespace %s", errs.ErrNotFound, ns)
	}

	if err := des.Commands(ctx, out, commands); err != nil {
		return fmt.Errorf("descriptor:%w", err)
	}

	return nil
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
