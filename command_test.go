package console_test

import (
	"context"
	"errors"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/argument"
	"gitoa.ru/go-4devs/console/input/option"
	"gitoa.ru/go-4devs/console/output"
)

//nolint:gochecknoinits
func init() {
	console.MustRegister(Command().With(console.WithName("fdevs:console:test")))
	console.MustRegister(Command().With(console.WithName("fdevs:console:arg")))
}

func Command() *console.Command {
	return &console.Command{
		Name:        "test:command",
		Description: "test command",
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			out.Print(ctx,
				"test argument:", in.Argument(ctx, "test_argument").String(), "\n",
				"bool option:", in.Option(ctx, "bool").Bool(), "\n",
				"duration option with default:", in.Option(ctx, "duration").Duration(), "\n",
				"array string:[", strings.Join(in.Option(ctx, "string").Strings(), ","), "]\n",
			)

			return nil
		},
		Configure: func(ctx context.Context, def *input.Definition) error {
			def.
				SetArguments(
					argument.String("test_argument", "test argument"),
				).
				SetOptions(
					option.String("string", "array string", option.Array),
					option.Bool("bool", "test bool option"),
					option.Duration("duration", "test duration with default", option.Default(time.Second)),
					option.Time("hidden", "hidden time", option.Default(time.Second), option.Hidden),
				)

			return nil
		},
	}
}

func TestChainPrepare(t *testing.T) {
	t.Parallel()

	var cnt int32

	ctx := context.Background()
	def := input.NewDefinition()

	prepare := func(ctx context.Context, def *input.Definition, n console.Configure) error {
		atomic.AddInt32(&cnt, 1)

		return n(ctx, def)
	}
	configure := func(context.Context, *input.Definition) error {
		return nil
	}

	for i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		prepares := make([]console.Prepare, i)
		for p := 0; p < i; p++ {
			prepares[p] = prepare
		}

		cnt = 0
		chain := console.ChainPrepare(prepares...)

		if err := chain(ctx, def, configure); err != nil {
			t.Errorf("expected nil err, got: %s", err)
		}

		if cnt != int32(i) {
			t.Fatalf("expected: call prepare 1, got: %d ", cnt)
		}
	}
}

func TestChainHandle(t *testing.T) {
	t.Parallel()

	var cnt int32

	ctx := context.Background()
	in := &input.Array{
		Map: input.Map{},
	}
	out := output.Stdout()

	handle := func(ctx context.Context, in input.Input, out output.Output, next console.Action) error {
		atomic.AddInt32(&cnt, 1)

		return next(ctx, in, out)
	}
	action := func(context.Context, input.Input, output.Output) error {
		return nil
	}

	for i := range []int{0, 1, 2, 30, 40, 50} {
		handles := make([]console.Handle, i)
		for p := 0; p < i; p++ {
			handles[p] = handle
		}

		cnt = 0
		chain := console.ChainHandle(handles...)

		if err := chain(ctx, in, out, action); err != nil {
			t.Errorf("expected nil err, got: %s", err)
		}

		if cnt != int32(i) {
			t.Fatalf("expected: call prepare 1, got: %d ", cnt)
		}
	}
}

func TestRunEmptyExecute(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	empty := console.Command{
		Name: "empty",
	}
	in := &input.Array{
		Map: input.Map{},
	}
	out := output.Stdout()

	err := empty.Run(ctx, in, out)
	if !errors.Is(err, console.ErrExecuteNil) {
		t.Fatalf("expected: %v, got: %v ", console.ErrExecuteNil, err)
	}
}
