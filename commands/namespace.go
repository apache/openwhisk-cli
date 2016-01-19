package commands

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "identifies the namespace you belong to",
}

var namespaceListCmd = &cobra.Command{
	Use:   "list",
	Short: "lists all available namespaces",

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

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid namespace argument")
			fmt.Println(err)
			return
		}

		namespace := args[0]

		props, _ := readProps(PropsFile)
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }
		props["NAMESPACE"] = namespace

		err := writeProps(PropsFile, props)
		if err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	namespaceCmd.AddCommand(
		namespaceListCmd,
		namespaceSetCmd,
	)
}
