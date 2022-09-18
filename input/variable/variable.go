package variable

import (
	"fmt"

	"gitoa.ru/go-4devs/console/input/errs"
	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/value"
)

type Option func(*Variable)

func WithType(t ArgType) Option {
	return func(v *Variable) {
		v.Type = t
	}
}

func ArgOption(v *Variable) {
	v.Type = TypeOption
}

func ArgArgument(v *Variable) {
	v.Type = TypeArgument
}

func Value(in flag.Flag) Option {
	return func(v *Variable) {
		v.Flag |= in
	}
}

func Default(in value.Value) Option {
	return func(v *Variable) {
		v.Default = in
	}
}

func Required(v *Variable) {
	v.Flag |= flag.Required
}

func Hidden(v *Variable) {
	v.hidden = true
}

func WithParse(create Create, update Append) Option {
	return func(v *Variable) {
		v.append = func(Param) Append { return update }
		v.create = func(Param) Create { return create }
	}
}

func WithParamParse(create func(Param) Create, update func(Param) Append) Option {
	return func(v *Variable) {
		v.append = update
		v.create = create
	}
}

func Valid(f ...func(value.Value) error) Option {
	return func(v *Variable) {
		v.Valid = f
	}
}

func Array(o *Variable) {
	o.Flag |= flag.Array
}

func WithParam(name string, fn func(interface{}) error) Option {
	return func(v *Variable) {
		v.params[name] = fn
	}
}

type (
	Create func(s string) (value.Value, error)
	Append func(old value.Value, s string) (value.Value, error)
)

func New(name, description string, opts ...Option) Variable {
	res := Variable{
		Name:        name,
		Description: description,
		Type:        TypeOption,
		create:      func(Param) Create { return CreateString },
		append:      func(Param) Append { return AppendString },
		params:      make(Params),
	}

	for _, opt := range opts {
		opt(&res)
	}

	return res
}

type Variable struct {
	Name        string
	Description string
	Alias       string
	Flag        flag.Flag
	Type        ArgType
	Default     value.Value
	hidden      bool
	Valid       []func(value.Value) error
	params      Params
	create      func(Param) Create
	append      func(Param) Append
}

func (v Variable) Validate(in value.Value) error {
	for _, valid := range v.Valid {
		if err := valid(in); err != nil {
			return Err(v.Name, v.Type, err)
		}
	}

	return nil
}

func (v Variable) IsHidden() bool {
	return v.hidden
}

func (v Variable) IsArray() bool {
	return v.Flag.IsArray()
}

func (v Variable) IsRequired() bool {
	return v.Flag.IsRequired()
}

func (v Variable) HasDefault() bool {
	return v.Default != nil
}

func (v Variable) IsBool() bool {
	return v.Flag.IsBool()
}

func (v Variable) HasShort() bool {
	return v.Type == TypeOption && len(v.Alias) == 1
}

func (v Variable) Create(s string) (value.Value, error) {
	return v.create(v.params)(s)
}

func (v Variable) Append(old value.Value, s string) (value.Value, error) {
	return v.append(v.params)(old, s)
}

type Param interface {
	Value(name string, v interface{}) error
}

type Params map[string]func(interface{}) error

func (p Params) Value(name string, v interface{}) error {
	if p, ok := p[name]; ok {
		return p(v)
	}

	return fmt.Errorf("%w: param %v", errs.ErrNotFound, name)
}
