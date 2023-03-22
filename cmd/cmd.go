package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A CLI app for interacting with ChatGPT API",
}

// Execute runs the root command
func Execute() error {
	return rootCmd.Execute()
}
