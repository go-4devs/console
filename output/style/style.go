package style

import (
	"errors"
	"fmt"
	"strings"
	"sync"
)

//nolint:gochecknoglobals
var (
	styles = map[string]Style{
		"error":    {Foreground: White, Background: Red},
		"info":     {Foreground: Green},
		"comment":  {Foreground: Yellow},
		"question": {Foreground: Black, Background: Cyan},
	}
	stylesMu sync.Mutex
	empty    = Style{}
)

var (
	ErrNotFound       = errors.New("console: style not found")
	ErrDuplicateStyle = errors.New("console: Register called twice")
)

func Empty() Style {
	return empty
}

func Find(name string) (Style, error) {
	if st, has := styles[name]; has {
		return st, nil
	}

	return empty, ErrNotFound
}

func Register(name string, style Style) error {
	stylesMu.Lock()
	defer stylesMu.Unlock()

	if _, has := styles[name]; has {
		return fmt.Errorf("%w for style %s", ErrDuplicateStyle, name)
	}

	styles[name] = style

	return nil
}

func MustRegister(name string, style Style) {
	if err := Register(name, style); err != nil {
		panic(err)
	}
}

type Style struct {
	Background Color
	Foreground Color
	Options    []Option
}

func (s Style) Apply(msg string) string {
	return s.Set(ActionSet) + msg + s.Set(ActionUnset)
}

func (s Style) Set(action int) string {
	style := make([]string, 0, len(s.Options))

	if s.Foreground != "" {
		style = append(style, "3"+s.Foreground.Apply(action))
	}

	if s.Background != "" {
		style = append(style, "4"+s.Background.Apply(action))
	}

	for _, opt := range s.Options {
		style = append(style, opt.Apply(action))
	}

	if len(style) == 0 {
		return ""
	}

	return "\033[" + strings.Join(style, ";") + "m"
}
