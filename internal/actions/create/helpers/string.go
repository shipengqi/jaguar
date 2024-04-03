package helpers

import (
	"regexp"
	"strings"
)

func NormalizeAppName(value string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	value = reg.ReplaceAllString(value, "")
	return strings.ToLower(value)
}

func NormalizeAppEnv(value string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
	value = reg.ReplaceAllString(value, "_")
	value = strings.Trim(value, "_")
	return strings.ToUpper(value)
}
