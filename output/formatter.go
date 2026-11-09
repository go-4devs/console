package output

import (
	"context"

	"github.com/go-4devs/console/output/formatter"
	"github.com/go-4devs/console/output/label"
	"github.com/go-4devs/console/output/verbosity"
)

func Format(out Output, format *formatter.Formatter) Output {
	return func(ctx context.Context, v verbosity.Verbosity, msg string, kv ...label.KeyValue) (int, error) {
		return out(ctx, v, format.Format(ctx, msg), kv...)
	}
}

func Ansi(out Output) Output {
	return Format(out, formatter.Ansi())
}

func None(out Output) Output {
	return Format(out, formatter.None())
}
