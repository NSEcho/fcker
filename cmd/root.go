package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "fcker",
	Short: "Fake person generator",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}
