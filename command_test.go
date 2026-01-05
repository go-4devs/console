package console_test

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"testing"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/group"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/definition/proto"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/command"
	cerr "gitoa.ru/go-4devs/console/errors"
	"gitoa.ru/go-4devs/console/output"
)

//nolint:gochecknoinits
func init() {
	console.MustRegister(command.With(Command(), command.WithName("fdevs:console:test")))
	console.MustRegister(command.With(Command(), command.WithName("fdevs:console:arg")))
}

func Command() command.Command {
	return command.New("test:command", "test command", Execute, command.Configure(Configure))
}

func Execute(ctx context.Context, in config.Provider, out output.Output) error {
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
}

func Configure(_ context.Context, def config.Definition) error {
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
	if !errors.Is(err, cerr.ErrExecuteNil) {
		t.Fatalf("expected: %v, got: %v ", cerr.ErrExecuteNil, err)
	}
}
