package command

import (
	"context"

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
				option.Bool("foo", "foo option", option.Short("f")),
				input.NewOption("bar", "required bar option", option.Required, option.Short("b")),
				input.NewOption("cat", "cat option", option.Short("c")),
			)

			return nil
		},
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			out.Println(ctx, "foo: <info>", in.Option(ctx, "foo").Bool(), "</info>")
			out.Println(ctx, "bar: <info>", in.Option(ctx, "bar").String(), "</info>")
			out.Println(ctx, "cat: <info>", in.Option(ctx, "cat").String(), "</info>")

			return nil
		},
	}
}
