package ui

import "github.com/charmbracelet/lipgloss"

var (
	// SupportedGoFrameworks represents the list of supported Go frameworks.
	SupportedGoFrameworks = map[string][]string{
		"gin":   {"gin", "Gin"},
		"fiber": {"fiber", "Fiber"},
	}

	// SupportedProjectTypes represents the list of supported types.
	SupportedProjectTypes = map[string][]string{
		"api":  {"api", "API"},
		"cli":  {"cli", "CLI"},
		"grpc": {"grpc", "gRPC"},
	}
)

var (
	// ColorBlue represents blue color.
	ColorBlue = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}

	// ColorGreen represents green color.
	ColorGreen = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}

	// ColorYellow represents yellow color.
	ColorYellow = lipgloss.AdaptiveColor{Light: "#E0C900", Dark: "#FFE600"}

	// ColorRed represents red color.
	ColorRed = lipgloss.AdaptiveColor{Light: "#FF4672", Dark: "#ED567A"}

	// ColorGray represents grey color.
	ColorGray = lipgloss.AdaptiveColor{Light: "#FFFDF5", Dark: "#B6B4AD"}
)
