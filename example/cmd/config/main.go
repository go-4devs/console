package main

import (
	"context"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/provider/env"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/cfg"
)

const (
	Namespace = "fdevs"
	AppName   = "console"
)

// FDEVS_CONSOLE_CAT=env FDEVS_CONSOLE_HIDDEN=2022-09-18T23:07:49+03:00 go run cmd/config/main.go fdevs:console:arg -b tmp.
// FDEVS_CONSOLE_CAT=env go run cmd/config/main.go fdevs:console:arg --hidden=2022-09-18T23:07:49+03:00 -b tmp.
func main() {
	env := config.New(Namespace, AppName, []config.Provider{
		env.New(),
	})

	console.
		New(console.WithInput(
			input.Chain(
				input.NewArgs(0),
				cfg.New(env.Value),
			),
		)).
		Add(
			command.Long(),
			command.Args(),
		).
		Execute(context.Background())
}
