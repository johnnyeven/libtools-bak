package gen

import (
	"os"
	"testing"

	"profzone/libtools/codegen"
)

func init() {
	os.Chdir("./types")
}

func TestGen(t *testing.T) {
	enumGenerator := EnumGenerator{}
	codegen.Generate(&enumGenerator)
}
