package wrap

import (
	"context"

	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/output/formatter"
)

func Format(out output.Output, format *formatter.Formatter) output.Output {
	return func(ctx context.Context, v output.Verbosity, msg string, kv ...output.KeyValue) (int, error) {
		return out(ctx, v, format.Format(ctx, msg), kv...)
	}
}

func Ansi(out output.Output) output.Output {
	return Format(out, formatter.Ansi())
}

func None(out output.Output) output.Output {
	return Format(out, formatter.None())
}
