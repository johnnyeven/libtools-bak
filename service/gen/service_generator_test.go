package gen

import (
	"testing"

	"golib/tools/codegen"
)

func TestGen(t *testing.T) {
	clientGenerator := ServiceGenerator{
		ServiceName:  "test",
		DatabaseName: "test",
	}
	codegen.Generate(&clientGenerator)
}
