package commands

import "github.com/spf13/cobra"

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list triggers, actions, and rules in the registry",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		//List returns lists of all actions, triggers, rules, and activations.

		// actions, _, err := wsk.Actions.List(nil)
		// if err != nil {
		// 	return
		// }
		//
		// triggers, _, err := wsk.Triggers.List(nil)
		// if err != nil {
		// 	return
		// }
		//
		// rules, _, err := wsk.Rules.List(nil)
		// if err != nil {
		// 	return
		// }
		//
		// activations, _, err := wsk.Activations.List(nil)
		// if err != nil {
		// 	return
		// }
		//
		// return

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
