package cmd

import (
	"github.com/spf13/cobra"
)

// WhiskCmd defines the entry point for the cli.  All other commands are registered on this.
var (
	WhiskCmd = &cobra.Command{
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
	// WhiskCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// WhiskCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	// WhiskCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	// WhiskCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	// WhiskCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	// viper.BindPFlag("author", WhiskCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("projectbase", WhiskCmd.PersistentFlags().Lookup("projectbase"))
	// viper.BindPFlag("useViper", WhiskCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}
