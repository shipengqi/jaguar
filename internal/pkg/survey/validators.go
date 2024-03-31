package survey

import (
	"errors"
	"fmt"
	"strings"
)

func NumberValidator(min, max int) func(input interface{}) error {
	return func(input interface{}) error {
		v, ok := input.(int)
		if !ok {
			return errors.New("value must be a valid string")
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

func StringValidator(minLength, maxLength int) func(input interface{}) error {
	return func(input interface{}) error {
		str, ok := input.(string)
		if !ok {
			return errors.New("value must be a valid string")
		}
		str = strings.TrimSpace(str)
		if minLength != NumberEmpty && len(str) < minLength {
			return fmt.Errorf("length must be at least %d", minLength)
		}
		if maxLength != NumberEmpty && len(str) > maxLength {
			return fmt.Errorf("length cannot exceed %d", maxLength)
		}
		return nil
	}
}
