package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd defines the entry point for the cli.  All other commands are registered on this.
var (
	RootCmd = &cobra.Command{
		Use:   "wsk",
		Short: "Whisk cloud computing command line interface.",
		Long: `This is a longer description
						of Whisk!`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
)

func init() {
	// TODO :: configure cobra

	// cobra.OnInitialize(initConfig)
	// RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// RootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	// RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// RootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	// RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}
