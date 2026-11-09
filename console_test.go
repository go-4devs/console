package console_test

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/go-4devs/config"
	"github.com/go-4devs/config/provider/chain"
	"github.com/go-4devs/config/provider/memory"
	"github.com/go-4devs/config/value"
	"github.com/go-4devs/console"
	"github.com/go-4devs/console/command"
	"github.com/go-4devs/console/output"
)

type errWriter struct{}

func (errWriter) Write([]byte) (int, error) {
	return 0, io.ErrClosedPipe
}

func TestRunWriteError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	cmd := command.New("write", "w", func(ctx context.Context, _ config.Provider, out output.Output) error {
		out.Println(ctx, "hi")

		return nil
	})

	err := console.Run(ctx, cmd, chain.New(&memory.Map{}, &memory.Default{}), output.New(errWriter{}, output.FormatString))
	if !errors.Is(err, io.ErrClosedPipe) {
		t.Fatalf("got %v", err)
	}
}

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
