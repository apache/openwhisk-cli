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
		Use:   "wsk",
		Short: "Whisk cloud computing command line interface.",
		Long:  `[TODO] Put "WHISK" in cool ascii font`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Whisk!")
			// Do Stuff Here
		},
	}

	// flags

	whisk *client.Client
)

const (
	// TODO :: configure this properly
	propsPath = "~/.wskprops"

	testAuthToken = "6c31860c-67ec-4adf-84f5-e421a9d3050e:CShXVzgb0KmlLJ2Iej02p60SBsnZJXA7FCQThVDXLEw2z5faOZBnc9efgp8BuQ9U"
	testNamespace = "_"
	testBaseURL   = "https://whisk.stage1.ng.bluemix.net:443/api/v1/" // TODO :: relplace with local address
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

	// read props
	// props, _ := readProps(propsPath)
	// if err != nil {–≠–
	// 	fmt.Errorf()
	// 	// What now?
	// }

	// Setup client
	u, err := url.Parse(testBaseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	clientConfig := &client.Config{
		AuthToken: testAuthToken,
		Namespace: testNamespace,
		BaseURL:   u,
	}

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
