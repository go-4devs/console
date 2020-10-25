package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/console/input/flag"
	"gitoa.ru/go-4devs/console/input/validator"
	"gitoa.ru/go-4devs/console/input/value"
)

func TestValid(t *testing.T) {
	validValue := value.New("one")
	invalidValue := value.New([]string{"one"})

	valid := validator.Valid(
		validator.NotBlank(flag.String),
		validator.Enum("one", "two"),
	)

	if err := valid(validValue); err != nil {
		t.Errorf("expected valid value, got: %s", err)
	}

	if err := valid(invalidValue); !errors.Is(err, validator.ErrNotBlank) {
		t.Errorf("expected not blank, got:%s", err)
	}
}
