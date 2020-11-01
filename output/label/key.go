package label

type Key string

func (k Key) Any(v interface{}) KeyValue {
	return KeyValue{
		Key:   k,
		Value: AnyValue(v),
	}
}

func (k Key) Bool(v bool) KeyValue {
	return KeyValue{
		Key:   k,
		Value: BoolValue(v),
	}
}

func (k Key) Int(v int) KeyValue {
	return KeyValue{
		Key:   k,
		Value: IntValue(v),
	}
}

func (k Key) Int64(v int64) KeyValue {
	return KeyValue{
		Key:   k,
		Value: Int64Value(v),
	}
}

func (k Key) Uint(v uint) KeyValue {
	return KeyValue{
		Key:   k,
		Value: UintValue(v),
	}
}

func (k Key) Uint64(v uint64) KeyValue {
	return KeyValue{
		Key:   k,
		Value: Uint64Value(v),
	}
}

func (k Key) Float64(v float64) KeyValue {
	return KeyValue{
		Key:   k,
		Value: Float64Value(v),
	}
}

func (k Key) String(v string) KeyValue {
	return KeyValue{
		Key:   k,
		Value: StringValue(v),
	}
}
