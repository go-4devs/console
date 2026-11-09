package command

import (
	"context"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/param"
	argument "github.com/go-4devs/config/provider/arg"
	"github.com/go-4devs/console"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/output"
	"github.com/go-4devs/console/setting"
)

func CreateUser(required bool) command.Command {
	return command.New(
		"app:create-user",
		"Creates a new user.",
		UserExecute,
		command.Configure(UserConfigure(required)),
		command.Help(func(setting.HData) (string, error) {
			return "This command allows you to create a user...", nil
		}),
	)
}

func UserConfigure(required bool) func(_ context.Context, cfg config.Definition) error {
	return func(_ context.Context, cfg config.Definition) error {
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
	}
}

func UserExecute(ctx context.Context, in config.Provider, out output.Output) error {
	// outputs a message followed by a "\n"
	out.Println(ctx, "User Creator")
	out.Println(ctx, "Username: ", console.ReadValue(ctx, in, "username").String())

	return nil
}
