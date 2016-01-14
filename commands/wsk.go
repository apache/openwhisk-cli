package commands

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/spf13/cobra"
	"github.ibm.com/Bluemix/whisk-cli/client"
)

// WskCmd defines the entry point for the cli.  All other commands are registered on this.
var (
	WskCmd = &cobra.Command{
		Use:   "wsk",
		Short: "Whisk cloud computing command line interface.",
		Long:  `Whisk cloud computing command line interface.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Print("Whisk!")
			// Do Stuff Here
		},
	}

	// Top-level flags
	Verbose bool
	Edge    bool

	whisk *client.Client
)

const (
	// TODO :: configure this properly
	propsPath = "~/.wskprops"

	testAuthToken = "6c31860c-67ec-4adf-84f5-e421a9d3050e:CShXVzgb0KmlLJ2Iej02p60SBsnZJXA7FCQThVDXLEw2z5faOZBnc9efgp8BuQ9U"
	testNamespace = "_"
	testBaseURL   = "whisk.stage1.ng.bluemix.net" // TODO :: relplace with local address
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
	// if err != nil {
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

	WskCmd.Flags().BoolVarP(&Verbose, "verbose", "v", false, "[TODO] verbose output")
	WskCmd.Flags().BoolVarP(&Edge, "edge", "e", false, "[TODO] test edge server directly, bypassing the master router")

	// Some Examples ...

	// cobra.OnInitialize(initConfig)
	// WskCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// WskCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	// WskCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// WskCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	// WskCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// viper.BindPFlag("author", WskCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", WskCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", WskCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}
