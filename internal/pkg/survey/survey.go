package survey

import (
	"github.com/AlecAivazis/survey/v2"
)

const (
	NumberEmpty = -1
)

func Select(label string, options []string) (string, error) {
	prompt := &survey.Select{
		Message: label,
		Options: options,
	}
	result := ""
	err := survey.AskOne(prompt, &result, survey.WithValidator(survey.Required))
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputString(label string, minLength, maxLength int) (string, error) {
	prompt := &survey.Input{
		Message: label,
	}
	result := ""
	err := survey.AskOne(prompt, &result,
		survey.WithValidator(survey.MinLength(minLength)),
		survey.WithValidator(survey.MaxLength(maxLength)))
	if err != nil {
		return "", err
	}
	return result, nil
}

func InputNumber(label string, min, max int) (int, error) {
	prompt := &survey.Input{
		Message: label,
	}
	var result int
	err := survey.AskOne(prompt, &result,
		survey.WithValidator(NumberValidator(min, max)))
	if err != nil {
		return NumberEmpty, err
	}
	return result, nil
}

func Confirm(label string) (bool, error) {
	prompt := &survey.Confirm{
		Message: label,
	}
	var result bool
	err := survey.AskOne(prompt, &result)
	if err != nil {
		return false, err
	}
	return result, nil
}
