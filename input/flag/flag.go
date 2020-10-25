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

func (f Flag) With(v Flag) Flag {
	return f | v
}

func (f Flag) IsString() bool {
	return f|Required|Array^Required^Array == 0
}

func (f Flag) IsRequired() bool {
	return f&Required > 0
}

func (f Flag) IsArray() bool {
	return f&Array > 0
}

func (f Flag) IsInt() bool {
	return f&Int > 0
}

func (f Flag) IsInt64() bool {
	return f&Int64 > 0
}

func (f Flag) IsUint() bool {
	return f&Uint > 0
}

func (f Flag) IsUint64() bool {
	return f&Uint64 > 0
}

func (f Flag) IsFloat64() bool {
	return f&Float64 > 0
}

func (f Flag) IsBool() bool {
	return f&Bool > 0
}

func (f Flag) IsDuration() bool {
	return f&Duration > 0
}

func (f Flag) IsTime() bool {
	return f&Time > 0
}

func (f Flag) IsAny() bool {
	return f&Any > 0
}

func (f Flag) Type() Flag {
	switch {
	case f.IsInt():
		return Int
	case f.IsInt64():
		return Int64
	case f.IsUint():
		return Uint
	case f.IsUint64():
		return Uint64
	case f.IsFloat64():
		return Float64
	case f.IsBool():
		return Bool
	case f.IsDuration():
		return Duration
	case f.IsTime():
		return Time
	case f.IsAny():
		return Any
	default:
		return String
	}
}
