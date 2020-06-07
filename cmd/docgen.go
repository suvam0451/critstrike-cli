// Copyright © 2020 Debashish Patra <suvam0451@outlook.com> -- MPL-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// docgenCmd represents the docgen command
var docgenCmd = &cobra.Command{
	Use:   "docgen",
	Short: "Internal documentation generator for a project template",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("no project template provided. See help for list")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(docgenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// docgenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// docgenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// option to modify input and output parameters must be inherited
	docgenCmd.PersistentFlags().StringP("inputpath", "i", "./snippets", "location of input .json files")
	docgenCmd.PersistentFlags().StringP("outputpath", "o",
		"./content/sleeping-forest/02-Snippets/02-Full-list",
		"location to output .mdx files to")
}
