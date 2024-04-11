package helpers

import (
	"moul.io/banner"
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

func NormalizeAppLogo(name string) string {
	re, _ := regexp.Compile(`[^a-z-_?.]+`)
	name = re.ReplaceAllString(strings.ToLower(name), " ")
	return banner.Inline(name)
}
