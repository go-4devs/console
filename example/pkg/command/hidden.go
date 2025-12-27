package command

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/output"
)

func Hidden() *console.Command {
	return &console.Command{
		Name:        "fdevs:console:hidden",
		Description: "hidden command exmale",
		Hidden:      true,
		Execute: func(ctx context.Context, _ config.Provider, out output.Output) error {
			out.Println(ctx, "<info> call hidden command</info>")

			return nil
		},
	}
}
