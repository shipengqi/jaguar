package version

import (
	"fmt"
	"runtime"
)

// Build-time variables set via ldflags.
var (
	Version      = "v0.0.0-dev"
	GitCommit    = "unknown"
	GitTreeState = "unknown"
	BuildTime    = "unknown"
)

// Info holds version information.
type Info struct {
	Version      string `json:"version"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildTime    string `json:"buildTime"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

// String returns a human-readable version string.
func (i Info) String() string {
	return fmt.Sprintf("Version: %s\nCommit: %s\nGitTreeState: %s\nBuildTime: %s\nGoVersion: %s\nCompiler: %s\nPlatform: %s",
		i.Version, i.GitCommit, i.GitTreeState, i.BuildTime, i.GoVersion, i.Compiler, i.Platform)
}

// Get returns the current build's version info.
func Get() Info {
	return Info{
		Version:      Version,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildTime:    BuildTime,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
