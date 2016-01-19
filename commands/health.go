package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "general system health",

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("not yet implemented")
	},
}

func init() {}
