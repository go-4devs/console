package command

import (
	"context"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/output"
)

func CreateUser(required bool) *console.Command {
	return &console.Command{
		Name:        "app:create-user",
		Description: "Creates a new user.",
		Help:        "This command allows you to create a user...",
		Configure: func(ctx context.Context, cfg *input.Definition) error {
			var opts []func(*input.Argument)
			if required {
				opts = append(opts, argument.Required)
			}
			cfg.
				SetArgument("username", "The username of the user.", argument.Required).
				SetArgument("password", "User password", opts...)

			return nil
		},
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			// outputs a message followed by a "\n"
			out.Println(ctx, "User Creator")
			out.Println(ctx, "Username: ", in.Argument(ctx, "username").String())

			return nil
		},
	}
}
