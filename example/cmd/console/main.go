package main

import (
	"context"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
)

func main() {
	console.
		New().
		Add(
			command.Hello(),
			command.Args(),
			command.Hidden(),
			command.Namespace(),
			command.CreateUser(false),
		).
		Execute(context.Background())
}
