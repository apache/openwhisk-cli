package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.ibm.com/Bluemix/go-whisk/whisk"

	"github.com/spf13/cobra"
)

// triggerCmd represents the trigger command
var triggerCmd = &cobra.Command{
	Use:   "trigger",
	Short: "work with triggers",
}

var triggerFireCmd = &cobra.Command{
	Use:   "fire <name string> <payload string>",
	Short: "fire trigger event",

	Run: func(cmd *cobra.Command, args []string) {

		var err error
		var triggerName, payloadArg string
		if len(args) < 1 || len(args) > 2 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		triggerName = args[0]

		payload := map[string]interface{}{}

		if len(flags.common.param) > 0 {
			parameters, err := parseParameters(flags.common.param)
			if err != nil {
				fmt.Printf("error: %s", err)
				return
			}

			for key, value := range parameters {
				payload[key] = value
			}
		}

		if len(args) == 2 {
			payloadArg = args[1]
			reader := strings.NewReader(payloadArg)
			err = json.NewDecoder(reader).Decode(&payload)
			if err != nil {
				payload["payload"] = payloadArg
			}
		}

		_, _, err = client.Triggers.Fire(triggerName, payload)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("ok: fired trigger")
	},
}

var triggerCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "create new trigger",

	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		triggerName := args[0]

		parameters, err := parseParameters(flags.common.param)
		if err != nil {
			fmt.Println(err)
			return
		}
		annotations, err := parseAnnotations(flags.common.annotation)
		if err != nil {
			fmt.Println(err)
			return
		}

		trigger := &whisk.Trigger{
			Name:        triggerName,
			Parameters:  parameters,
			Annotations: annotations,
		}

		trigger, _, err = client.Triggers.Insert(trigger, false)

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

	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		triggerName := args[0]
		parameters, err := parseParameters(flags.common.param)
		if err != nil {
			fmt.Println(err)
			return
		}
		annotations, err := parseAnnotations(flags.common.annotation)
		if err != nil {
			fmt.Println(err)
			return
		}

		trigger := &whisk.Trigger{
			Name:        triggerName,
			Parameters:  parameters,
			Annotations: annotations,
		}

		trigger, _, err = client.Triggers.Insert(trigger, true)

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

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		triggerName := args[0]

		trigger, _, err := client.Triggers.Get(triggerName)
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

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, err = client.Triggers.Delete(ruleName)
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

	Run: func(cmd *cobra.Command, args []string) {
		options := &whisk.TriggerListOptions{
			Skip:  flags.common.skip,
			Limit: flags.common.limit,
		}
		triggers, _, err := client.Triggers.List(options)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(triggers)
		printJSON(triggers)
	},
}

func init() {

	triggerCreateCmd.Flags().StringVarP(&flags.common.annotation, "annotation", "a", "", "annotations")
	triggerCreateCmd.Flags().StringVarP(&flags.common.param, "param", "p", "", "default parameters")
	triggerCreateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	triggerUpdateCmd.Flags().StringVarP(&flags.common.annotation, "annotation", "a", "", "annotations")
	triggerUpdateCmd.Flags().StringVarP(&flags.common.param, "param", "p", "", "default parameters")
	triggerUpdateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	triggerFireCmd.Flags().StringVarP(&flags.common.param, "param", "p", "", "default parameters")

	triggerListCmd.Flags().IntVarP(&flags.common.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	triggerListCmd.Flags().IntVarP(&flags.common.limit, "limit", "l", 0, "only return this many entities from the collection")

	triggerCmd.AddCommand(
		triggerFireCmd,
		triggerCreateCmd,
		triggerUpdateCmd,
		triggerGetCmd,
		triggerDeleteCmd,
		triggerListCmd,
	)

}
