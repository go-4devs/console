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

const newline = "\n"

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
		nline := ""
		if msg[len(msg)-1:] == newline {
			nline = newline
		}

		return "msg=\"" + strings.TrimSpace(msg) + "\", " + output.KeyValues(kv).String() + nline
	}

	return msg
}

func New(w io.Writer, format func(verb output.Verbosity, msg string, kv ...output.KeyValue) string) output.Output {
	return func(ctx context.Context, verb output.Verbosity, msg string, kv ...output.KeyValue) (int, error) {
		return fmt.Fprint(w, format(verb, msg, kv...))
	}
}
