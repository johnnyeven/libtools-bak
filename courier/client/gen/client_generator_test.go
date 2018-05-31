package gen

import (
	"os"
	"testing"

	"golib/tools/codegen"
)

func init() {
	os.Chdir("..")
}

func TestGen(t *testing.T) {
	clientGenerator := ClientGenerator{
		SpecURL:    "http://service-i-cashdesk.staging.g7pay.net/cashdesk",
		BaseClient: "golib/tools/courier/client.Client",
	}
	codegen.Generate(&clientGenerator)
}

func TestGenV3(t *testing.T) {
	clientGenerator := ClientGenerator{
		SpecURL:    "http://service-demo.staging.g7pay.net/demo",
		BaseClient: "golib/tools/courier/client.Client",
	}
	codegen.Generate(&clientGenerator)
}
