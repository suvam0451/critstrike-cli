// Copyright © 2020 Debashish Patra <suvam0451@outlook.com> -- MPL-2.0

package cmd

import (
	"fmt"
	"os"
	"suvam0451/critstrike/docgen"

	"github.com/spf13/cobra"
)

// docgenSFCmd represents the docgenSF command
var docgenSFCmd = &cobra.Command{
	Use:   "sf",
	Short: "Generate docs for project Sleeping-Forest.",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Generate
		ipath, _ := cmd.Flags().GetString("inputpath")
		opath, _ := cmd.Flags().GetString("outputpath")

		// First generate the intermediate files
		docgen.GenerateIntermediateJSON(ipath, "./tmp_snippets")
		// Then transpile them to .mdx files
		docgen.GenerateSnippetDocs("./tmp_snippets", opath)
		// Delete intermediate build files
		if err := os.RemoveAll("tmp_snippets"); err != nil {
			fmt.Println("Failed to remove intermediate build files")
		}
		// Then signal completion
		fmt.Println("Completed generating documentation for the requested project")
	},
}

func init() {
	docgenCmd.AddCommand(docgenSFCmd)
}
