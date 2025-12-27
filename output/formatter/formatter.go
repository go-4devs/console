package formatter

import (
	"bytes"
	"context"
	"regexp"

	"gitoa.ru/go-4devs/console/output/style"
)

var re = regexp.MustCompile(`<(([a-z][^<>]+)|/([a-z][^<>]+)?)>`)

func WithStyle(styles func(string) (style.Style, error)) func(*Formatter) {
	return func(f *Formatter) {
		f.styles = styles
	}
}

func New(opts ...func(*Formatter)) *Formatter {
	formatter := &Formatter{
		styles: style.Find,
	}

	for _, opt := range opts {
		opt(formatter)
	}

	return formatter
}

type Formatter struct {
	styles func(string) (style.Style, error)
}

func (a *Formatter) Format(_ context.Context, msg string) string {
	var (
		out bytes.Buffer
		cur int
	)

	for _, idx := range re.FindAllStringIndex(msg, -1) {
		tag := msg[idx[0]+1 : idx[1]-1]

		if cur < idx[0] {
			out.WriteString(msg[cur:idx[0]])
		}

		var (
			st  style.Style
			err error
		)

		switch tag[0:1] {
		case "/":
			st, err = a.styles(tag[1:])
			if err == nil {
				out.WriteString(st.Set(style.ActionUnset))
			}
		default:
			st, err = a.styles(tag)
			if err == nil {
				out.WriteString(st.Set(style.ActionSet))
			}
		}

		if err != nil {
			cur = idx[0]
		} else {
			cur = idx[1]
		}
	}

	if len(msg) > cur {
		out.WriteString(msg[cur:])
	}

	return out.String()
}
