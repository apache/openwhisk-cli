package commands

import (
	"errors"
	"fmt"

	"github.ibm.com/Bluemix/go-whisk/whisk"

	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "work with rules",
}

var ruleEnableCmd = &cobra.Command{
	Use:   "enable <name string>",
	Short: "enable rule",

	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, _, err = client.Rules.SetState(ruleName, "enable")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: enabled rule ", ruleName)

	},
}
var ruleDisableCmd = &cobra.Command{
	Use:   "disable <name string>",
	Short: "disable rule",

	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, _, err = client.Rules.SetState(ruleName, "disable")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: disabled rule ", ruleName)

	},
}

var ruleStatusCmd = &cobra.Command{
	Use:   "status <name string>",
	Short: "get rule status",

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: how is this different than "rule get" ??
		fmt.Println("rule status called")
	},
}

var ruleCreateCmd = &cobra.Command{
	Use:   "create <name string> <trigger string> <action string>",
	Short: "create new rule",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 3 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		ruleName := args[0]
		triggerName := args[1]
		actionName := args[2]

		rule := &whisk.Rule{
			Name:    ruleName,
			Trigger: triggerName,
			Action:  actionName,
			Publish: flags.common.shared,
		}

		rule, _, err = client.Rules.Insert(rule, false)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: created rule ", ruleName)
		printJSON(rule)
	},
}

var ruleUpdateCmd = &cobra.Command{
	Use:   "update <name string> <trigger string> <action string>",
	Short: "update existing rule",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 3 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		ruleName := args[0]
		triggerName := args[1]
		actionName := args[2]

		rule := &whisk.Rule{
			Name:    ruleName,
			Trigger: triggerName,
			Action:  actionName,
			Publish: flags.common.shared,
		}

		rule, _, err = client.Rules.Insert(rule, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: updated rule ", ruleName)
		printJSON(rule)
	},
}

var ruleGetCmd = &cobra.Command{
	Use:   "get <name string>",
	Short: "get rule",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		rule, _, err := client.Rules.Get(ruleName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: got rule ", ruleName)
		printJSON(rule)
	},
}

var ruleDeleteCmd = &cobra.Command{
	Use:   "delete <name string>",
	Short: "delete rule",

	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, err = client.Rules.Delete(ruleName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: deleted rule ", ruleName)
	},
}

var ruleListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all rules",

	Run: func(cmd *cobra.Command, args []string) {

		ruleListOptions := &whisk.RuleListOptions{
			Skip:  flags.common.skip,
			Limit: flags.common.limit,
		}

		rules, _, err := client.Rules.List(ruleListOptions)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("rules")
		printJSON(rules)
	},
}

func init() {

	ruleCreateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")
	ruleCreateCmd.Flags().BoolVar(&flags.rule.auto, "auto", false, "autmatically enable rule after creating it")

	ruleUpdateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	ruleDeleteCmd.Flags().BoolVar(&flags.rule.auto, "auto", false, "autmatically disable rule before deleting it")

	ruleListCmd.Flags().IntVarP(&flags.common.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	ruleListCmd.Flags().IntVarP(&flags.common.limit, "limit", "l", 30, "only return this many entities from the collection")

	ruleCmd.AddCommand(
		ruleEnableCmd,
		ruleDisableCmd,
		ruleStatusCmd,
		ruleCreateCmd,
		ruleUpdateCmd,
		ruleGetCmd,
		ruleDeleteCmd,
		ruleListCmd,
	)

}
