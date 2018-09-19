package common

import (
	"io"

	"github.com/profzone/libtools/codegen"
)

type Op interface {
	ID() string
	Method() string
	Path() string
	HasRequest() bool
	WriteReqType(w io.Writer, importer *codegen.Importer)
	WriteRespBodyType(w io.Writer, importer *codegen.Importer)
}
