package gen

import (
	"os"
	"testing"

	"profzone/libtools/codegen"
)

func init() {
	os.Chdir("..")
}

func TestGen(t *testing.T) {
	clientGenerator := ClientGenerator{
		SpecURL:    "http://service-i-cashdesk.staging.g7pay.net/cashdesk",
		BaseClient: "profzone/libtools/courier/client.Client",
	}
	codegen.Generate(&clientGenerator)
}

func TestGenV3(t *testing.T) {
	clientGenerator := ClientGenerator{
		SpecURL:    "http://service-demo.staging.g7pay.net/demo",
		BaseClient: "profzone/libtools/courier/client.Client",
	}
	codegen.Generate(&clientGenerator)
}
