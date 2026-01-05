package registry

import (
	"fmt"

	"gitoa.ru/go-4devs/console/command"
)

//nolint:gochecknoglobals
var commands = command.Commands{}

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
