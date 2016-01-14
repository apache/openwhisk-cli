package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

// activationCmd represents the activation command
var activationCmd = &cobra.Command{
	Use:   "activation",
	Short: "work with activations",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("activation called")
	},
}

var activationListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("activation list called")
	},
}

var activationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("activation get called")
	},
}

var activationLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "A brief description of your command",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("activation logs called")
	},
}

var activationResultCmd = &cobra.Command{
	Use:   "result",
	Short: "A brief description of your command",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("activation called")
	},
}

var activationPollCmd = &cobra.Command{
	Use:   "poll",
	Short: "A brief description of your command",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("activation called")
	},
}

func init() {

	activationCmd.AddCommand(
		activationListCmd,
		activationGetCmd,
		activationLogsCmd,
		activationResultCmd,
		activationPollCmd,
	)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// activationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// activationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
