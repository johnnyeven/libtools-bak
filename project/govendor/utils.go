package govendor

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path"
	"strings"

	"golib/tools/executil"
)

var pkgsNoGoGet = []string{
	"g7pay/",
	"devops/",
	"golib/",
}

func isLocalPkg(pkgName string) bool {
	for _, pkgPrefix := range pkgsNoGoGet {
		if strings.HasPrefix(pkgName, pkgPrefix) {
			return true
		}
	}
	return false
}

func UpdatePkg(importPath string) {
	if !isLocalPkg(importPath) {
		executil.StdRun(exec.Command("go", "get", "-u", "-v", importPath))
		return
	}

	pkg, err := build.Import(importPath, "", build.FindOnly)
	if err != nil {
		goPath := os.Getenv("GOPATH")
		os.Chdir(path.Join(strings.Split(goPath, ":")[0], "src"))

		gitRepo := fmt.Sprintf(`git@git.chinawayltd.com:%s.git`, importPath)
		executil.StdRun(exec.Command("git", "clone", gitRepo, importPath))
	} else {
		os.Chdir(pkg.Dir)
		executil.StdRun(exec.Command("git", "pull", "--rebase"))
	}
}
