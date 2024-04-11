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

func TestNormalizeAppLogo(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"jaguar", "   _\n  (_) __ _  __ _  _  _  __ _  _ _\n  | |/ _` |/ _` || || |/ _` || '_|\n _/ |\\__,_|\\__, | \\_,_|\\__,_||_|\n|__/       |___/"},
		{"Jaguar", "   _\n  (_) __ _  __ _  _  _  __ _  _ _\n  | |/ _` |/ _` || || |/ _` || '_|\n _/ |\\__,_|\\__, | \\_,_|\\__,_||_|\n|__/       |___/"},
		{"Jag uar", "   _\n  (_) __ _  __ _    _  _  __ _  _ _\n  | |/ _` |/ _` |  | || |/ _` || '_|\n _/ |\\__,_|\\__, |   \\_,_|\\__,_||_|\n|__/       |___/"},
		{"Jag uar ", "   _\n  (_) __ _  __ _    _  _  __ _  _ _\n  | |/ _` |/ _` |  | || |/ _` || '_|\n _/ |\\__,_|\\__, |   \\_,_|\\__,_||_|\n|__/       |___/"},
		{"Jag-uar", "   _\n  (_) __ _  __ _  ___  _  _  __ _  _ _\n  | |/ _` |/ _` ||___|| || |/ _` || '_|\n _/ |\\__,_|\\__, |      \\_,_|\\__,_||_|\n|__/       |___/"},
		{".Jag-uar?", "      _                                    ___\n     (_) __ _  __ _  ___  _  _  __ _  _ _ |__ \\\n _   | |/ _` |/ _` ||___|| || |/ _` || '_|  /_/\n(_) _/ |\\__,_|\\__, |      \\_,_|\\__,_||_|   (_)\n   |__/       |___/"},
		{"Jag]uar", "   _\n  (_) __ _  __ _    _  _  __ _  _ _\n  | |/ _` |/ _` |  | || |/ _` || '_|\n _/ |\\__,_|\\__, |   \\_,_|\\__,_||_|\n|__/       |___/"},
		{"Jag#uar", "   _\n  (_) __ _  __ _    _  _  __ _  _ _\n  | |/ _` |/ _` |  | || |/ _` || '_|\n _/ |\\__,_|\\__, |   \\_,_|\\__,_||_|\n|__/       |___/"},
		{"%Jag#uar&", "   _\n  (_) __ _  __ _    _  _  __ _  _ _\n  | |/ _` |/ _` |  | || |/ _` || '_|\n _/ |\\__,_|\\__, |   \\_,_|\\__,_||_|\n|__/       |___/"},
	}

	for _, v := range tests {
		t.Run(fmt.Sprintf("%s should nomalize success", v.input), func(t *testing.T) {
			got := NormalizeAppLogo(v.input)
			if got != v.expected {
				t.Fatalf("want %s, but got %s", v.expected, got)
			}
		})
	}
}
