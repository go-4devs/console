package array

import (
	"context"
	"sync"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/wrap"
)

var _ input.ReadInput = (*Input)(nil)

func Argument(name string, v interface{}) func(*Input) {
	return func(i *Input) {
		i.args[name] = value.New(v)
	}
}

func Option(name string, v interface{}) func(*Input) {
	return func(i *Input) {
		i.opt[name] = value.New(v)
	}
}

func New(opts ...func(*Input)) *wrap.Input {
	i := &Input{
		args: make(map[string]input.Value),
		opt:  make(map[string]input.Value),
	}

	for _, opt := range opts {
		opt(i)
	}

	return &wrap.Input{ReadInput: i}
}

type Input struct {
	args map[string]input.Value
	opt  map[string]input.Value
	mu   sync.Mutex
}

func (i *Input) ReadOption(_ context.Context, name string) (input.Value, error) {
	if o, has := i.opt[name]; has {
		return o, nil
	}

	return nil, input.ErrorOption(name, input.ErrNotFound)
}

func (i *Input) HasOption(name string) bool {
	_, has := i.opt[name]

	return has
}

func (i *Input) SetOption(name string, val input.Value) {
	i.mu.Lock()
	i.opt[name] = val
	i.mu.Unlock()
}

func (i *Input) ReadArgument(_ context.Context, name string) (input.Value, error) {
	if a, has := i.args[name]; has {
		return a, nil
	}

	return nil, input.ErrorArgument(name, input.ErrNotFound)
}

func (i *Input) HasArgument(name string) bool {
	_, has := i.args[name]

	return has
}

func (i *Input) SetArgument(name string, val input.Value) {
	i.mu.Lock()
	i.args[name] = val
	i.mu.Unlock()
}

func (i *Input) Bind(_ context.Context, def *input.Definition) error {
	return nil
}
