package output_test

import (
	"bytes"
	"context"
	"testing"

	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/label"
)

func TestNew(t *testing.T) {
	ctx := context.Background()
	buf := bytes.Buffer{}
	wr := output.New(&buf, output.FormatString)

	cases := map[string]struct {
		ex string
		kv []label.KeyValue
	}{
		"message": {
			ex: "message",
		},
		"msg with kv": {
			ex: "msg=\"msg with kv\", string key=\"string value\", bool key=\"false\", int key=\"42\"",
			kv: []label.KeyValue{
				label.String("string key", "string value"),
				label.Bool("bool key", false),
				label.Int("int key", 42),
			},
		},
		"msg with newline \n": {
			ex: "msg=\"msg with newline\", int=\"42\"\n",
			kv: []label.KeyValue{
				label.Int("int", 42),
			},
		},
	}

	for msg, data := range cases {
		wr.InfoKV(ctx, msg, data.kv...)

		if data.ex != buf.String() {
			t.Errorf("message not equals expext:%s, got:%s", data.ex, buf.String())
		}

		buf.Reset()
	}
}
