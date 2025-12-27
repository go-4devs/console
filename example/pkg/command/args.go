package command

import (
	"context"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/output"
)

func Args() *console.Command {
	return &console.Command{
		Name:        "fdevs:console:arg",
		Description: "Understanding how Console Arguments and Options Are Handled",
		Configure: func(_ context.Context, def config.Definition) error {
			def.Add(
				option.Bool("foo", "foo option", option.Short('f')),
				option.String("bar", "required bar option", option.Required, option.Short('b')),
				option.String("cat", "cat option", option.Short('c')),
				option.Time("time", "time example"),
				option.Time("hidden", "hidden time example", option.Hidden),
			)

			return nil
		},
		Execute: func(ctx context.Context, in config.Provider, out output.Output) error {
			out.Println(ctx, "foo: <info>", console.ReadValue(ctx, in, "foo").Bool(), "</info>")
			out.Println(ctx, "bar: <info>", console.ReadValue(ctx, in, "bar").String(), "</info>")
			out.Println(ctx, "cat: <info>", console.ReadValue(ctx, in, "cat").String(), "</info>")
			out.Println(ctx, "time: <info>", console.ReadValue(ctx, in, "time").Time().Format(time.RFC3339), "</info>")
			out.Println(ctx, "hidden: <info>", console.ReadValue(ctx, in, "hidden").Time().Format(time.RFC3339), "</info>")

			return nil
		},
	}
}
