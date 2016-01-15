package commands

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list triggers, actions, and rules in the registry",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		listAll := flags.x_type == ""

		if (listAll) || (flags.x_type == "actions") {
			actions, _, err := whisk.Actions.List(nil)
			if err != nil {
				return
			}
			fmt.Println("actions")
			spew.Dump(actions)
		}

		if (listAll) || (flags.x_type == "triggers") {

			triggers, _, err := whisk.Triggers.List(nil)
			if err != nil {
				return
			}
			fmt.Println("triggers")
			spew.Dump(triggers)
		}

		if (listAll) || (flags.x_type == "rules") {

			rules, _, err := whisk.Rules.List(nil)
			if err != nil {
				return
			}
			fmt.Println("rules")
			spew.Dump(rules)
		}

		if (listAll) || (flags.x_type == "activations") {

			activations, _, err := whisk.Activations.List(nil)
			if err != nil {
				return
			}
			fmt.Println("activations")
			spew.Dump(activations)
		}

	},
}

func init() {

	listCmd.Flags().StringVarP(&flags.x_type, "type", "t", "", "only list given type")
	listCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	listCmd.Flags().IntVarP(&flags.limit, "limit", "l", 0, "only return this many entities from the collection")

}
