package root

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Nabby12/ggen/cmd/handlers/version"
)

type Options struct {
	Count       int
	NoLowerCase bool
	NoNumeric   bool
	NoSymbol    bool
	NoUpperCase bool
}

var (
	o *Options = &Options{}
)

func NewRootCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ggen",
		Short: "Generate random password by ggen",
		Long: `'ggen' is a simple CLI tool via Golang.
This application quickly generate random password.`,
		Example: `  ggen                   # 30 characters containing lowercase, uppercase, numeric and symbol.
  ggen -c 15 --no-symbol # 15 characters not containing symbol.`,

		RunE: func(rootCmd *cobra.Command, args []string) error {
			result, err := Generate(
				o.Count,
				o.NoLowerCase,
				o.NoNumeric,
				o.NoSymbol,
				o.NoUpperCase,
			)
			if err != nil {
				return err
			}

			fmt.Printf("%v\n", result)
			return nil
		},
	}

	// set options variables.
	rootCmd.PersistentFlags().IntVarP(&o.Count, "count", "c", 30, "specify the number of characters")
	rootCmd.PersistentFlags().BoolVarP(&o.NoLowerCase, "no-lowercase", "", false, "not contain lowercase character")
	rootCmd.PersistentFlags().BoolVarP(&o.NoNumeric, "no-numeric", "", false, "not contain numeric character")
	rootCmd.PersistentFlags().BoolVarP(&o.NoSymbol, "no-symbol", "", false, "not contain symbol character")
	rootCmd.PersistentFlags().BoolVarP(&o.NoUpperCase, "no-uppercase", "", false, "not contain uppercase character")

	// disale 'completion' command.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(version.NewVersionCommand())

	return rootCmd
}
