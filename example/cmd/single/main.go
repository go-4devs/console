package main

import (
	"context"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
)

func main() {
	console.Execute(context.Background(), command.Hello())
}
