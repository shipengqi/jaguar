package helpers

import (
	"fmt"
	"testing"
)

func TestNormalizeAppName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"APP NAME", "appname"},
		{"APP-NAME", "appname"},
		{"APP_NAME", "appname"},
		{"APP,NAME", "appname"},
		{"APP#NAME", "appname"},
		{"App NaME ", "appname"},
		{"_App NaME ", "appname"},
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("%s should nomalize to %s", v.input, v.expected), func(t *testing.T) {
			got := NormalizeAppName(v.input)
			if got != v.expected {
				t.Fatalf("want %s, but got %s", v.expected, got)
			}
		})
	}
}

func TestNormalizeAppEnv(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"APP#NAME", "APP_NAME"},
		{"App NaME ", "APP_NAME"},
		{"_App NaME ", "APP_NAME"},
		{"__App NaME ", "APP_NAME"},
		{"App  NaME ", "APP_NAME"},
		{"  App     NaME ", "APP_NAME"},
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("%s should nomalize to %s", v.input, v.expected), func(t *testing.T) {
			got := NormalizeAppEnv(v.input)
			if got != v.expected {
				t.Fatalf("want %s, but got %s", v.expected, got)
			}
		})
	}
}
