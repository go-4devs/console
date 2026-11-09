package main

import (
	"context"

	"github.com/go-4devs/console"
	"github.com/go-4devs/console/example/pkg/command"
)

func main() {
	console.Execute(context.Background(), command.Hello())
}
