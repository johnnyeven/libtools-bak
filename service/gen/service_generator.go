package gen

import (
	"go/build"
	"path"

	"golib/tools/codegen"
)

type ServiceGenerator struct {
	ServiceName  string
	DatabaseName string
	Root         string
}

func (g *ServiceGenerator) Load(cwd string) {
}

func (g *ServiceGenerator) Pick() {
}

func (g *ServiceGenerator) Output(cwd string) codegen.Outputs {
	outputs := codegen.Outputs{}

	codegen.NewGenFile("main", path.Join(g.ServiceName, "doc.go")).
		WithData(g).
		OutputTo(outputs)

	outputs.WriteFiles()

	pkg, _ := build.ImportDir(path.Join(cwd, g.ServiceName), build.ImportComment)

	g.Root = pkg.ImportPath

	if g.DatabaseName != "" {
		codegen.NewGenFile("database", path.Join(g.ServiceName, "database/db.go")).
			WithData(g).
			Block(`
		var DB{{ .ExposeVar .Data.DatabaseName }} = {{ ( .PureUse "golib/tools/sqlx" )}}.NewDatabase("{{ .Data.DatabaseName }}")
`,
			).
			OutputTo(outputs)

		outputs.WriteFiles()
	}

	codegen.NewGenFile("global", path.Join(g.ServiceName, "global/config.go")).
		WithData(g).
		Block(`
func init() {
	{{ .PureUse "golib/tools/servicex" }}.SetServiceName("{{ .Data.ServiceName }}")
	{{ .PureUse "golib/tools/servicex" }}.ConfP(&Config)

	{{ if .Data.DatabaseName }}
		{{ .PureUse .Data.Root "database" }}.DB{{ .ExposeVar .Data.DatabaseName }}.MustMigrateTo(Config.MasterDB.Get(), !{{ .PureUse "golib/tools/servicex" }}.AutoMigrate)
	{{ end }}
}

var Config = struct {
	Log      *{{ ( .PureUse "golib/tools/log" ) }}.Log
	Server   {{ ( .PureUse "golib/tools/courier/transport_http" ) }}.ServeHTTP
{{ if .Data.DatabaseName }}
	MasterDB *{{ .PureUse "golib/tools/sqlx/mysql" }}.MySQL
	SlaveDB  *{{ .PureUse "golib/tools/sqlx/mysql" }}.MySQL
{{ end }}
}{
	Log: &{{ ( .PureUse "golib/tools/log" ) }}.Log{
		Level: "DEBUG",
	},
	Server: {{ ( .PureUse "golib/tools/courier/transport_http" ) }}.ServeHTTP{
		WithCORS: true,
		Port:     8000,
	},
{{ if .Data.DatabaseName }}
	MasterDB: &{{ .PureUse "golib/tools/sqlx/mysql" }}.MySQL{
		Name: "{{ .Data.DatabaseName }}",
		Port: 33306,
		User: "root",
		Password: "root",
		Host: "....",
	},
	SlaveDB: &{{ .PureUse "golib/tools/sqlx/mysql" }}.MySQL{
		Name: "{{ .Data.DatabaseName }}-readonly",
		Port: 33306,
		User: "root",
		Password: "root",
		Host: "....",
	},
{{ end }}
}
`,
		).OutputTo(outputs)

	codegen.NewGenFile("types", path.Join(g.ServiceName, "constants/types/doc.go")).WithData(g).Block(`
// Defined enum types here
	`).OutputTo(outputs)

	codegen.NewGenFile("modules", path.Join(g.ServiceName, "modules/doc.go")).WithData(g).Block(`
// Defined sub modules here
	`).OutputTo(outputs)

	codegen.NewGenFile("errors", path.Join(g.ServiceName, "constants/errors/status_err_codes.go")).
		WithData(g).
		Block(`
//go:generate tools gen error
const ServiceStatusErrorCode = 0 * 1e3 // todo rename this

const (
	// 请求参数错误
	BadRequest {{ .PureUse "golib/tools/courier/status_error" }}.StatusErrorCode = http.StatusBadRequest*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 未找到
	NotFound {{ .PureUse "golib/tools/courier/status_error" }}.StatusErrorCode = http.StatusNotFound*1e6 + ServiceStatusErrorCode + iota
)

const (
	// @errTalk 未授权
	Unauthorized {{ .PureUse "golib/tools/courier/status_error" }}.StatusErrorCode = http.StatusUnauthorized*1e6 + ServiceStatusErrorCode + iota
)

const (
	// @errTalk 操作冲突
	Conflict {{ .PureUse "golib/tools/courier/status_error" }}.StatusErrorCode = http.StatusConflict*1e6 + ServiceStatusErrorCode + iota
)

const (
	// @errTalk 不允许操作
	Forbidden {{ .PureUse "golib/tools/courier/status_error" }}.StatusErrorCode = http.StatusForbidden*1e6 + ServiceStatusErrorCode + iota
)

const (
	// 内部处理错误
	InternalError {{ .PureUse "golib/tools/courier/status_error" }}.StatusErrorCode = http.StatusInternalServerError*1e6 + ServiceStatusErrorCode + iota
)
		`,
		).
		OutputTo(outputs)

	codegen.NewGenFile("routes", path.Join(g.ServiceName, "routes/root.go")).
		WithData(g).
		Block(`
var RootRouter = {{ .PureUse "golib/tools/courier" }}.NewRouter(GroupRoot{})

func init() {
	RootRouter.Register({{ .PureUse "golib/tools/courier/swagger" }}.SwaggerRouter)
}

type GroupRoot struct {
	courier.EmptyOperator
}

func (root GroupRoot) Path() string {
	return "/{{ .Data.ServiceName }}"
}
`,
		).
		OutputTo(outputs)

	outputs.WriteFiles()

	codegen.NewGenFile("main", path.Join(g.ServiceName, "main.go")).
		WithData(g).
		Block(`
	func main() {
		{{( .PureUse "golib/tools/servicex" )}}.Execute()
		{{( .PureUse .Data.Root "global" )}}.Config.Server.Serve({{ ( .PureUse .Data.Root "routes" ) }}.RootRouter)
	}
	`,
		).
		OutputTo(outputs)

	return outputs
}
