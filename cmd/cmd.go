package cmd

import (
	"os"

	"github.com/Nabby12/ggen/cmd/handlers/root"
)

func Execute() {
	rootCmd := root.NewRootCommand()

	rootCmd.SetOutput(os.Stdout)
	if err := rootCmd.Execute(); err != nil {
		rootCmd.SetOutput(os.Stderr)
		rootCmd.Println(err)
		os.Exit(1)
	}
}
