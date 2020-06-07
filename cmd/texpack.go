// Copyright © 2020 Debashish Patra <suvam0451@outlook.com> -- MPL-2.0

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// texpackCmd represents the texpack command
var texpackCmd = &cobra.Command{
	Use:   "texpack",
	Short: "Packs textures from same folder into single texture/atlas",
	Long: `Texture packing is a textureset optimization/management feature
Multiple single channel masks are better off as one texture to save space/memory and for future convenience.
The original textures will remain unaffected
	Channel Packing : Used to generate RMA, MRA maps`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("texpack called")
	},
}

func init() {
	rootCmd.AddCommand(texpackCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// texpackCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// texpackCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
