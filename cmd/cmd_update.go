package cmd

import (
	"os/exec"

	"github.com/spf13/cobra"

	"golib/tools/executil"
	"golib/tools/project/govendor"
)

var cmdUpdateFlagUpdate bool

var cmdUpdate = &cobra.Command{
	Use:   "update",
	Short: "go vendor update",
	Run: func(cmd *cobra.Command, args []string) {
		vendor, err := govendor.LoadGoVendorJSON("./vendor/vendor.json")
		if err == nil {
			importPathList := vendor.ListImportPath()
			govendor.UpdatePkgs(importPathList...)
			if cmdUpdateFlagUpdate {
				executil.StdRun(exec.Command("govendor", "update", "+v"))
			}
		}
	},
}

func init() {
	cmdRoot.AddCommand(cmdUpdate)

	cmdUpdate.Flags().
		BoolVarP(&cmdUpdateFlagUpdate, "update", "u", false, "update vendor")
}
