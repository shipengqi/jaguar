package ui

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
