package commands

import (
	"errors"
	"fmt"

	"github.ibm.com/Bluemix/whisk-cli/client"

	"github.com/spf13/cobra"
)

// triggerCmd represents the trigger command
var triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "work with triggers",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("trigger called")
	},
}

var triggerFireCmd = &cobra.Command{
	Use:   "fire <name string> <payload ?>",
	Short: "fire trigger event",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO :: parse payload from args... how ?
		// whisk.Triggers.Fire(triggerName, payload)
	},
}

var triggerCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create new trigger",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO :: parse annotation
		// TODO :: parse param

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		triggerName := args[0]

		trigger := &client.Trigger{
			Name: triggerName,
			// Param
			// Annotation
		}

		trigger, _, err = whisk.Triggers.Insert(trigger, false)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: created trigger")

		printJSON(trigger)
	},
}

var triggerUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update existing trigger",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {

		// TODO :: parse annotation
		// TODO :: parse param

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		triggerName := args[0]

		trigger := &client.Trigger{
			Name: triggerName,
			// Param
			// Annotation
		}

		trigger, _, err = whisk.Triggers.Insert(trigger, true)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: updated trigger")

		printJSON(trigger)
	},
}

var triggerGetCmd = &cobra.Command{
	Use:   "get",
	Short: "get trigger",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		triggerName := args[0]

		trigger, _, err := whisk.Triggers.Fetch(triggerName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: got trigger ", triggerName)
		printJSON(trigger)
	},
}

var triggerDeleteCmd = &cobra.Command{
	Use:   "delete <name string>",
	Short: "delete trigger",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, err = whisk.Triggers.Delete(ruleName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: deleted rule ", ruleName)
	},
}

var triggerListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all triggers",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		options := &client.TriggerListOptions{
			Skip:  flags.skip,
			Limit: flags.limit,
		}
		triggers, _, err := whisk.Triggers.List(options)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(triggers)
		printJSON(triggers)
	},
}

func init() {

	triggerCreateCmd.Flags().StringSliceVarP(&flags.annotation, "annotation", "a", []string{}, "annotations")
	triggerCreateCmd.Flags().StringSliceVarP(&flags.param, "param", "p", []string{}, "default parameters")
	triggerCreateCmd.Flags().BoolVar(&flags.shared, "shared", false, "shared action (default: private)")

	triggerUpdateCmd.Flags().StringSliceVarP(&flags.annotation, "annotation", "a", []string{}, "annotations")
	triggerUpdateCmd.Flags().StringSliceVarP(&flags.param, "param", "p", []string{}, "default parameters")
	triggerUpdateCmd.Flags().BoolVar(&flags.shared, "shared", false, "shared action (default: private)")

	triggerFireCmd.Flags().StringSliceVarP(&flags.param, "param", "p", []string{}, "default parameters")

	triggerListCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	triggerListCmd.Flags().IntVarP(&flags.limit, "limit", "l", 0, "only return this many entities from the collection")

	triggerCmd.AddCommand(
		triggerFireCmd,
		triggerCreateCmd,
		triggerUpdateCmd,
		triggerGetCmd,
		triggerDeleteCmd,
		triggerListCmd,
	)

}
