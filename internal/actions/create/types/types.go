package types

const (
	ProjectTypeCLI  = "cli"
	ProjectTypeAPI  = "api"
	ProjectTypeGRPC = "grpc"
)

type TemplateData struct {
	App   AppData
	Build BuildData
}

type AppData struct {
	Type           string
	Name           string
	Logo           string
	EnvPrefix      string
	ModuleName     string
	DocumentLink   string
	NormalizedName string
}

type BuildData struct {
	Bin  string
	Root string
}
