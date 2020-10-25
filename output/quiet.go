package output

import (
	"context"

	"gitoa.ru/go-4devs/console/output/verbosity"
)

func Quiet() Output {
	return func(context.Context, verbosity.Verbosity, string, ...KeyValue) (int, error) {
		return 0, nil
	}
}
