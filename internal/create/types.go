package create

const (
	ProjectTypeGoAPI   = "go-api"
	ProjectTypeGoEmbed = "go-embed"
	ProjectTypeGoCLI   = "go-cli"
	ProjectTypeGoGRPC  = "go-grpc"
	ProjectTypeNodeJS  = "nodejs"
	ProjectTypePython  = "python"

	FrameworkGin     = "gin"
	FrameworkFiber   = "fiber"
	FrameworkKoa     = "koa"
	FrameworkNestJS  = "nestjs"
	FrameworkNextJS  = "nextjs"
	FrameworkFastAPI = "fastapi"

	FrontendReact   = "react"
	FrontendVue     = "vue"
	FrontendAngular = "angular"

	ProjectTypeFrontendReact   = "frontend-react"
	ProjectTypeFrontendVue     = "frontend-vue"
	ProjectTypeFrontendAngular = "frontend-angular"

	LanguageGo       = "go"
	LanguageNodeJS   = "nodejs"
	LanguagePython   = "python"
	LanguageFrontend = "frontend"

	SkeletonVersion1 = "v1"
)

type TemplateData struct {
	App      AppData
	Build    BuildData
	Frontend FrontendData
}

type AppData struct {
	Type           string
	Name           string
	Language       string
	Framework      string
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

type FrontendData struct {
	Enabled   bool
	Framework string
	Embedded  bool
}
