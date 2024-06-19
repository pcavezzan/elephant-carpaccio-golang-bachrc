package cmd

import (
	"fmt"

	"github.com/bachrc/test-stonal/internal/app"
	"github.com/spf13/cobra"
)
  
func init() {
	rootCmd.AddCommand(statesCmd)
}

var statesCmd = &cobra.Command{
	Use:   "states",
	Short: "Print the states",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		app := app.New()

		fmt.Println(app.ListStates())
	},
}