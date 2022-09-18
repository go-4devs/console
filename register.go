package console

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
	"sync"
)

const (
	CommandHelp = "help"
	CommandList = "list"
)

var (
	ErrNotFound         = errors.New("command not found")
	ErrCommandNil       = errors.New("console: Register command is nil")
	ErrCommandDuplicate = errors.New("console: duplicate command")
)

//nolint:gochecknoglobals
var (
	commandsMu  sync.RWMutex
	commands    = make(map[string]*Command)
	findCommand = regexp.MustCompile("([^:]+|)")
)

type AlternativesError struct {
	alt []string
	err error
}

func (e AlternativesError) Error() string {
	return fmt.Sprintf("%s, alternatives: [%s]", e.err, strings.Join(e.alt, ","))
}

func (e AlternativesError) Is(err error) bool {
	return errors.Is(e.err, err)
}

func (e AlternativesError) Unwrap() error {
	return e.err
}

func (e AlternativesError) Alternatives() []string {
	return e.alt
}

// MustRegister register command or panic if err.
func MustRegister(cmd *Command) {
	if err := Register(cmd); err != nil {
		panic(err)
	}
}

// Register makes a command available execute in app. If Register is called twice with the same name or if driver is nil, return error.
func Register(cmd *Command) error {
	if cmd == nil {
		return ErrCommandNil
	}

	if _, err := Find(cmd.Name); !errors.Is(err, ErrNotFound) {
		return fmt.Errorf("%w: command %s", ErrCommandDuplicate, cmd.Name)
	}

	register(cmd)

	return nil
}

func register(cmd *Command) {
	commandsMu.Lock()
	defer commandsMu.Unlock()

	if cmd != nil && cmd.Name != "" {
		commands[cmd.Name] = cmd
	}
}

// Commands returns a sorted list of the names of the registered commands.
func Commands() []string {
	commandsMu.RLock()
	defer commandsMu.RUnlock()

	return commandNames()
}

func commandNames() []string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}

	sort.Strings(names)

	return names
}

// Find command by name, tries to find the best match if you give it an abbreviation of a name.
func Find(name string) (*Command, error) {
	commandsMu.RLock()
	defer commandsMu.RUnlock()

	if cmd, ok := commands[name]; ok {
		return cmd, nil
	}

	nameRegexp := findCommand.ReplaceAllStringFunc(name, func(in string) string {
		return in + "[^:]*"
	})

	findCommands := make([]*Command, 0)

	cmdRegexp, err := regexp.Compile("^" + nameRegexp + "$")
	if err != nil {
		return nil, fmt.Errorf("find by regexp:%w", err)
	}

	for name := range commands {
		if !commands[name].Hidden && cmdRegexp.MatchString(name) {
			findCommands = append(findCommands, commands[name])
		}
	}

	if len(findCommands) == 1 {
		return findCommands[0], nil
	}

	if len(findCommands) > 1 {
		names := make([]string, len(findCommands))
		for i := range findCommands {
			names[i] = findCommands[i].Name
		}

		return nil, AlternativesError{alt: names, err: ErrNotFound}
	}

	return nil, ErrNotFound
}
