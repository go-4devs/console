package formatter_test

import (
	"context"
	"testing"

	"gitoa.ru/go-4devs/console/output/formatter"
)

func TestFormatter(t *testing.T) {
	ctx := context.Background()
	formatter := formatter.New()

	cases := map[string]string{
		"<info>info message</info>": "\x1b[32minfo message\x1b[39m",
		"<info><command></info>":    "\x1b[32m<command>\x1b[39m",
		"<html>...</html>":          "<html>...</html>",
	}

	for msg, ex := range cases {
		got := formatter.Format(ctx, msg)
		if ex != got {
			t.Errorf("ivalid expected:%#v, got: %#v", ex, got)
		}
	}
}
