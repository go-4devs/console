package console_test

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"gitoa.ru/go-4devs/config/provider/chain"
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/output"
)

func ExampleRun() {
	cmd := Command()
	ctx := context.Background()
	out := output.Stdout()
	in := chain.New(&memory.Map{}, &memory.Default{})

	err := console.Run(ctx, cmd, in, out)
	fmt.Println("err:", err)
	// Output:
	// test argument:
	// bool option:false
	// duration option with default:1s
	// array string:[]
	// group string:group string default value
	// log http service:debug
	// err: <nil>
}

func ExampleExecute() {
	cmd := Command()
	ctx := context.Background()
	in := &memory.Map{}

	jb, err := json.Marshal([]string{"same value", "other value"})
	if err != nil {
		log.Print(err)
	}

	// Run command: ./bin "argument value" -b --string="same value" --string="other value"
	in.SetOption(value.New(true), "bool")
	in.SetOption(value.JBytes(jb), "string")
	in.SetOption(value.New("argument value"), "test_argument")
	in.SetOption(value.New("error"), "log", "http", "level")

	console.Execute(ctx, cmd, console.WithInput(chain.New(in, &memory.Default{})), console.WithExit(func(int) {}))
	// Output:
	// test argument:argument value
	// bool option:true
	// duration option with default:1s
	// array string:[same value,other value]
	// group string:group string default value
	// log http service:error
}
