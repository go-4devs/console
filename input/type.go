package input

//go:generate stringer -type=Type -linecomment

type Type int

const (
	Argument Type = iota // argument
	Option               // option
)

func (t Type) IsArgument() bool {
	return t == Argument
}

func (t Type) IsOption() bool {
	return t == Option
}
