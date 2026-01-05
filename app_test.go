package console_test

import (
	"context"
	"os"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/command"
)

//nolint:lll
func ExampleNew_help() {
	ctx := context.Background()
	os.Args = []string{
		"bin/console",
		"test:command",
		"-h",
		"--no-ansi",
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
		"bin/console",
		"--no-ansi",
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
