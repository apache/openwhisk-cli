package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// WskCmd defines the entry point for the cli.
var WskCmd = &cobra.Command{
	Use:              "wsk",
	Short:            "Whisk cloud computing command line interface.",
	Long:             `[TODO] Put "WHISK" in cool ascii font`,
	PersistentPreRun: parseConfigFlags,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Whisk!")
		// Do Stuff Here
	},
}

func init() {
	WskCmd.Flags().BoolVarP(&flags.edge, "edge", "e", false, "[TODO] test edge server directly, bypassing the master router")
	WskCmd.PersistentFlags().BoolVarP(&flags.verbose, "verbose", "v", false, "[TODO] verbose output")
	WskCmd.PersistentFlags().StringVarP(&flags.auth, "auth", "u", "", "[TODO] authorization key")
	WskCmd.PersistentFlags().StringVarP(&flags.namespace, "namespace", "n", "", "[TODO] namespace")
}
