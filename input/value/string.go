package value

import "gitoa.ru/go-4devs/console/input/flag"

type String struct {
	Empty
	Val  []string
	Flag flag.Flag
}

func (s *String) Append(in string) error {
	s.Val = append(s.Val, in)

	return nil
}

func (s *String) String() string {
	if s.Flag.IsArray() {
		return ""
	}

	if len(s.Val) == 1 {
		return s.Val[0]
	}

	return ""
}

func (s *String) Strings() []string {
	return s.Val
}

func (s *String) Any() interface{} {
	if s.Flag.IsArray() {
		return s.Strings()
	}

	return s.String()
}
