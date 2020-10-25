package output

import (
	"context"

	"gitoa.ru/go-4devs/console/output/verbosity"
)

func Verbosity(out Output, verb verbosity.Verbosity) Output {
	return func(ctx context.Context, v verbosity.Verbosity, msg string, kv ...KeyValue) (int, error) {
		if verb >= v {
			return out(ctx, v, msg, kv...)
		}

		return 0, nil
	}
}
