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
		actionCmd,
		activationCmd,
		packageCmd,
		ruleCmd,
		triggerCmd,
		sdkCmd,
		propertyCmd,
		namespaceCmd,
		listCmd,
	)

	WskCmd.PersistentFlags().BoolVarP(&flags.global.verbose, "verbose", "v", false, "verbose output")
	WskCmd.PersistentFlags().StringVarP(&flags.global.auth, "auth", "u", "", "authorization key")
	WskCmd.PersistentFlags().StringVarP(&flags.global.namespace, "namespace", "n", "", "namespace")
	WskCmd.PersistentFlags().StringVar(&flags.global.apihost, "apihost", "", "whisk API host")
	WskCmd.PersistentFlags().StringVar(&flags.global.apiversion, "apiversion", "", "whisk API version")
}
