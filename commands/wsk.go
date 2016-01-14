package commands

import (
	"net/http"
	"os"

	"github.com/spf13/cobra"
	"github.ibm.com/Bluemix/whisk-cli/client"
)

// WskCmd defines the entry point for the cli.  All other commands are registered on this.
var (
	WskCmd = &cobra.Command{
		Use:   "wsk",
		Short: "Whisk cloud computing command line interface.",
		Long: `This is a longer description
						of Whisk!`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}

	// Top-level flags
	Verbose bool
	// Config string

	whisk *client.Client
)

const (
	// TODO :: configure this properly
	PROPS_PATH = "~/.wskprops"

	TEST_AUTH_TOKEN = "6c31860c-67ec-4adf-84f5-e421a9d3050e:CShXVzgb0KmlLJ2Iej02p60SBsnZJXA7FCQThVDXLEw2z5faOZBnc9efgp8BuQ9U"
	TEST_NAMESPACE  = "wilsonth@us.ibm.com_dev"
	TEST_BASE_URL   = "whisk.stage1.ng.bluemix.net" // TODO :: relplace with local address
)

// Execute runs main whisk command.
func Execute() {

	// Where to parse flags ??

	WskCmd.AddCommand(
		actionCmd,
		activationCmd,
		authCmd,
		cleanCmd,
		listCmd,
		packageCmd,
		ruleCmd,
		sdkCmd,
		triggerCmd,
		versionCmd,
	)

	if err := WskCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}

func init() {

	// read props
	// props, _ := readProps(PROPS_PATH)
	// if err != nil {
	// 	fmt.Errorf()
	// 	// What now?
	// }

	clientConfig := &client.Config{
		AuthToken: TEST_AUTH_TOKEN,
		Namespace: TEST_NAMESPACE,
		BaseURL:   TEST_BASE_URL,
	}
	whisk = client.New(http.DefaultClient, clientConfig)

	// TODO :: configure cobra

	// add common flags

	WskCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

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
