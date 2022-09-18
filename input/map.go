package input

import (
	"context"
	"fmt"
	"sync"

	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/variable"
)

type Map struct {
	opts map[string]value.Value
	args map[string]value.Value
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

func (m *Map) SetOption(name string, val interface{}) {
	m.Lock()
	defer m.Unlock()

	if m.opts == nil {
		m.opts = make(map[string]value.Value)
	}

	m.opts[name] = value.New(val)
}

func (m *Map) HasArgument(name string) bool {
	_, ok := m.args[name]

	return ok
}

func (m *Map) SetArgument(name string, val interface{}) {
	m.Lock()
	defer m.Unlock()

	if m.args == nil {
		m.args = make(map[string]value.Value)
	}

	m.args[name] = value.New(val)
}

func (m *Map) AppendOption(opt variable.Variable, val string) error {
	old, ok := m.opts[opt.Name]
	if !ok {
		value, err := opt.Create(val)
		if err != nil {
			return fmt.Errorf("append option:%w", err)
		}

		m.SetOption(opt.Name, value)

		return nil
	}

	value, err := opt.Append(old, val)
	if err != nil {
		return fmt.Errorf("append option:%w", err)
	}

	m.SetOption(opt.Name, value)

	return nil
}

func (m *Map) AppendArgument(arg variable.Variable, val string) error {
	old, ok := m.args[arg.Name]
	if !ok {
		value, err := arg.Create(val)
		if err != nil {
			return fmt.Errorf("append option:%w", err)
		}

		m.SetArgument(arg.Name, value)

		return nil
	}

	value, err := arg.Append(old, val)
	if err != nil {
		return fmt.Errorf("append option:%w", err)
	}

	m.SetArgument(arg.Name, value)

	return nil
}
