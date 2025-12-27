package command

import (
	"context"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/input/key"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/output"
)

func Hello() *console.Command {
	return &console.Command{
		Name:        "fdevs:console:hello",
		Description: "example hello command",
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			name := in.Value(ctx, key.Dash("name")).String()
			out.Println(ctx, "<error>Hello</error> <info>", name, "</info>")

			out.Info(ctx, "same trace info\n")
			out.Debug(ctx, "have some question?\n")
			out.Trace(ctx, "this message shows with -vvv\n")

			pass := in.Value(ctx, key.Dash("pass")).String()
			out.Println(ctx, "hidden option pass <info>", pass, "</info>")

			return nil
		},
		Configure: func(_ context.Context, def *input.Definition) error {
			def.SetOptions(
				argument.String("name", "Same name", argument.Default("World")),
				option.String("pass", "password", option.Hidden),
			)

			return nil
		},
	}
}
