package output

import (
	"context"

	"github.com/go-4devs/console/output/label"
	"github.com/go-4devs/console/output/verbosity"
)

func Quiet() Output {
	return func(context.Context, verbosity.Verbosity, string, ...label.KeyValue) (int, error) {
		return 0, nil
	}
}
