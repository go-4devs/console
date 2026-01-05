package command

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/output"
)

func Namespace() command.Command {
	return command.New("app:start", "example command in other namespace", NSExecute)
}

func NSExecute(ctx context.Context, _ config.Provider, out output.Output) error {
	out.Println(ctx, "example command in other namespace")

	return nil
}
