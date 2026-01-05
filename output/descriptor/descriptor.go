package descriptor

import (
	"context"
	"errors"
	"sync"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/console/output"
)

var ErrDescriptorNotFound = errors.New("descriptor not found")

//nolint:gochecknoglobals
var (
	descriptors = map[string]Descriptor{
		"txt": &txt{},
	}
	descriptorMu sync.Mutex
)

type Command struct {
	config.Options

	Bin         string
	Name        string
	Description string
	Help        string
}

type Commands struct {
	config.Options

	Namespace string
	Commands  []NSCommand
}

type NSCommand struct {
	Name     string
	Commands []ShortCommand
}

func (n *NSCommand) Append(name, desc string) {
	n.Commands = append(n.Commands, ShortCommand{Name: name, Description: desc})
}

type ShortCommand struct {
	Name        string
	Description string
}

type Descriptor interface {
	Command(ctx context.Context, out output.Output, cmd Command) error
	Commands(ctx context.Context, out output.Output, cmds Commands) error
}

func Find(name string) (Descriptor, error) {
	descriptorMu.Lock()
	defer descriptorMu.Unlock()

	if d, has := descriptors[name]; has {
		return d, nil
	}

	return nil, ErrDescriptorNotFound
}

func Descriptors() []string {
	names := make([]string, 0, len(descriptors))

	for name := range descriptors {
		names = append(names, name)
	}

	return names
}

func Register(name string, descriptor Descriptor) {
	descriptorMu.Lock()
	defer descriptorMu.Unlock()

	if descriptor == nil {
		panic("console: Register descriptor is nil")
	}

	if _, has := descriptors[name]; has {
		panic("console: Register called twice for descriptor " + name)
	}

	descriptors[name] = descriptor
}
