package ui

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	NumberEmpty = -1
)

func NumberValidator(field string, minVal, maxVal int) func(string) error {
	return func(input string) error {
		v, err := strconv.Atoi(input)
		if err != nil {
			return fmt.Errorf("%s must be a valid number", field)
		}
		if minVal != NumberEmpty && v < minVal {
			return fmt.Errorf("%s must be at least %d", field, minVal)
		}
		if maxVal != NumberEmpty && v > maxVal {
			return fmt.Errorf("%s cannot exceed %d", field, maxVal)
		}
		return nil
	}
}

func StringValidator(field string, minLength, maxLength int) func(string) error {
	return func(input string) error {
		input = strings.TrimSpace(input)
		if minLength != NumberEmpty && len(input) < minLength {
			return fmt.Errorf("%s length must be at least %d", field, minLength)
		}
		if maxLength != NumberEmpty && len(input) > maxLength {
			return fmt.Errorf("%s length cannot exceed %d", field, maxLength)
		}
		return nil
	}
}
