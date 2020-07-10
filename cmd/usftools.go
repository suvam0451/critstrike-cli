// Copyright © 2020 Debashish Patra <suvam0451@outlook.com> -- MPL-2.0

package cmd

import (
	"fmt"
	"log"
	"os"
	"suvam0451/critstrike/unrealremark"

	"github.com/spf13/cobra"
)

// usftoolsCmd represents the usftools command
var usftoolsCmd = &cobra.Command{
	Use:   "usftools",
	Short: "Parse/lint shader files(.usf/.ush) for ue4 (using python API)",
	Long:  `Although the input file should be .usf/.ush. and output to a python file, you may run it on any file you want`,
	Run: func(cmd *cobra.Command, args []string) {
		// Generate
		ipath, _ := cmd.Flags().GetString("inputpath")
		opath, _ := cmd.Flags().GetString("outputpath")
		pre, _ := cmd.Flags().GetString("alias")

		// Validate file
		if _, err := os.Stat(ipath); os.IsNotExist(err) {
			log.Fatalf("Input file does not exist. Check if -i flag has valid path/to/file.")
		}

		if ipath != "" && opath != "" {
			unrealremark.ParseHlslFile(ipath, opath, pre)
			fmt.Println("Finished parsing the shader file.")
		} else {
			fmt.Println("Input or output specified was invalid.")
		}
	},
}

func init() {
	rootCmd.AddCommand(usftoolsCmd)
	usftoolsCmd.PersistentFlags().StringP("inputpath", "i", "", "usf file to parse.")
	usftoolsCmd.PersistentFlags().StringP("outputpath", "o", "./Output.py", "python file location to output to.")
	usftoolsCmd.PersistentFlags().StringP("alias", "a", "", "Virtual shader source path alias.")
}
