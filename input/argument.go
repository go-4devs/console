package input

func NewArgument(name, description string, opts ...func(*Argument)) Argument {
	a := Argument{
		Name:        name,
		Description: description,
	}

	for _, opt := range opts {
		opt(&a)
	}

	return a
}

type Argument struct {
	Name        string
	Description string
	Default     Value
	Flag        Flag
	Valid       []func(Value) error
}

func (a Argument) HasDefault() bool {
	return a.Default != nil
}

func (a Argument) IsBool() bool {
	return a.Flag.IsBool()
}

func (a Argument) IsRequired() bool {
	return a.Flag.IsRequired()
}

func (a Argument) IsArray() bool {
	return a.Flag.IsArray()
}

func (a Argument) Validate(v Value) error {
	for _, valid := range a.Valid {
		if err := valid(v); err != nil {
			return ErrorArgument(a.Name, err)
		}
	}

	return nil
}
