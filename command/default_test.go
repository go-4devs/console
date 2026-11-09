package command_test

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/output"
)

func TestNoColorAnyValue(t *testing.T) {
	t.Setenv("NO_COLOR", "")

	ctx := context.Background()
	buf := &bytes.Buffer{}
	out := command.Ansi(ctx, &memory.Map{}, output.Buffer(buf))
	out.Println(ctx, "<info>x</info>")

	if strings.Contains(buf.String(), "\x1b") {
		t.Fatalf("NO_COLOR left color codes: %q", buf.String())
	}
}
