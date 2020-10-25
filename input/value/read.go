package value

import (
	"fmt"

	"gitoa.ru/go-4devs/console/input"
)

var _ input.AppendValue = (*Read)(nil)

type Read struct {
	input.Value
}

func (r *Read) Append(string) error {
	return fmt.Errorf("%w: read value", input.ErrInvalidName)
}
