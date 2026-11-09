package command_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/console"
	"github.com/go-4devs/console/example/pkg/command"
	"github.com/go-4devs/console/output"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	buf := bytes.Buffer{}
	out := output.Buffer(&buf)
	in := &memory.Map{}
	in.SetOption("andrey", "username")

	err := console.Run(ctx, command.CreateUser(false), in, out)
	if err != nil {
		t.Fatalf("expect nil err, got :%s", err)
	}

	expect := "User Creator\nUsername:  andrey\n"

	if expect != buf.String() {
		t.Errorf("expect: %s, got:%s", expect, buf.String())
	}
}
