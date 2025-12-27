package command_test

import (
	"bytes"
	"context"
	"testing"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
	"gitoa.ru/go-4devs/console/input/key"
	"gitoa.ru/go-4devs/console/input/provider/memory"
	"gitoa.ru/go-4devs/console/output"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	buf := bytes.Buffer{}
	out := output.Buffer(&buf)
	in := &memory.Map{}
	in.SetOption(key.Dash("username"), "andrey")

	err := console.Run(ctx, command.CreateUser(false), in, out)
	if err != nil {
		t.Fatalf("expect nil err, got :%s", err)
	}

	expect := "User Creator\nUsername:  andrey\n"

	if expect != buf.String() {
		t.Errorf("expect: %s, got:%s", expect, buf.String())
	}
}
