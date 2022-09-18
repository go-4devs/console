package variable

//go:generate stringer -type=ArgType -linecomment

type ArgType int

const (
	TypeOption   ArgType = iota + 1 // option
	TypeArgument                    // argument
)
