package dump

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/errs"
	"gitoa.ru/go-4devs/console/internal/registry"
	"gitoa.ru/go-4devs/console/output"
)

//go:generate go tool config config:generate

const NameRefernce = "config:dump-reference"

func Command() command.Command {
	return command.New(NameRefernce, "dump reference by command", RExecute, command.Configure(RConfigure))
}

func RExecute(ctx context.Context, in config.Provider, out output.Output) error {
	provs, ok := in.(config.Providers)
	if !ok {
		return fmt.Errorf("%w: expect %T got %T", errs.ErrWrongType, (config.Providers)(nil), in)
	}

	cfg := NewRConfigureConfig(in)

	cmd, err := registry.Find(cfg.CommandName(ctx))
	if err != nil {
		return fmt.Errorf("cmd:%w", err)
	}

	def := definition.New()
	if err := cmd.Configure(ctx, def); err != nil {
		return fmt.Errorf("configure:%w", err)
	}

	prov, err := provs.Provider(cfg.Format(ctx))
	if err != nil {
		return fmt.Errorf("prov:%w", errs.AlternativesError{Alt: provs.Names(), Err: err})
	}

	bind, ok := prov.(config.DumpProvider)
	if !ok {
		return fmt.Errorf("%w: expect config.DunpProvider got %T", errs.ErrWrongType, prov)
	}

	if err := bind.DumpReference(ctx, out, def); err != nil {
		return fmt.Errorf("dump:%w", err)
	}

	return nil
}

func RConfigure(_ context.Context, def config.Definition) error {
	def.Add(
		arg.String("command-name", "command name", option.Required),
		option.String("format", "format", option.Default(arg.Name)),
	)

	return nil
}
