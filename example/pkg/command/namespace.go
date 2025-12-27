package command

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/output"
)

func Namespace() *console.Command {
	return &console.Command{
		Name:        "app:start",
		Description: "example command in other namespace",
		Execute: func(ctx context.Context, _ config.Provider, out output.Output) error {
			out.Println(ctx, "example command in other namespace")

			return nil
		},
	}
}
