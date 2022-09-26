package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	version string
)

func NewVersionCommand() *cobra.Command {
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Show the version of ggen",
		Long:  `This command shows ggen's current version.`,
		Run: func(versionCmd *cobra.Command, args []string) {
			if version == "" {
				version = "x.x.x"
			}
			fmt.Printf("ggen version ==> %v\n", version)
		},
	}

	return versionCmd
}
