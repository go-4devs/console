package input

func NewOption(name, description string, opts ...func(*Option)) Option {
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
	Flag        Flag
	Default     Value
	Valid       []func(Value) error
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

func (o Option) Validate(v Value) error {
	for _, valid := range o.Valid {
		if err := valid(v); err != nil {
			return ErrorOption(o.Name, err)
		}
	}

	return nil
}
