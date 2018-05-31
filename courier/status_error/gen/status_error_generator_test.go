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
	statusErrorGenerator := StatusErrorGenerator{}
	codegen.Generate(&statusErrorGenerator)
}
