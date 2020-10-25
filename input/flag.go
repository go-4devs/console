package input

//go:generate stringer -type=Flag -linecomment

type Flag int

const (
	ValueString   Flag = 0         // string
	ValueRequired Flag = 1 << iota // required
	ValueArray                     // array
	ValueInt                       // int
	ValueInt64                     // int64
	ValueUint                      // uint
	ValueUint64                    // uint64
	ValueFloat64                   // float64
	ValueBool                      // bool
	ValueDuration                  // duration
	ValueTime                      // time
	ValueAny                       // any
)

func (f Flag) Type() Flag {
	return Type(f)
}

func (f Flag) With(v Flag) Flag {
	return f | v
}

func (f Flag) IsString() bool {
	return f|ValueRequired|ValueArray^ValueRequired^ValueArray == 0
}

func (f Flag) IsRequired() bool {
	return f&ValueRequired > 0
}

func (f Flag) IsArray() bool {
	return f&ValueArray > 0
}

func (f Flag) IsInt() bool {
	return f&ValueInt > 0
}

func (f Flag) IsInt64() bool {
	return f&ValueInt64 > 0
}

func (f Flag) IsUint() bool {
	return f&ValueUint > 0
}

func (f Flag) IsUint64() bool {
	return f&ValueUint64 > 0
}

func (f Flag) IsFloat64() bool {
	return f&ValueFloat64 > 0
}

func (f Flag) IsBool() bool {
	return f&ValueBool > 0
}

func (f Flag) IsDuration() bool {
	return f&ValueDuration > 0
}

func (f Flag) IsTime() bool {
	return f&ValueTime > 0
}

func (f Flag) IsAny() bool {
	return f&ValueAny > 0
}
