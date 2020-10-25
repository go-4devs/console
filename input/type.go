package input

//go:generate stringer -type=Type -linecomment

type Type int

const (
	Argument Type = iota // argument
	Option               // option
)

func (i Type) IsArgument() bool {
	return i == Argument
}

func (i Type) IsOption() bool {
	return i == Option
}
