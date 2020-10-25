package validator

import "gitoa.ru/go-4devs/console/input/value"

func Valid(v ...func(value.Value) error) func(value.Value) error {
	return func(in value.Value) error {
		for _, valid := range v {
			if err := valid(in); err != nil {
				return err
			}
		}

		return nil
	}
}
