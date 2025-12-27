package console_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console"
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
		Execute: func(ctx context.Context, in config.Provider, out output.Output) error {
			var astr []string
			if aerr := console.ReadValue(ctx, in, "string").Unmarshal(&astr); aerr != nil && !errors.Is(aerr, config.ErrNotFound) {
				return fmt.Errorf("unmarshal string:%w", aerr)
			}

			out.Print(ctx,
				"test argument:", console.ReadValue(ctx, in, "test_argument").String(), "\n",
				"bool option:", console.ReadValue(ctx, in, "bool").Bool(), "\n",
				"duration option with default:", console.ReadValue(ctx, in, "duration").Duration(), "\n",
				"array string:[", strings.Join(astr, ","), "]\n",
				"group string:", console.ReadValue(ctx, in, "group", "test", "string").String(), "\n",
				"log http service:", console.ReadValue(ctx, in, "log", "http", "level").String(), "\n",
			)

			return nil
		},
		Configure: func(_ context.Context, def config.Definition) error {
			def.
				Add(
					group.New("group", "group example",
						option.Bool("bool", "bool"),
						group.New("test", "test", option.String("string", "test group string", option.Default("group string default value"))),
					),
					group.New("log", "log",
						proto.New("service", "service level",
							option.String("level", "service level", option.Default("debug")),
						),
					),
					arg.String("test_argument", "test argument"),
					option.String("string", "array string", option.Slice),
					option.Bool("bool", "test bool option"),
					option.Duration("duration", "test duration with default", option.Default(value.New(time.Second))),
					option.Time("hidden", "hidden time", option.Default(value.New(time.Second)), option.Hidden),
				)

			return nil
		},
	}
}

func TestChainPrepare(t *testing.T) {
	t.Parallel()

	var cnt int64

	ctx := context.Background()
	def := definition.New()

	prepare := func(ctx context.Context, def config.Definition, n console.Configure) error {
		atomic.AddInt64(&cnt, 1)

		return n(ctx, def)
	}
	configure := func(context.Context, config.Definition) error {
		return nil
	}

	for i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		prepares := make([]console.Prepare, i)
		for p := range i {
			prepares[p] = prepare
		}

		cnt = 0
		chain := console.ChainPrepare(prepares...)

		err := chain(ctx, def, configure)
		if err != nil {
			t.Errorf("expected nil err, got: %s", err)
		}

		if cnt != int64(i) {
			t.Fatalf("expected: call prepare 1, got: %d ", cnt)
		}
	}
}

func TestChainHandle(t *testing.T) {
	t.Parallel()

	var cnt int64

	ctx := context.Background()
	in := &memory.Map{}
	out := output.Stdout()

	handle := func(ctx context.Context, in config.Provider, out output.Output, next console.Action) error {
		atomic.AddInt64(&cnt, 1)

		return next(ctx, in, out)
	}
	action := func(context.Context, config.Provider, output.Output) error {
		return nil
	}

	for i := range []int{0, 1, 2, 30, 40, 50} {
		handles := make([]console.Handle, i)
		for p := range i {
			handles[p] = handle
		}

		cnt = 0
		chain := console.ChainHandle(handles...)

		err := chain(ctx, in, out, action)
		if err != nil {
			t.Errorf("expected nil err, got: %s", err)
		}

		if cnt != int64(i) {
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
	in := &memory.Map{}
	out := output.Stdout()

	err := empty.Run(ctx, in, out)
	if !errors.Is(err, console.ErrExecuteNil) {
		t.Fatalf("expected: %v, got: %v ", console.ErrExecuteNil, err)
	}
}
