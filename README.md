# Console


[![Build Status](https://drone.gitoa.ru/api/badges/go-4devs/console/status.svg)](https://drone.gitoa.ru/go-4devs/console)
[![Go Report Card](https://goreportcard.com/badge/gitoa.ru/go-4devs/console)](https://goreportcard.com/report/gitoa.ru/go-4devs/console)
[![GoDoc](https://godoc.org/gitoa.ru/go-4devs/console?status.svg)](http://godoc.org/gitoa.ru/go-4devs/console)


## Creating a Command

Commands are defined in struct extending `pkg/command/create_user.go`. For example, you may want a command to create a user:

```go
package command

import (
	"context"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/output"
)

func Createuser() *console.Command {
	return &console.Command{
		Name: "app:create-user",
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			return nil
		},
	}
}
```
## Configure command

```go
func Createuser() *console.Command {
	return &console.Command{
        //...
		Description: "Creates a new user.",
		Help:        "This command allows you to create a user...",
	}
}
```


## Add arguments

```go
func Createuser(required bool) *console.Command {
	return &console.Command{
        //....
		Configure: func(ctx context.Context, cfg *input.Definition) error {
			var opts []func(*input.Argument)
			if required {
				opts = append(opts, argument.Required)
			}
			cfg.SetArgument("password", "User password", opts...)

			return nil
		},
	}
}
```

## Registering the Command

`cmd/console/main.go`

```go
package main

import (
	"context"

	"gitoa.ru/go-4devs/console"
	"pkg/command"
)

func main() {
	console.
		New().
		Add(
			command.Createuser(false),
		).
		Execute(context.Background())
}
```

## Executing the Command

build command `go build -o bin/console cmd/console/main.go`
run command `bin/console app:create-user``

## Console Output

The Execute field has access to the output stream to write messages to the console:
```go
func Createuser(required bool) *console.Command {
	return &console.Command{
        // ....
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			// outputs a message followed by a "\n"
			out.Println(ctx, "User Creator")
			out.Println(ctx, "Whoa!")

			// outputs a message without adding a "\n" at the end of the line
			out.Print(ctx, "You are about to ", "create a user.")

			return nil
		},
	}
}
```

Now, try build and executing the command:

```bash
bin/console app:create-user
User Creator
Whoa!
You are about to create a user.
```

## Console Input

Use input options or arguments to pass information to the command:

```go
func CreateUser(required bool) *console.Command {
	return &console.Command{
		Configure: func(ctx context.Context, cfg *input.Definition) error {
			var opts []func(*input.Argument)
			if required {
				opts = append(opts, argument.Required)
			}
			cfg.
				SetArgument("username", "The username of the user.", argument.Required).
				SetArgument("password", "User password", opts...)

			return nil
		},
		Execute: func(ctx context.Context, in input.Input, out output.Output) error {
			// outputs a message followed by a "\n"
			out.Println(ctx, "User Creator")
			out.Println(ctx, "Username: ", in.Argument(ctx, "username").String())

			return nil
		},
	}
}
```

Now, you can pass the username to the command:

```bash
bin/console app:create-user AwesomeUsername
User Creator
Username: AwesomeUsername
```

## Testing Commands

```go
package command_test

import (
	"bytes"
	"context"
	"testing"

	"gitoa.ru/go-4devs/console"
	"gitoa.ru/go-4devs/console/example/pkg/command"
	"gitoa.ru/go-4devs/console/input/array"
	"gitoa.ru/go-4devs/console/output/writer"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	in := array.New(array.Argument("username", "andrey"))
	buf := bytes.Buffer{}
	out := writer.Buffer(&buf)

	console.Run(ctx, command.CreateUser(false), in, out)

	expect := `User Creator
Username: andrey
`

	if expect != buf.String() {
		t.Errorf("expect: %s, got:%s", expect, buf.String())
	}
}
```
