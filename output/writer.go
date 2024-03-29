package output

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"strings"

	"gitoa.ru/go-4devs/console/output/label"
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

func FormatString(_ verbosity.Verbosity, msg string, kv ...label.KeyValue) string {
	if len(kv) > 0 {
		nline := ""
		if msg[len(msg)-1:] == newline {
			nline = newline
		}

		return "msg=\"" + strings.TrimSpace(msg) + "\", " + label.KeyValues(kv).String() + nline
	}

	return msg
}

func New(w io.Writer, format func(verb verbosity.Verbosity, msg string, kv ...label.KeyValue) string) Output {
	return func(ctx context.Context, verb verbosity.Verbosity, msg string, kv ...label.KeyValue) (int, error) {
		out, err := fmt.Fprint(w, format(verb, msg, kv...))
		if err != nil {
			return 0, fmt.Errorf("writer fprint:%w", err)
		}

		return out, nil
	}
}
