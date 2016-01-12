package commands

import (
	"github.com/spf13/cobra"
)

var WhiskCmd = &cobra.Command{
	Use:   "wsk",
	Short: "Whisk cloud computing command line interface.",
	Long: `This is a longer description
						of Whisk!`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}
