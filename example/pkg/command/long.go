package command

import (
	"context"
	"time"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/validator"
)

const defaultTimeout = time.Second * 30

// Long example of a command that takes a long time to run.
func Long() *console.Command {
	return &console.Command{
		Name: "fdevs:command:long",
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			timeout := in.Option(ctx, "timeout").Duration()
			timer := time.NewTimer(timeout)
			ticker := time.NewTicker(time.Second)
			defer ticker.Stop()
			for {
				select {
				case t := <-ticker.C:
					out.Println(ctx, "ticker: <info>", t, "</info>")
				case <-timer.C:
					out.Println(ctx, "<error>stop timer</error>")

					return nil
				case <-ctx.Done():
					out.Println(ctx, "<info>cancel context</info>")

					return nil
				}
			}
		},
		Configure: func(ctx context.Context, def *input.Definition) error {
			def.SetOptions(option.Duration("timeout", "set duration run command",
				option.Default(defaultTimeout),
				option.Short("t"),
				option.Valid(validator.NotBlank(input.ValueDuration)),
			))

			return nil
		},
	}
}
