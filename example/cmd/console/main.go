package main

import (
	"context"

	"github.com/go-4devs/console"
	"github.com/go-4devs/console/command/dump"
	"github.com/go-4devs/console/example/pkg/command"
)

func main() {
	console.
		New().
		Add(
			dump.Command(),
			command.Hello(),
			command.Args(),
			command.Hidden(),
			command.Namespace(),
			command.CreateUser(false),
		).
		Execute(context.Background())
}
