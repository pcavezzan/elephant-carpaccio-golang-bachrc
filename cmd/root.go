package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "elephant",
	Short: "Elephant is fun",
	Long: `Pouet pouet`,
	Run: func(cmd *cobra.Command, args []string) {
	  // Do Stuff Here
	},
}
  
func Execute() error {
	return rootCmd.Execute()
}