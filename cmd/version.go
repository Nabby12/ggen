package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of ggen",
	Long:  `This command shows ggen's current version.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ggen version ==> 1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
