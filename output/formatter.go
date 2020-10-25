package output

import (
	"context"

	"gitoa.ru/go-4devs/console/output/formatter"
	"gitoa.ru/go-4devs/console/output/verbosity"
)

func Format(out Output, format *formatter.Formatter) Output {
	return func(ctx context.Context, v verbosity.Verbosity, msg string, kv ...KeyValue) (int, error) {
		return out(ctx, v, format.Format(ctx, msg), kv...)
	}
}

func Ansi(out Output) Output {
	return Format(out, formatter.Ansi())
}

func None(out Output) Output {
	return Format(out, formatter.None())
}
