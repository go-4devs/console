package validator_test

import (
	"errors"
	"testing"

	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/validator"
)

func TestEnum(t *testing.T) {
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
