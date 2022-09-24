package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Nabby12/ggen/cmd/handlers"
)

type Options struct {
	Count   int
	NoLower bool
	Numeric bool
	Symbol  bool
	Upper   bool
}

var (
	o = &Options{}
)

var rootCmd = &cobra.Command{
	Use:   "ggen",
	Short: "Generate random password by ggen",
	Long: `'ggen' is a simple CLI tool via Golang.
This application quickly generate random password.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		result, err := handlers.Generate(
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

func Execute() {
	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}

func init() {
	// disale 'completion' command.
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	// set options variables.
	rootCmd.Flags().IntVarP(&o.Count, "count", "c", 10, "generated password word count")
	rootCmd.Flags().BoolVarP(&o.NoLower, "no-lowercase", "", false, "not contain lowercase letters")
	rootCmd.Flags().BoolVarP(&o.Numeric, "numeric", "n", false, "contain numbers")
	rootCmd.Flags().BoolVarP(&o.Symbol, "symbol", "s", false, "contain symbols")
	rootCmd.Flags().BoolVarP(&o.Upper, "uppercase", "u", false, "contain uppercase letters")
}
