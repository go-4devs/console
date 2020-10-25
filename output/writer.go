package output

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"gitoa.ru/go-4devs/console/output/verbosity"
)

const newline = "\n"

func Stderr() Output {
	return New(os.Stderr, FormatString)
}

func Stdout() Output {
	return New(os.Stdout, FormatString)
}

func Buffer(buf *bytes.Buffer) Output {
	return New(buf, FormatString)
}

func FormatString(_ verbosity.Verbosity, msg string, kv ...KeyValue) string {
	if len(kv) > 0 {
		nline := ""
		if msg[len(msg)-1:] == newline {
			nline = newline
		}

		return "msg=\"" + strings.TrimSpace(msg) + "\", " + KeyValues(kv).String() + nline
	}

	return msg
}

func New(w io.Writer, format func(verb verbosity.Verbosity, msg string, kv ...KeyValue) string) Output {
	return func(ctx context.Context, verb verbosity.Verbosity, msg string, kv ...KeyValue) (int, error) {
		return fmt.Fprint(w, format(verb, msg, kv...))
	}
}
