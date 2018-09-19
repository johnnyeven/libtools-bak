package gen

import (
	"testing"

	"github.com/profzone/libtools/codegen"
)

func TestGen(t *testing.T) {
	clientGenerator := ServiceGenerator{
		ServiceName:  "test",
		DatabaseName: "test",
	}
	codegen.Generate(&clientGenerator)
}
