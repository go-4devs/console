package console_test

import (
	"context"
	"fmt"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/output"
)

func ExampleRun() {
	cmd := Command()
	ctx := context.Background()
	out := output.Stdout()
	in := &input.Array{}

	err := console.Run(ctx, cmd, in, out)
	fmt.Println("err:", err)
	// Output:
	// test argument:
	// bool option:false
	// duration option with default:1s
	// array string:[]
	// err: <nil>
}

func ExampleExecute() {
	cmd := Command()
	ctx := context.Background()
	in := &input.Array{}

	// Run command: ./bin "argument value" -b --string="same value" --string="other value"
	in.SetOption("bool", value.New(true))
	in.SetOption("string", value.New([]string{"same value", "other value"}))
	in.SetArgument("test_argument", value.New("argument value"))

	console.Execute(ctx, cmd, console.WithInput(in), console.WithExit(func(int) {}))
	// Output:
	// test argument:argument value
	// bool option:true
	// duration option with default:1s
	// array string:[same value,other value]
}
