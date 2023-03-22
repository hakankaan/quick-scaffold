package cmd

import (
	"github.com/hakankaan/quick-scaffold/pkg/config"
	"github.com/spf13/cobra"
)

var setAPIKeyCmd = &cobra.Command{
	Use:   "setapikey [API_KEY]",
	Short: "Set the API key for the ChatGPT API",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := args[0]
		err := config.SetAPIKey(apiKey)
		if err != nil {
			cmd.PrintErrln("Failed to set the API key:", err)
			return
		}

		cmd.Print("API key set successfully")
	},
}

func init() {
	rootCmd.AddCommand(setAPIKeyCmd)
}
