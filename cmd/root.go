package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "replacertool",
	Version: "1.0.0",
	Short:   "A Text Replacement Tool",
	Long: `Text Replacer Tool version 1.0.0
For help, type replacertool -h | For commands list, type replacertool -l`,
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
