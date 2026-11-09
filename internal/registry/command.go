package registry

import (
	"context"
	"fmt"

	"github.com/go-4devs/console/command"
)

type ctxKey struct{}

//nolint:gochecknoglobals
var commands = command.Commands{}

// WithContext stores cmds for commands that look up other commands while running.
func WithContext(ctx context.Context, cmds *command.Commands) context.Context {
	return context.WithValue(ctx, ctxKey{}, cmds)
}

// Current returns the command list for this run, or the process-wide list.
func Current(ctx context.Context) *command.Commands {
	if cmds, ok := ctx.Value(ctxKey{}).(*command.Commands); ok && cmds != nil {
		return cmds
	}

	return &commands
}

func Find(name string) (command.Command, error) {
	prov, err := commands.Find(name)
	if err != nil {
		return prov, fmt.Errorf("%w", err)
	}

	return prov, nil
}

func Commands() []string {
	return commands.Names()
}

func Add(cmds ...command.Command) error {
	if err := commands.Add(cmds...); err != nil {
		return fmt.Errorf("add:%w", err)
	}

	return nil
}

func Set(cmds ...command.Command) error {
	if err := commands.Set(cmds...); err != nil {
		return fmt.Errorf("set:%w", err)
	}

	return nil
}
