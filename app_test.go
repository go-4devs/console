package console_test

import (
	"bytes"
	"context"
	"os"
	"strings"
	"testing"

	"github.com/go-4devs/console"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/output"
)

const (
	binConsole = "bin/console"
	flagHelp   = "-h"
	flagNoANSI = "--no-ansi"
)

func TestShortFlagIsNotCommandName(t *testing.T) { //nolint:paralleltest // os.Args is process-wide
	old := os.Args

	t.Cleanup(func() { os.Args = old })

	os.Args = []string{binConsole, flagHelp, flagNoANSI}

	var buf bytes.Buffer

	code := 0

	console.New(
		console.WithExit(func(c int) { code = c }),
		console.WithOutput(output.Buffer(&buf)),
	).Execute(context.Background())

	if code != 0 {
		t.Fatalf("exit %d, output: %s", code, buf.String())
	}

	if strings.Contains(buf.String(), "not found") {
		t.Fatalf("-h was used as a command name: %s", buf.String())
	}

	if !strings.Contains(buf.String(), "Lists commands") {
		t.Fatalf("expected list help, got: %s", buf.String())
	}
}

func TestBadArgExitsNonZero(t *testing.T) { //nolint:paralleltest // os.Args is process-wide
	old := os.Args

	t.Cleanup(func() { os.Args = old })

	os.Args = []string{binConsole, "--bogus"}

	var buf bytes.Buffer

	code := 0

	console.New(
		console.WithExit(func(c int) { code = c }),
		console.WithOutput(output.Buffer(&buf)),
	).Execute(context.Background())

	if code == 0 {
		t.Fatalf("exit 0, output: %s", buf.String())
	}

	if strings.Count(buf.String(), "bogus") != 1 {
		t.Fatalf("error printed %d times: %s", strings.Count(buf.String(), "bogus"), buf.String())
	}

	if !strings.Contains(buf.String(), "Lists commands") {
		t.Fatalf("expected help after the error, got: %s", buf.String())
	}
}

//nolint:lll
func ExampleNew_help() {
	ctx := context.Background()
	os.Args = []string{
		binConsole,
		"test:command",
		flagHelp,
		flagNoANSI,
	}

	console.New(console.WithExit(func(int) {})).
		Add(
			Command(),
		).
		Execute(ctx)
	// Output:
	// Description:
	//   test command
	//
	// Usage:
	//   test:command [options] [--] [<test_argument>]
	//
	// Arguments:
	//   test_argument                                      test argument
	//
	// Options:
	//       --string[=STRING]                              array string  (multiple values allowed)
	//       --group-bool                                   bool
	//       --group-test-string[=GROUP-TEST-STRING]        test group string [default:group string default value]
	//       --log-{service}-level[=LOG-{SERVICE}-LEVEL]    service level [default:debug]
	//       --bool                                         test bool option
	//       --duration[=DURATION]                          test duration with default
	//       --ansi                                         Do not ask any interactive question
	//   -V, --version                                      Display this application version
	//   -h, --help                                         Display this help message
	//   -v, --verbose                                      Increase the verbosity of messages: -v for info output, -vv for debug and -vvv for trace  (multiple values allowed)
	//   -q, --quiet                                        Do not output any message
	//       --no-ansi                                      Disable ANSI output
}

func ExampleNew_list() {
	ctx := context.Background()

	os.Args = []string{
		binConsole,
		flagNoANSI,
	}

	console.New(
		console.WithExit(func(int) {}),
		console.WithReplaceCommand,
	).
		Add(
			Command(),
			command.New("fdevs:console:arg", "Understanding how Console Arguments and Options Are Handled", Execute),
			command.New("fdevs:console:hello", "example hello command", Execute),
			command.New("app:start", "example command in other namespace", Execute),
		).
		Execute(ctx)
	// Output:
	// Usage:
	//   command [options] [arguments]
	//
	// Options:
	//       --no-ansi  Disable ANSI output
	//       --ansi     Do not ask any interactive question
	//   -V, --version  Display this application version
	//   -h, --help     Display this help message
	//   -v, --verbose  Increase the verbosity of messages: -v for info output, -vv for debug and -vvv for trace  (multiple values allowed)
	//   -q, --quiet    Do not output any message
	//
	// Available commands:
	//   help                 Displays help for a command
	//   list                 Lists commands
	// app
	//   app:start            example command in other namespace
	// fdevs
	//   fdevs:console:arg    Understanding how Console Arguments and Options Are Handled
	//   fdevs:console:hello  example hello command
	//   fdevs:console:test   test command
	// test
	//   test:command         test command
}
