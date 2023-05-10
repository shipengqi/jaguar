package survey

import (
	"strconv"

	"github.com/manifoldco/promptui"
)

const (
	NumberEmpty = -1
)

func Select(label string, items []string) string {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}
	_, result, err := prompt.Run()
	if err != nil {
		return ""
	}
	return result
}

func InputString(label string, min, max int) string {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: StringValidator(min, max),
	}
	result, err := prompt.Run()
	if err != nil {
		return ""
	}
	return result
}

func InputNumber(label string, min, max int) int {
	prompt := promptui.Prompt{
		Label:    label,
		Validate: NumberValidator(min, max),
	}
	result, err := prompt.Run()
	if err != nil {
		return NumberEmpty
	}
	v, _ := strconv.Atoi(result)
	return v
}

func Confirm(label string) bool {
	prompt := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	result, err := prompt.Run()
	if err != nil {
		return false
	}
	v, err := strconv.ParseBool(result)
	if err != nil {
		return false
	}
	return v
}
