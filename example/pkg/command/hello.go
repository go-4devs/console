package command

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/output"
)

func Hello() command.Command {
	return command.New(
		"fdevs:console:hello",
		"example hello command",
		HelloExecute,
		command.Configure(HelloConfigure),
	)
}

func HelloExecute(ctx context.Context, in config.Provider, out output.Output) error {
	name := console.ReadValue(ctx, in, "name").String()
	out.Println(ctx, "<error>Hello</error> <info>", name, "</info>")

	out.Info(ctx, "same trace info\n")
	out.Debug(ctx, "have some question?\n")
	out.Trace(ctx, "this message shows with -vvv\n")

	pass := console.ReadValue(ctx, in, "pass").String()
	out.Println(ctx, "hidden option pass <info>", pass, "</info>")

	return nil
}

func HelloConfigure(_ context.Context, def config.Definition) error {
	def.Add(
		arg.String("name", "Same name", arg.Default("World")),
		option.String("pass", "password", option.Hidden),
	)

	return nil
}
