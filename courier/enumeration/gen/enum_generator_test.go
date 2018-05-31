package gen

import (
	"os"
	"testing"

	"golib/tools/codegen"
)

func init() {
	os.Chdir("./types")
}

func TestGen(t *testing.T) {
	enumGenerator := EnumGenerator{}
	codegen.Generate(&enumGenerator)
}
