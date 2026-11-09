package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/go-4devs/console"
	"github.com/go-4devs/console/example/pkg/command"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan os.Signal, 1)
	defer close(ch)

	signal.Notify(ch, os.Interrupt)

	go func() {
		<-ch
		cancel()
	}()

	console.Execute(ctx, command.Long())
}
