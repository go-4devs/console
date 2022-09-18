package formatter_test

import (
	"context"
	"testing"

	"gitoa.ru/go-4devs/console/output/formatter"
)

func TestNone(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	none := formatter.None()

	cases := map[string]string{
		"<info>message info</info>":    "message info",
		"<error>message error</error>": "message error",
		"<comment><scheme></comment>":  "<scheme>",
		"<body>body</body>":            "<body>body</body>",
	}

	for msg, ex := range cases {
		got := none.Format(ctx, msg)
		if ex != got {
			t.Errorf("expect:%#v, got:%#v", ex, got)
		}
	}
}
