package commands

import "github.com/spf13/cobra"

var testCmd = &cobra.Command{
	Hidden: true,
	Use:    "test",
	Run: func(cmd *cobra.Command, args []string) {
		printJSON(args)
		printJSON(flags.param)
	},
}

func init() {
	testCmd.Flags().StringVarP(&flags.param, "param", "p", "", "parameters")
}
