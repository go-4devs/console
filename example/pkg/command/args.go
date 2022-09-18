package command

import (
	"context"
	"time"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/output"
)

func Args() *console.Command {
	return &console.Command{
		Name:        "fdevs:console:arg",
		Description: "Understanding how Console Arguments and Options Are Handled",
		Configure: func(ctx context.Context, def *input.Definition) error {
			def.SetOptions(
				option.Bool("foo", "foo option", option.Short('f')),
				option.String("bar", "required bar option", option.Required, option.Short('b')),
				option.String("cat", "cat option", option.Short('c')),
				option.Time("time", "time example"),
				option.Time("hidden", "hidden time example", option.Hidden),
			)

			return nil
		},
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			out.Println(ctx, "foo: <info>", in.Option(ctx, "foo").Bool(), "</info>")
			out.Println(ctx, "bar: <info>", in.Option(ctx, "bar").String(), "</info>")
			out.Println(ctx, "cat: <info>", in.Option(ctx, "cat").String(), "</info>")
			out.Println(ctx, "time: <info>", in.Option(ctx, "time").Time().Format(time.RFC3339), "</info>")
			out.Println(ctx, "hidden: <info>", in.Option(ctx, "hidden").Time().Format(time.RFC3339), "</info>")

			return nil
		},
	}
}
