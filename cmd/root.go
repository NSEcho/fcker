package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "fcker",
	Short: "Fetcher/manipulator of fake data",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}
