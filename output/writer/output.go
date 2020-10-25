package writer

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"gitoa.ru/go-4devs/console/output"
)

func Stderr() output.Output {
	return New(os.Stderr, String)
}

func Stdout() output.Output {
	return New(os.Stdout, String)
}

func Buffer(buf *bytes.Buffer) output.Output {
	return New(buf, String)
}

func String(_ output.Verbosity, msg string, kv ...output.KeyValue) string {
	if len(kv) > 0 {
		newline := ""
		if msg[len(msg)-1:] == "\n" {
			newline = "\n"
		}

		return "msg=\"" + strings.TrimSpace(msg) + "\", " + output.KeyValues(kv).String() + newline

	}

	return msg
}

func New(w io.Writer, format func(verb output.Verbosity, msg string, kv ...output.KeyValue) string) output.Output {
	return func(ctx context.Context, verb output.Verbosity, msg string, kv ...output.KeyValue) (int, error) {
		return fmt.Fprint(w, format(verb, msg, kv...))
	}
}
