package commands

import (
	"errors"
	"fmt"

	"github.ibm.com/BlueMix-Fabric/go-whisk/whisk"

	"github.com/fatih/color"
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
		fmt.Printf("%s enabled rule %s\n", color.GreenString("ok:"), boldString(ruleName))
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
		fmt.Printf("%s disabled rule %s\n", color.GreenString("ok:"), boldString(ruleName))
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

		if flags.rule.enable {
			rule, _, err = client.Rules.SetState(ruleName, "enabled")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		fmt.Printf("%s created rule %s\n", color.GreenString("ok:"), boldString(ruleName))
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
		fmt.Printf("%s updated rule %s\n", color.GreenString("ok:"), boldString(ruleName))
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
		fmt.Printf("%s got rule %s\n", color.GreenString("ok:"), boldString(ruleName))
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

		if flags.rule.disable {
			_, _, err := client.Rules.SetState(ruleName, "disabled")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

		_, err = client.Rules.Delete(ruleName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s deleted rule %s\n", color.GreenString("ok:"), boldString(ruleName))
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
		printList(rules)
	},
}

func init() {

	ruleCreateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")
	ruleCreateCmd.Flags().BoolVar(&flags.rule.enable, "enable", false, "autmatically enable rule after creating it")

	ruleUpdateCmd.Flags().BoolVar(&flags.common.shared, "shared", false, "shared action (default: private)")

	ruleDeleteCmd.Flags().BoolVar(&flags.rule.disable, "disable", false, "autmatically disable rule before deleting it")

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
