package command

import (
	"context"

	"github.com/go-4devs/config"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/output"
)

func Namespace() command.Command {
	return command.New("app:start", "example command in other namespace", NSExecute)
}

func NSExecute(ctx context.Context, _ config.Provider, out output.Output) error {
	out.Println(ctx, "example command in other namespace")

	return nil
}
