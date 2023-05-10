package create

const (
	stageProjFiles = iota + 1
	stageGitIgnore
	stageCIYaml
	stageReleaserYml
	stageSemverYml
	stageMakeRules
	stageInitMod
	stageSwaggerConf
)

type stage struct {
	phase string
	title string
	run   func() error
}

var commonStages = map[int]stage{
	stageProjFiles:   {"1/2", "Creating project files"},
	stageGitIgnore:   {"2/2", "Creating .gitignore"},
	stageCIYaml:      {"2/2", "Creating .golangci.yaml"},
	stageReleaserYml: {"2/2", "Creating .goreleaser.yaml"},
	stageSemverYml:   {"2/2", "Creating .gsemver.yaml"},
	stageMakeRules:   {"2/2", "Creating make rules"},
	stageInitMod:     {"2/2", "Initializing Go module"},
}

var apiStages = map[int]stage{
	stageSwaggerConf: {"2/2", "Creating swagger config files"},
}
