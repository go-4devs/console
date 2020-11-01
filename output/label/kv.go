package label

import (
	"fmt"
	"strings"
)

var (
	_ fmt.Stringer = KeyValue{}
	_ fmt.Stringer = KeyValues{}
)

type KeyValues []KeyValue

func (kv KeyValues) String() string {
	s := make([]string, len(kv))
	for i, v := range kv {
		s[i] = v.String()
	}

	return strings.Join(s, ", ")
}

type KeyValue struct {
	Key   Key
	Value Value
}

func (k KeyValue) String() string {
	return string(k.Key) + "=\"" + k.Value.String() + "\""
}

func Any(k string, v interface{}) KeyValue {
	return Key(k).Any(v)
}

func Bool(k string, v bool) KeyValue {
	return Key(k).Bool(v)
}

func Int(k string, v int) KeyValue {
	return Key(k).Int(v)
}

func Int64(k string, v int64) KeyValue {
	return Key(k).Int64(v)
}

func Uint(k string, v uint) KeyValue {
	return Key(k).Uint(v)
}

func Uint64(k string, v uint64) KeyValue {
	return Key(k).Uint64(v)
}

func Float64(k string, v float64) KeyValue {
	return Key(k).Float64(v)
}

func String(k string, v string) KeyValue {
	return Key(k).String(v)
}
