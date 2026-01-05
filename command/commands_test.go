package command_test

import (
	"context"
	"testing"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/output"
)

func testEmtyExecute(context.Context, config.Provider, output.Output) error {
	return nil
}

func TestCommandsCommand(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		"fdevs:console:test": "fdevs:console:test",
		"fd:c:t":             "fdevs:console:test",
		"fd::t":              "fdevs:console:test",
		"f:c:t":              "fdevs:console:test",
		"f:c:a":              "fdevs:console:arg",
	}

	var commands command.Commands

	_ = commands.Add(
		command.New("fdevs:console:test", "fdevs console test", testEmtyExecute),
		command.New("fdevs:console:arg", "fdevs console arg", testEmtyExecute),
	)

	for name, ex := range cases {
		res, err := commands.Find(name)
		if err != nil {
			t.Errorf("%v expect <nil> err, got:%s", name, err)

			continue
		}

		if res.Name() != ex {
			t.Errorf("%v expect: %s, got: %s", name, ex, res)
		}
	}
}
