package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "work with namespaces",
}

var namespaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "list available namespaces",

	Run: func(cmd *cobra.Command, args []string) {
		// add "TYPE" --> public / private

		namespaces, _, err := client.Namespaces.List()
		if err != nil {
			fmt.Println(err)
			return
		}
		printList(namespaces)
	},
}

var namespaceGetCmd = &cobra.Command{
	Use:   "get <namespace string>",
	Short: "get triggers, actions, and rules in the registry for a namespace",

	Run: func(cmd *cobra.Command, args []string) {
		var nsName string
		if len(args) > 0 {
			nsName = args[0]
		}

		namespace, _, err := client.Namespaces.Get(nsName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("entities in namespace: %s\n", boldString(namespace.Name))
		printList(namespace.Contents.Packages)
		printList(namespace.Contents.Actions)
		printList(namespace.Contents.Triggers)
		printList(namespace.Contents.Rules)

	},
}

// listCmd is a shortcut for "wsk namespace get _"
var listCmd = &cobra.Command{
	Use:   "list <namespace string>",
	Short: "list triggers, actions, and rules in the registry for a namespace",
	Run:   namespaceGetCmd.Run,
}

func init() {
	namespaceCmd.AddCommand(
		namespaceListCmd,
		namespaceGetCmd,
	)
}
