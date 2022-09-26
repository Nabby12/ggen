package root

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/Nabby12/ggen/cmd/handlers/version"
)

type Options struct {
	Count   int
	NoLower bool
	Numeric bool
	Symbol  bool
	Upper   bool
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

		RunE: func(rootCmd *cobra.Command, args []string) error {
			result, err := Generate(
				o.Count,
				o.NoLower,
				o.Numeric,
				o.Symbol,
				o.Upper,
			)
			if err != nil {
				return err
			}

			fmt.Printf("%v\n", result)
			return nil
		},
	}

	// set options variables.
	rootCmd.PersistentFlags().IntVarP(&o.Count, "count", "c", 20, "generated password word count")
	rootCmd.PersistentFlags().BoolVarP(&o.NoLower, "no-lowercase", "", false, "not contain lowercase letters")
	rootCmd.PersistentFlags().BoolVarP(&o.Numeric, "numeric", "n", false, "contain numbers")
	rootCmd.PersistentFlags().BoolVarP(&o.Symbol, "symbol", "s", false, "contain symbols")
	rootCmd.PersistentFlags().BoolVarP(&o.Upper, "uppercase", "u", false, "contain uppercase letters")

	// disale 'completion' command.
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(version.NewVersionCommand())

	return rootCmd
}
