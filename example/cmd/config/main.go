package main

import (
	"context"

	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/provider/chain"
	"gitoa.ru/go-4devs/config/provider/env"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
)

const (
	Namespace = "fdevs"
	AppName   = "console"
)

// FDEVS_CONSOLE_CAT=env FDEVS_CONSOLE_HIDDEN=2022-09-18T23:07:49+03:00 go run cmd/config/main.go fdevs:console:arg -b tmp.
// FDEVS_CONSOLE_CAT=env go run cmd/config/main.go fdevs:console:arg --hidden=2022-09-18T23:07:49+03:00 -b tmp.
func main() {
	console.
		New(console.WithInput(
			chain.New(
				arg.New(arg.WithSkip(0)),
				env.New(Namespace, AppName),
				&memory.Default{},
			),
		)).
		Add(
			command.Long(),
			command.Args(),
		).
		Execute(context.Background())
}
