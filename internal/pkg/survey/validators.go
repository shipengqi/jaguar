package survey

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func NumberValidator(min, max int) func(input string) error {
	return func(input string) error {
		v, err := strconv.Atoi(input)
		if err != nil {
			return errors.New("value must be a valid integer")
		}
		if min != NumberEmpty && v < min {
			return fmt.Errorf("value must be at least %d", min)
		}
		if max != NumberEmpty && v > max {
			return fmt.Errorf("value cannot exceed %d", max)
		}
		return nil
	}
}

func StringValidator(min, max int) func(input string) error {
	return func(input string) error {
		input = strings.TrimSpace(input)
		if min != NumberEmpty && len(input) < min {
			return fmt.Errorf("length must be at least %d", min)
		}
		if max != NumberEmpty && len(input) > max {
			return fmt.Errorf("length cannot exceed %d", max)
		}
		return nil
	}
}
