package command

import (
	"context"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/definition/option"
	"github.com/go-4devs/config/provider/arg"
	"github.com/go-4devs/console"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/output"
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
