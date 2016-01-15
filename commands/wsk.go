package commands

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.ibm.com/Bluemix/whisk-cli/client"
)

var (
	// WskCmd defines the entry point for the cli.
	WskCmd = &cobra.Command{
		Use:              "wsk",
		Short:            "Whisk cloud computing command line interface.",
		Long:             `[TODO] Put "WHISK" in cool ascii font`,
		PersistentPreRun: parseConfigFlags,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Whisk!")
			// Do Stuff Here
		},
	}

	// flags

	whisk *client.Client
)

// Execute runs main whisk command.
func Execute() error {

	// Where to parse flags ??

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
	)

	return WskCmd.Execute()
}

// init sets up flags + loads props + initializes client
func init() {
	clientConfig := &client.Config{}

	// read props
	props, err := readProps(propsFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	if namespace, hasProp := props["NAMESPACE"]; hasProp {
		clientConfig.Namespace = namespace
	}

	if authToken, hasProp := props["AUTH"]; hasProp {
		clientConfig.AuthToken = authToken
	}

	// Setup client
	whisk, err = client.New(http.DefaultClient, clientConfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	WskCmd.Flags().BoolVarP(&flags.edge, "edge", "e", false, "[TODO] test edge server directly, bypassing the master router")
	WskCmd.PersistentFlags().BoolVarP(&flags.verbose, "verbose", "v", false, "[TODO] verbose output")
	WskCmd.PersistentFlags().StringVarP(&flags.auth, "auth", "u", "", "[TODO] authorization key")
	WskCmd.PersistentFlags().StringVarP(&flags.namespace, "namespace", "n", "", "[TODO] namespace")
}

// parseConfigFlags applies command line configs for auth, namespace, and edge.
func parseConfigFlags(cmd *cobra.Command, args []string) {

	if flags.auth != "" {
		whisk.Config.AuthToken = flags.auth
	}

	if flags.namespace != "" {
		whisk.Config.Namespace = flags.namespace
	}

	if flags.edge != false {
		u, err := url.Parse(edgeHost)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		whisk.Config.BaseURL = u
	}

}
