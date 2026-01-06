# Console

![Build Status](https://gitoa.ru/go-4devs/console/actions/workflows/goaction.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/gitoa.ru/go-4devs/console)](https://goreportcard.com/report/gitoa.ru/go-4devs/console)
[![GoDoc](https://godoc.org/gitoa.ru/go-4devs/console?status.svg)](http://godoc.org/gitoa.ru/go-4devs/console)


## Creating a Command

Commands are defined in struct extending `pkg/command/create_user.go`. For example, you may want a command to create a user:

```go
package command

import (
	"context"

	"gitoa.ru/go-4devs/console/output"
	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/config"
)

func CreateUser() command.Command {
	return command.New(
		"app:create-user",
		"create user",
		func(ctx context.Context, in config.Provider, out output.Output) error {
			return nil
		},
	}
}
```
## Configure command

```go
func CreateUser() command.Command {
	return command.New(
    "app:create-user",
		"Creates a new user.",
		Execute,
    command.Help( "This command allows you to create a user..."),
	}
}

func Execute(ctx context.Context, in config.Provider, out output.Output) error{
	return nil
}
```


## Add arguments

```go
func CreateUser(required bool) command.Command {
	return command.New(
			"name",
			"description",
			Execute,
			command.Configure(Configure(required)),
		)
	}
}

func Configure(required bool) func(ctx context.Context, cfg config.Definition) error {
	return func (ctx context.Context, cfg config.Definition) error{
			var opts []func(*arg.Option)
			if required {
				opts = append(opts, arg.Required)
			}
			cfg.Add(
					arg.String("password", "User password", opts...)
				)

			return nil
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
	"gitoa.ru/go-4devs/console/example/pkg/command"
)

func main() {
	console.
		New().
		Add(
			command.CreateUser(false),
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
func CreateUser(required bool) command.Command {
	return command.New(
		"app:user:create",
		"create user",
		Execute,
	)
}

func Execute(ctx context.Context, in config.Provider, out output.Output) error {
	// outputs a message followed by a "\n"
	out.Println(ctx, "User Creator")
	out.Println(ctx, "Whoa!")

	// outputs a message without adding a "\n" at the end of the line
	out.Print(ctx, "You are about to ", "create a user.")

	return nil
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
func CreateUser() command.Command {
	return command.New(
		"app:user:create",
		"create user",
		Execute,
		command.Configure(Configure),
	)
}

func Configure(ctx context.Context, cfg config.Definition) error {
	cfg.Add(
		arg.String("username", "The username of the user.", arg.Required),
		arg.String("password", "User password"),
	)

	return nil
}

func Execute(ctx context.Context, in config.Provider, out output.Output) error {
	// outputs a message followed by a "\n"
	username, _ := in.Value(ctx, "username")
	out.Println(ctx, "User Creator")
	out.Println(ctx, "Username: ", username.String())

	return nil
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
	"gitoa.ru/go-4devs/config/provider/memory"
	"gitoa.ru/go-4devs/console/output"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	in := memory.Map{}
  in.SetOption("andrey","username")
	buf := bytes.Buffer{}
	out := output.Buffer(&buf)

	console.Run(ctx, command.CreateUser(false), in, out)

	expect := `User Creator
Username: andrey
`

	if expect != buf.String() {
		t.Errorf("expect: %s, got:%s", expect, buf.String())
	}
}
```
