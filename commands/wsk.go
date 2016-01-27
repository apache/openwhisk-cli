package commands

import "github.com/spf13/cobra"

// WskCmd defines the entry point for the cli.
var WskCmd = &cobra.Command{
	Use:              "wsk",
	Short:            "Whisk cloud computing command line interface.",
	Long:             logoText(),
	PersistentPreRun: parseConfigFlags,
}

func init() {

	WskCmd.AddCommand(
		triggerCmd,
		actionCmd,
		sdkCmd,
		ruleCmd,
		activationCmd,
		packageCmd,
	)

	WskCmd.Flags().BoolVarP(&flags.global.edge, "edge", "e", false, "test edge server directly, bypassing the master router")
	WskCmd.PersistentFlags().BoolVarP(&flags.global.verbose, "verbose", "v", false, "verbose output")
	WskCmd.PersistentFlags().StringVarP(&flags.global.auth, "auth", "u", "", "authorization key")
	WskCmd.PersistentFlags().StringVarP(&flags.global.namespace, "namespace", "n", "", "namespace")
}
