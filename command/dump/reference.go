package dump

import (
	"context"
	"fmt"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/arg"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/errs"
	"github.com/go-4devs/console/internal/registry"
	"github.com/go-4devs/console/output"
)

//go:generate go tool config config:generate

const NameRefernce = "config:dump-reference"

func Command() command.Command {
	return command.New(NameRefernce, "dump reference by command", RExecute, command.Configure(RConfigure))
}

func RExecute(ctx context.Context, in config.Provider, out output.Output) error {
	provs, ok := in.(config.Providers)
	if !ok {
		return fmt.Errorf("%w: expect %T got %T", errs.ErrWrongType, config.Providers(nil), in)
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

	prov, err := provs.ByName(cfg.Format(ctx))
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
