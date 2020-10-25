package option

import (
	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

func Required(o *Option) {
	o.Flag |= flag.Required
}

func Default(in interface{}) func(*Option) {
	return func(o *Option) {
		o.Default = value.New(in)
	}
}

func Short(s string) func(*Option) {
	return func(o *Option) {
		o.Short = s
	}
}

func Array(o *Option) {
	o.Flag |= flag.Array
}

func Value(flag flag.Flag) func(*Option) {
	return func(o *Option) {
		o.Flag |= flag
	}
}

func Flag(in flag.Flag) func(*Option) {
	return func(o *Option) {
		o.Flag = in
	}
}

func Valid(f ...func(value.Value) error) func(*Option) {
	return func(o *Option) {
		o.Valid = f
	}
}

func New(name, description string, opts ...func(*Option)) Option {
	o := Option{
		Name:        name,
		Description: description,
	}

	for _, opt := range opts {
		opt(&o)
	}

	return o
}

type Option struct {
	Name        string
	Description string
	Short       string
	Flag        flag.Flag
	Default     value.Value
	Valid       []func(value.Value) error
}

func (o Option) HasShort() bool {
	return len(o.Short) == 1
}

func (o Option) HasDefault() bool {
	return o.Default != nil
}

func (o Option) IsBool() bool {
	return o.Flag.IsBool()
}

func (o Option) IsArray() bool {
	return o.Flag.IsArray()
}

func (o Option) IsRequired() bool {
	return o.Flag.IsRequired()
}

func (o Option) Validate(v value.Value) error {
	for _, valid := range o.Valid {
		if err := valid(v); err != nil {
			return Error(o.Name, err)
		}
	}

	return nil
}

func Error(name string, err error) error {
	return err
}
