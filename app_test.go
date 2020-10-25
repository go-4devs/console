package console_test

import (
	"context"
	"os"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
)

//nolint: lll
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
	//   test_argument                  test argument
	//
	// Options:
	//       --duration[=DURATION]        test duration with default [default: 1s]
	//       --bool                       test bool option
	//       --string[=STRING]            array string (multiple values allowed)
	//   -q, --quiet                      Do not output any message
	//   -v, --verbose                    Increase the verbosity of messages: -v for info output, -vv for debug and -vvv for trace (multiple values allowed)
	//   -h, --help                       Display this help message
	//   -V, --version                    Display this application version
	//       --ansi                       Do not ask any interactive question
	//       --no-ansi                    Disable ANSI output
}

func ExampleNew_list() {
	ctx := context.Background()

	os.Args = []string{
		"bin/console",
		"--no-ansi",
	}

	console.New(console.WithExit(func(int) {})).
		Add(
			Command(),
			command.Hello(),
			command.Args(),
			command.Namespace(),
		).
		Execute(ctx)
	// Output:
	// Usage:
	//   command [options] [arguments]
	//
	// Options:
	//   -q, --quiet    Do not output any message
	//   -v, --verbose  Increase the verbosity of messages: -v for info output, -vv for debug and -vvv for trace (multiple values allowed)
	//   -h, --help     Display this help message
	//   -V, --version  Display this application version
	//       --ansi     Do not ask any interactive question
	//       --no-ansi  Disable ANSI output
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
