package flag

//go:generate stringer -type=Flag -linecomment

type Flag int

const (
	String   Flag = 0         // string
	Required Flag = 1 << iota // required
	Array                     // array
	Int                       // int
	Int64                     // int64
	Uint                      // uint
	Uint64                    // uint64
	Float64                   // float64
	Bool                      // bool
	Duration                  // duration
	Time                      // time
	Any                       // any
)

func (i Flag) With(v Flag) Flag {
	return i | v
}

func (i Flag) IsString() bool {
	return i|Required|Array^Required^Array == 0
}

func (i Flag) IsRequired() bool {
	return i&Required > 0
}

func (i Flag) IsArray() bool {
	return i&Array > 0
}

func (i Flag) IsInt() bool {
	return i&Int > 0
}

func (i Flag) IsInt64() bool {
	return i&Int64 > 0
}

func (i Flag) IsUint() bool {
	return i&Uint > 0
}

func (i Flag) IsUint64() bool {
	return i&Uint64 > 0
}

func (i Flag) IsFloat64() bool {
	return i&Float64 > 0
}

func (i Flag) IsBool() bool {
	return i&Bool > 0
}

func (i Flag) IsDuration() bool {
	return i&Duration > 0
}

func (i Flag) IsTime() bool {
	return i&Time > 0
}

func (i Flag) IsAny() bool {
	return i&Any > 0
}

//nolint:cyclop
func (i Flag) Type() Flag {
	switch {
	case i.IsInt():
		return Int
	case i.IsInt64():
		return Int64
	case i.IsUint():
		return Uint
	case i.IsUint64():
		return Uint64
	case i.IsFloat64():
		return Float64
	case i.IsBool():
		return Bool
	case i.IsDuration():
		return Duration
	case i.IsTime():
		return Time
	case i.IsAny():
		return Any
	default:
		return String
	}
}
