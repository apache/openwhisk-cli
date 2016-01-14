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
	Short: "A brief description of your command",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("namespace list called")
	},
}

var namespaceSetCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
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
