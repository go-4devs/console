package verbosity

import (
	"context"

	"gitoa.ru/go-4devs/console/output"
)

func Verb(out output.Output, verb output.Verbosity) output.Output {
	return func(ctx context.Context, v output.Verbosity, msg string, kv ...output.KeyValue) (int, error) {
		if verb >= v {
			return out(ctx, v, msg, kv...)
		}

		return 0, nil
	}
}

func Quiet() output.Output {
	return func(context.Context, output.Verbosity, string, ...output.KeyValue) (int, error) {
		return 0, nil
	}
}
