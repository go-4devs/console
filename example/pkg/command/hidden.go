package command

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/output"
)

func Hidden() command.Command {
	return command.New("fdevs:console:hidden", "hidden command exmale", HiddenExecute, command.Hidden)
}

func HiddenExecute(ctx context.Context, _ config.Provider, out output.Output) error {
	out.Println(ctx, "<info> call hidden command</info>")

	return nil
}
