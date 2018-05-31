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
	statusErrorGenerator := StatusErrorGenerator{}
	codegen.Generate(&statusErrorGenerator)
}
