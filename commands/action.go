package commands

import (
	"errors"
	"fmt"

	"github.ibm.com/Bluemix/whisk-cli/client"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

//////////////
// Commands //
//////////////

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "work with actions",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("action called")
	},
}

var actionCreateCmd = &cobra.Command{
	Use:   "create <name string> <artifact string>",
	Short: "create a new action",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		actionName := args[0]
		// artifactName := args[1]

		// flags.docker
		// flags.copy
		// flags.pipe
		// flags.lib
		// flags.package
		// flags.param
		// flags.annotation

		exec := client.Exec{}
		annotations := client.Annotations{}
		parameters := client.Parameters{}
		limits := client.Limits{
			Timeout: flags.timeout,
			Memory:  flags.memory,
		}

		action := &client.Action{
			Name:        actionName,
			Publish:     flags.shared,
			Exec:        exec,
			Annotations: annotations,
			Parameters:  parameters,
			Limits:      limits,
		}

		action, _, err = whisk.Actions.Insert(action, false)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: created action")
		spew.Dump(action)

	},
}

var actionUpdateCmd = &cobra.Command{
	Use:   "update []",
	Short: "update an existing action",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		actionName := args[0]
		// artifactName := args[1]

		// flags.docker
		// flags.copy
		// flags.pipe
		// flags.lib
		// flags.package
		// flags.param
		// flags.annotation

		exec := client.Exec{}
		annotations := client.Annotations{}
		parameters := client.Parameters{}
		limits := client.Limits{
			Timeout: flags.timeout,
			Memory:  flags.memory,
		}

		action := &client.Action{
			Name:        actionName,
			Publish:     flags.shared,
			Exec:        exec,
			Annotations: annotations,
			Parameters:  parameters,
			Limits:      limits,
		}

		action, _, err = whisk.Actions.Insert(action, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: created action")
		spew.Dump(action)

	},
}

var actionInvokeCmd = &cobra.Command{
	Use:     "invoke <name string>",
	Short:   "invoke action",
	Long:    `[ TODO :: add longer description here ]`,
	Example: "invoke action --json --blocking -p key_1,val_1 -p key_2,val_2 action_name",
	Run: func(cmd *cobra.Command, args []string) {

		// params, _ := cmd.Flags().GetStringSlice("param")
		//
		// spew.Dump(params)
		// // TODO :: parse params into K|V pairs

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}
		actionName := args[0]

		activation, _, err := whisk.Actions.Invoke(actionName, flags.blocking)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		// print out response
		fmt.Printf("ok: invoked %s with id %s\n", actionName, activation.ActivationID)
		spew.Dump(activation)
	},
}

var actionGetCmd = &cobra.Command{
	Use:   "get <name string>",
	Short: "get action",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		actionName := args[0]
		action, _, err := whisk.Actions.Fetch(actionName)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		// print out response
		fmt.Printf("ok: got action %s\n", actionName)
		spew.Dump(action)
	},
}

var actionDeleteCmd = &cobra.Command{
	Use:   "delete <name string>",
	Short: "delete action",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		actionName := args[0]
		_, err := whisk.Actions.Delete(actionName)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		// print out response
		fmt.Printf("ok: deleted action %s\n", actionName)
	},
}

var actionListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all actions",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		options := &client.ActionListOptions{
			Skip:  flags.skip,
			Limit: flags.limit,
		}
		actions, _, err := whisk.Actions.List(options)
		if err != nil {
			fmt.Printf("error: %s", err)
			return
		}
		fmt.Printf("actions\n")
		for _, action := range actions {
			var publishState string
			if action.Publish {
				publishState = "public"
			} else {
				publishState = "private"
			}

			fmt.Printf("%s\t\t\t\t%s", action.Name, publishState)
		}

	},
}

///////////
// Flags //
///////////

func init() {

	actionCreateCmd.Flags().BoolVar(&flags.docker, "docker", false, "treat artifact as docker image path on dockerhub")
	actionCreateCmd.Flags().BoolVar(&flags.copy, "copy", false, "treat artifact as the name of an existing action")
	actionCreateCmd.Flags().BoolVar(&flags.pipe, "pipe", false, "pipe treat artifact as comma separated sequence of actions to invoke")
	actionCreateCmd.Flags().BoolVar(&flags.shared, "shared", false, "add library to artifact (must be a gzipped tar file)")
	actionCreateCmd.Flags().StringVar(&flags.lib, "lib", "", "add library to artifact (must be a gzipped tar file)")
	actionCreateCmd.Flags().StringVar(&flags.xPackage, "package", "", "package")

	// actionUpdateCmd

	actionInvokeCmd.Flags().BoolP("json", "j", false, "output as JSON")
	actionInvokeCmd.Flags().StringSliceP("param", "p", []string{}, "parameters")
	actionInvokeCmd.Flags().BoolP("blocking", "b", false, "blocking invoke")

	// actionGetCmd.
	// actionDeleteCmd,

	actionCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entitites from the head of the collection")
	actionCmd.Flags().IntVarP(&flags.limit, "limit", "l", 30, "only return this many entities from the collection")
	actionCmd.Flags().BoolVar(&flags.full, "full", false, "include full entity description")

	actionCmd.AddCommand(
		actionCreateCmd,
		actionUpdateCmd,
		actionInvokeCmd,
		actionGetCmd,
		actionDeleteCmd,
		actionListCmd,
	)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// actionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// actionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
