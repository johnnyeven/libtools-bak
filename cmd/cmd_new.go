package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"profzone/libtools/codegen"
	"profzone/libtools/service/gen"
)

var cmdNewFlagName string

func init() {
	cmdRoot.AddCommand(cmdNew)

	cmdNew.Flags().
		StringVarP(&cmdNewFlagName, "db-name", "", "", "with db name")

}

var cmdNew = &cobra.Command{
	Use:   "new",
	Short: "new service",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic(fmt.Errorf("need service name"))
		}

		clientGenerator := gen.ServiceGenerator{
			ServiceName:  args[0],
			DatabaseName: cmdNewFlagName,
		}

		codegen.Generate(&clientGenerator)
	},
}
