package console

import (
	"fmt"

	"gitoa.ru/go-4devs/console/command"
	"gitoa.ru/go-4devs/console/internal/registry"
)

// MustRegister register command or panic if err.
func MustRegister(cmd ...command.Command) {
	err := registry.Add(cmd...)
	if err != nil {
		panic(err)
	}
}

// Register makes a command available execute in app. If Register is called twice with the same name or if driver is nil, return error.
func Register(cmd ...command.Command) error {
	if err := registry.Add(cmd...); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

// Find command by name, tries to find the best match if you give it an abbreviation of a name.
func Find(name string) (command.Command, error) {
	cmd, err := registry.Find(name)
	if err != nil {
		return cmd, fmt.Errorf("%w", err)
	}

	return cmd, nil
}
