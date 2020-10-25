package formatter

import "gitoa.ru/go-4devs/console/output/style"

func None() *Formatter {
	return New(
		WithStyle(func(name string) (style.Style, error) {
			if _, err := style.Find(name); err != nil {
				return style.Empty(), err
			}

			return style.Empty(), nil
		}))
}
