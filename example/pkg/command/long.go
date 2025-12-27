package command

import (
	"context"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/validator"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/output"
)

const defaultTimeout = time.Second * 30

// Long example of a command that takes a long time to run.
func Long() *console.Command {
	return &console.Command{
		Name: "fdevs:command:long",
		Execute: func(ctx context.Context, in config.Provider, out output.Output) error {
			timeout := console.ReadValue(ctx, in, "timeout").Duration()
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
		Configure: func(_ context.Context, def config.Definition) error {
			def.Add(option.Duration("timeout", "set duration run command",
				option.Default(value.New(defaultTimeout)),
				option.Short('t'),
				validator.Valid(validator.NotBlank),
			))

			return nil
		},
	}
}
