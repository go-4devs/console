package command_test

import (
	"context"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/output"
)

const testCommand = "fdevs:console:test"

func testEmtyExecute(context.Context, config.Provider, output.Output) error {
	return nil
}

func TestCommandsCommand(t *testing.T) {
	t.Parallel()

	cases := map[string]string{
		testCommand: testCommand,
		"fd:c:t":    testCommand,
		"fd::t":     testCommand,
		"f:c:t":     testCommand,
		"f:c:a":     "fdevs:console:arg",
	}

	var commands command.Commands

	_ = commands.Add(
		command.New(testCommand, "fdevs console test", testEmtyExecute),
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
