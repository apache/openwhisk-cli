package commands

import (
	"os"

	"github.com/spf13/cobra"
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
)

// Execute runs main whisk command.
func Execute() {

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
	// TODO :: configure cobra

	// add commands

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
