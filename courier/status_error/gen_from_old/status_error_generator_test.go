package gen_from_old

import (
	"os"
	"testing"

	"profzone/libtools/codegen"
)

func init() {
	os.Chdir("./fixtures")
}

func TestGen(t *testing.T) {
	statusErrorGenerator := StatusErrorGenerator{
		DryRun: true,
	}
	codegen.Generate(&statusErrorGenerator)
}
