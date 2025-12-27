package descriptor

import (
	"context"
	"errors"
	"sort"
	"strings"
	"sync"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/provider/arg"
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
	Bin         string
	Name        string
	Description string
	Help        string
	Definition  Definition
}

type Commands struct {
	Namespace  string
	Definition Definition
	Commands   []NSCommand
}

type NSCommand struct {
	Name     string
	Commands []ShortCommand
}

func (n *NSCommand) Append(name, desc string) {
	n.Commands = append(n.Commands, ShortCommand{Name: name, Description: desc})
}

func NewDefinition(opts []config.Variable) Definition {
	type data struct {
		name string
		pos  uint64
		opt  config.Variable
	}

	posArgs := make([]data, 0, len(opts))

	posOpt := make([]data, 0, len(opts))
	for _, opt := range opts {
		pos, ok := arg.ParamArgument(opt)
		if !ok {
			pos, _ = option.DataPosition(opt)
			posOpt = append(posOpt, data{pos: pos, opt: opt})

			continue
		}

		posArgs = append(posArgs, data{name: strings.Join(opt.Key(), "."), pos: pos, opt: opt})
	}

	sort.Slice(posArgs, func(i, j int) bool {
		return posArgs[i].pos > posArgs[j].pos && posArgs[i].name > posArgs[j].name
	})

	sort.Slice(posOpt, func(i, j int) bool {
		return posOpt[i].pos < posOpt[j].pos
	})

	args := make([]config.Variable, len(posArgs))
	for idx := range posArgs {
		args[idx] = posArgs[idx].opt
	}

	options := make([]config.Variable, len(posOpt))
	for idx := range posOpt {
		options[idx] = posOpt[idx].opt
	}

	return Definition{
		options: options,
		args:    args,
	}
}

type Definition struct {
	args    []config.Variable
	options []config.Variable
}

func (d Definition) Arguments() []config.Variable {
	return d.args
}

func (d Definition) Options() []config.Variable {
	return d.options
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
