package cmd

import (
	"github.com/spf13/cobra"

	"golib/tools/codegen"
	"golib/tools/courier/swagger/gen"
)

var cmdSwagger = &cobra.Command{
	Use:   "swagger",
	Short: "scan and generate swagger.json",
	Run: func(cmd *cobra.Command, args []string) {
		swaggerGenerator := gen.SwaggerGenerator{
			RootRouterName: "RootRouter",
		}
		codegen.Generate(&swaggerGenerator)
	},
}

func init() {
	cmdRoot.AddCommand(cmdSwagger)
}
