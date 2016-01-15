package commands

import (
	"errors"
	"fmt"

	"github.ibm.com/Bluemix/whisk-cli/client"

	"github.com/davecgh/go-spew/spew"
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
	Short: "list activations",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		options := &client.ActivationListOptions{
			Limit: flags.limit,
			Skip:  flags.skip,
			Upto:  flags.upto,
			Docs:  flags.full,
		}
		activations, _, err := whisk.Activations.List(options)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("activations")
		spew.Dump(activations)
	},
}

var activationGetCmd = &cobra.Command{
	Use:   "get <id string>",
	Short: "get activation",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid ID argument")
			fmt.Println(err)
			return
		}
		id := args[0]
		activation, _, err := whisk.Activations.Fetch(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: got activation ", id)
		spew.Dump(activation)

	},
}

var activationLogsCmd = &cobra.Command{
	Use:   "logs",
	Short: "get the logs of an activation",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid ID argument")
			fmt.Println(err)
			return
		}

		id := args[0]
		activation, _, err := whisk.Activations.Logs(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: got activation logs")
		spew.Dump(activation)
	},
}

var activationResultCmd = &cobra.Command{
	Use:   "result",
	Short: "get the result of an activation",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			err := errors.New("Invalid ID argument")
			fmt.Println(err)
			return
		}

		id := args[0]
		result, _, err := whisk.Activations.Result(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: got activation result")
		spew.Dump(result)
	},
}

var activationPollCmd = &cobra.Command{
	Use:   "poll",
	Short: "poll continuously for log messages from currently running actions",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("TODO :: implement activationPollCmd")
	},
}

func init() {

	activationListCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entitites from the head of the collection")
	activationListCmd.Flags().IntVarP(&flags.limit, "limit", "l", 30, "only return this many entities from the collection")
	activationListCmd.Flags().BoolVarP(&flags.full, "full", "f", false, "include full entity description")
	activationListCmd.Flags().IntVar(&flags.upto, "upto", 0, "return activations with timestamps earlier than UPTO; measured in miliseconds since Th, 01, Jan 1970")
	activationListCmd.Flags().IntVar(&flags.since, "since", 0, "return activations with timestamps earlier than UPTO; measured in miliseconds since Th, 01, Jan 1970")

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
