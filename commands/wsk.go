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
		authCmd,
		listCmd,
		whoamiCmd,
		triggerCmd,
		actionCmd,
		sdkCmd,
		ruleCmd,
		activationCmd,
		packageCmd,
		healthCmd,
		cleanCmd,
		namespaceCmd,
		versionCmd,

		// hidden
		configCmd,
		propsCmd,
		testCmd,
	)

	WskCmd.Flags().BoolVarP(&flags.edge, "edge", "e", false, "test edge server directly, bypassing the master router")
	WskCmd.PersistentFlags().BoolVarP(&flags.verbose, "verbose", "v", false, "verbose output")
	WskCmd.PersistentFlags().StringVarP(&flags.auth, "auth", "u", "", "authorization key")
	WskCmd.PersistentFlags().StringVarP(&flags.namespace, "namespace", "n", "", "namespace")
}
