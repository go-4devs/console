package command_test

import (
	"context"
	"sync/atomic"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/output"
)

func TestChainPrepare(t *testing.T) {
	t.Parallel()

	var cnt int64

	ctx := context.Background()
	def := definition.New()

	prepare := func(ctx context.Context, def config.Definition, n command.ConfigureFn) error {
		atomic.AddInt64(&cnt, 1)

		return n(ctx, def)
	}
	configure := func(context.Context, config.Definition) error {
		return nil
	}

	for i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10} {
		prepares := make([]command.PrepareFn, i)
		for p := range i {
			prepares[p] = prepare
		}

		cnt = 0
		chain := command.ChainPrepare(prepares...)

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

	handle := func(ctx context.Context, in config.Provider, out output.Output, next command.ExecuteFn) error {
		atomic.AddInt64(&cnt, 1)

		return next(ctx, in, out)
	}
	action := func(context.Context, config.Provider, output.Output) error {
		return nil
	}

	for i := range []int{0, 1, 2, 30, 40, 50} {
		handles := make([]command.HandleFn, i)
		for p := range i {
			handles[p] = handle
		}

		cnt = 0
		chain := command.ChainHandle(handles...)

		err := chain(ctx, in, out, action)
		if err != nil {
			t.Errorf("expected nil err, got: %s", err)
		}

		if cnt != int64(i) {
			t.Fatalf("expected: call prepare 1, got: %d ", cnt)
		}
	}
}
