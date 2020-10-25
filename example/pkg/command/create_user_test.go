package command_test

import (
	"bytes"
	"context"
	"testing"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
	"gitoa.ru/go-4devs/console/input/array"
	"gitoa.ru/go-4devs/console/output/writer"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	in := array.New(array.Argument("username", "andrey"))
	buf := bytes.Buffer{}
	out := writer.Buffer(&buf)

	err := console.Run(ctx, command.CreateUser(false), in, out)
	if err != nil {
		t.Fatalf("expect nil err, got :%s", err)
	}

	expect := "User Creator\nUsername:  andrey\n"

	if expect != buf.String() {
		t.Errorf("expect: %s, got:%s", expect, buf.String())
	}
}
