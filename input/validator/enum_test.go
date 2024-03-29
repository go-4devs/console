package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/console/input/validator"
	"gitoa.ru/go-4devs/console/input/value"
)

func TestEnum(t *testing.T) {
	t.Parallel()

	validValue := value.New("valid")
	invalidValue := value.New("invalid")

	enum := validator.Enum("valid", "other", "three")

	if err := enum(validValue); err != nil {
		t.Errorf("expected valid value got err:%s", err)
	}

	if err := enum(invalidValue); !errors.Is(err, validator.ErrInvalid) {
		t.Errorf("expected err:%s, got: %s", validator.ErrInvalid, err)
	}
}
