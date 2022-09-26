package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version string = "x.x.x"
)

func NewVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show the version of ggen",
		Long:  `This command shows ggen's current version.`,
		Run: func(versionCmd *cobra.Command, args []string) {
			fmt.Printf("ggen version ==> %v\n", Version)
		},
	}

	return versionCmd
}
