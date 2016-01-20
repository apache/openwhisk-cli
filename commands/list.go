package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list triggers, actions, and rules in the registry",

	Run: func(cmd *cobra.Command, args []string) {
		listAll := flags.xType == ""

		if (listAll) || (flags.xType == "actions") {
			actions, _, err := client.Actions.List(nil)
			if err != nil {
				return
			}
			fmt.Println("actions")
			printJSON(actions)
		}

		if (listAll) || (flags.xType == "triggers") {

			triggers, _, err := client.Triggers.List(nil)
			if err != nil {
				return
			}
			fmt.Println("triggers")
			printJSON(triggers)
		}

		if (listAll) || (flags.xType == "rules") {

			rules, _, err := client.Rules.List(nil)
			if err != nil {
				return
			}
			fmt.Println("rules")
			printJSON(rules)
		}

		if (listAll) || (flags.xType == "activations") {

			activations, _, err := client.Activations.List(nil)
			if err != nil {
				return
			}
			fmt.Println("activations")
			printJSON(activations)
		}

	},
}

func init() {

	listCmd.Flags().StringVarP(&flags.xType, "type", "t", "", "only list given type")
	listCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	listCmd.Flags().IntVarP(&flags.limit, "limit", "l", 0, "only return this many entities from the collection")

}
