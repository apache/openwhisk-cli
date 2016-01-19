package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print whisk version",

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println(apiVersion)
	},
}
