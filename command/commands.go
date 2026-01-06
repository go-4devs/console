package command

import (
	"fmt"
	"regexp"
	"sort"
	"sync"

	"gitoa.ru/go-4devs/console/errs"
	"gitoa.ru/go-4devs/console/setting"
)

var findCommand = regexp.MustCompile("([^:]+|)")

type Commands struct {
	sync.RWMutex

	cmds  []Command
	names map[string]int
}

func (c *Commands) Set(cmds ...Command) error {
	c.Lock()
	defer c.Unlock()

	return c.set(cmds...)
}

func (c *Commands) Add(cmds ...Command) error {
	c.Lock()
	defer c.Unlock()

	return c.add(cmds...)
}

func (c *Commands) Find(name string) (Command, error) {
	c.Lock()
	defer c.Unlock()

	return c.find(name)
}

func (c *Commands) Names() []string {
	c.Lock()
	defer c.Unlock()

	names := make([]string, 0, len(c.names))
	for name := range c.names {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

func (c *Commands) find(name string) (Command, error) {
	if idx, ok := c.names[name]; ok {
		return c.cmds[idx], nil
	}

	nameRegexp := findCommand.ReplaceAllStringFunc(name, func(in string) string {
		return in + "[^:]*"
	})

	findCommands := make([]Command, 0, len(c.cmds))

	cmdRegexp, err := regexp.Compile("^" + nameRegexp + "$")
	if err != nil {
		return Command{}, fmt.Errorf("find by regexp:%w", err)
	}

	for name, idx := range c.names {
		if cmdRegexp.MatchString(name) && !setting.IsHidden(c.cmds[idx]) {
			findCommands = append(findCommands, c.cmds[idx])
		}
	}

	if len(findCommands) == 1 {
		return findCommands[0], nil
	}

	if len(findCommands) > 1 {
		names := make([]string, len(findCommands))
		for i := range findCommands {
			names[i] = findCommands[i].Name()
		}

		return Command{}, errs.AlternativesError{Alt: names, Err: errs.ErrCommandDplicate}
	}

	return Command{}, fmt.Errorf("%w", errs.ErrNotFound)
}

func (c *Commands) set(cmds ...Command) error {
	if c.names == nil {
		c.names = make(map[string]int, len(cmds))
	}

	for _, cmd := range cmds {
		if cmd.IsZero() {
			return fmt.Errorf("command:%w", errs.ErrCommandNil)
		}

		if idx, ok := c.names[cmd.Name()]; ok {
			c.cmds[idx] = cmd

			continue
		}

		c.names[cmd.Name()] = len(c.cmds)
		c.cmds = append(c.cmds, cmd)
	}

	return nil
}

func (c *Commands) add(cmds ...Command) error {
	if c.names == nil {
		c.names = make(map[string]int, len(cmds))
	}

	for _, cmd := range cmds {
		if cmd.IsZero() {
			return fmt.Errorf("command:%w", errs.ErrCommandNil)
		}

		if _, ok := c.names[cmd.Name()]; ok {
			return fmt.Errorf("command %s:%w", cmd.Name(), errs.ErrCommandDplicate)
		}

		c.names[cmd.Name()] = len(c.cmds)
		c.cmds = append(c.cmds, cmd)
	}

	return nil
}
