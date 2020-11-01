package input

import (
	"context"
	"sync"

	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/value/flag"
)

type Map struct {
	opts map[string]value.Append
	args map[string]value.Append
	sync.Mutex
}

func (m *Map) Option(_ context.Context, name string) value.Value {
	m.Lock()
	defer m.Unlock()

	return m.opts[name]
}

func (m *Map) Argument(_ context.Context, name string) value.Value {
	m.Lock()
	defer m.Unlock()

	return m.args[name]
}

func (m *Map) Bind(_ context.Context, _ *Definition) error {
	return nil
}

func (m *Map) LenArguments() int {
	return len(m.args)
}

func (m *Map) HasOption(name string) bool {
	_, ok := m.opts[name]

	return ok
}

func (m *Map) SetOption(name string, v interface{}) {
	m.Lock()
	defer m.Unlock()

	if m.opts == nil {
		m.opts = make(map[string]value.Append)
	}

	m.opts[name] = value.New(v)
}

func (m *Map) HasArgument(name string) bool {
	_, ok := m.args[name]

	return ok
}

func (m *Map) SetArgument(name string, v interface{}) {
	m.Lock()
	defer m.Unlock()

	if m.args == nil {
		m.args = make(map[string]value.Append)
	}

	m.args[name] = value.New(v)
}

func (m *Map) AppendOption(f flag.Flag, name, val string) error {
	if _, ok := m.opts[name]; !ok {
		m.SetOption(name, value.ByFlag(f))
	}

	return m.opts[name].Append(val)
}

func (m *Map) AppendArgument(f flag.Flag, name, val string) error {
	if _, ok := m.args[name]; !ok {
		m.SetArgument(name, value.ByFlag(f))
	}

	return m.args[name].Append(val)
}
