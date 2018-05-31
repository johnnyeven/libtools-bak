package cmd

import (
	"github.com/spf13/cobra"

	"golib/tools/codegen"
	"golib/tools/courier/status_error/gen"
	"golib/tools/courier/status_error/gen_from_old"
)

var cmdGenError = &cobra.Command{
	Use:   "error",
	Short: "generate general error",
	Run: func(cmd *cobra.Command, args []string) {
		statusErrorGenerator := gen.StatusErrorGenerator{}
		codegen.Generate(&statusErrorGenerator)
	},
}

var cmdGenOldError = &cobra.Command{
	Use:   "old_error",
	Short: "generate general error",
	Run: func(cmd *cobra.Command, args []string) {
		statusErrorGenerator := gen_from_old.StatusErrorGenerator{}
		codegen.Generate(&statusErrorGenerator)
	},
}

func init() {
	cmdGen.AddCommand(cmdGenError)
	cmdGen.AddCommand(cmdGenOldError)
}
