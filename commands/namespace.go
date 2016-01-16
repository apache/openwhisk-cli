package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "identifies the namespace you belong to",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("namespace called")
	},
}

var namespaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all available namespaces",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		namespaces, _, err := whisk.Namespaces.List()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("namespaces")
		for _, namespace := range namespaces {
			fmt.Println(namespace)
		}

	},
}

var namespaceSetCmd = &cobra.Command{
	Use:   "set <namespace string>",
	Short: "sets the namespace to the desired option",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("namespace set called")

		// TODO :: update namespace in props.

	},
}

func init() {
	namespaceCmd.AddCommand(
		namespaceListCmd,
		namespaceSetCmd,
	)
}
