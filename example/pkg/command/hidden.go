package command

import (
	"context"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/output"
)

func Hidden() *console.Command {
	return &console.Command{
		Name:        "fdevs:console:hidden",
		Description: "hidden command exmale",
		Hidden:      true,
		Execute: func(ctx context.Context, _ input.Input, out output.Output) error {
			out.Println(ctx, "<info> call hidden command</info>")

			return nil
		},
	}
}
