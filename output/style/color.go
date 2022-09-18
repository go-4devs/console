package style

const (
	Black   Color = "0"
	Red     Color = "1"
	Green   Color = "2"
	Yellow  Color = "3"
	Blue    Color = "4"
	Magenta Color = "5"
	Cyan    Color = "6"
	White   Color = "7"
	Default Color = "9"
)

const (
	Bold       Option = "122"
	Underscore Option = "424"
	Blink      Option = "525"
	Reverse    Option = "727"
	Conseal    Option = "828"
)

const (
	ActionSet   = 1
	ActionUnset = 2
)

type Option string

func (o Option) Apply(action int) string {
	out := string(o)

	switch action {
	case ActionSet:
		return out[0:1]
	case ActionUnset:
		return out[1:]
	}

	return ""
}

type Color string

func (c Color) Apply(action int) string {
	if action == ActionSet {
		return string(c)
	}

	return string(Default)
}
