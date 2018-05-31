package loaderx

import (
	"strings"

	"profzone/libtools/codegen"
)

func GetPkgImportPathAndExpose(s string) (pkgImportPath string, expose string) {
	args := strings.Split(s, ".")
	lenOfArgs := len(args)
	if lenOfArgs > 1 {
		return codegen.DeVendor(strings.Join(args[0:lenOfArgs-1], ".")), args[lenOfArgs-1]
	}
	return "", s
}
