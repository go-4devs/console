package command

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	argument "gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/output"
)

func CreateUser(required bool) *console.Command {
	return &console.Command{
		Name:        "app:create-user",
		Description: "Creates a new user.",
		Help:        "This command allows you to create a user...",
		Configure: func(_ context.Context, cfg config.Definition) error {
			var opts []param.Option
			if required {
				opts = append(opts, option.Required)
			}

			cfg.
				Add(
					argument.String("username", "The username of the user.", option.Required),
					argument.String("password", "User password", opts...),
				)

			return nil
		},
		Execute: func(ctx context.Context, in config.Provider, out output.Output) error {
			// outputs a message followed by a "\n"
			out.Println(ctx, "User Creator")
			out.Println(ctx, "Username: ", console.ReadValue(ctx, in, "username").String())

			return nil
		},
	}
}
