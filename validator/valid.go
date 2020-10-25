package validator

import "gitoa.ru/go-4devs/console/input"

func Valid(v ...func(input.Value) error) func(input.Value) error {
	return func(in input.Value) error {
		for _, valid := range v {
			if err := valid(in); err != nil {
				return err
			}
		}

		return nil
	}
}
