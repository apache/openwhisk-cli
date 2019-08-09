/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package commands

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/apache/openwhisk-cli/wski18n"
	"github.com/apache/openwhisk-client-go/whisk"
)

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: wski18n.T("work with namespaces"),
}

var namespaceListCmd = &cobra.Command{
	Use:           "list",
	Short:         wski18n.T("list available namespaces"),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRunE:       SetupClientConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		// add "TYPE" --> public / private

		if whiskErr := CheckArgs(args, 0, 0, "Namespace list", wski18n.T("No arguments are required.")); whiskErr != nil {
			return whiskErr
		}

		namespaces, _, err := Client.Namespaces.List()
		if err != nil {
			whisk.Debug(whisk.DbgError, "Client.Namespaces.List() error: %s\n", err)
			errStr := wski18n.T("Unable to obtain the list of available namespaces: {{.err}}",
				map[string]interface{}{"err": err})
			werr := whisk.MakeWskErrorFromWskError(errors.New(errStr), err, whisk.EXIT_CODE_ERR_NETWORK, whisk.DISPLAY_MSG, whisk.NO_DISPLAY_USAGE)
			return werr
		}
		printList(namespaces, false) // `-n` flag applies to `namespace get`, not list, so must pass value false for printList here
		return nil
	},
}

var namespaceGetCmd = &cobra.Command{
	Use:           "get",
	Short:         wski18n.T("get triggers, actions, and rules in the registry for namespace"),
	SilenceUsage:  true,
	SilenceErrors: true,
	PreRunE:       SetupClientConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error
		var namespace string = getClientNamespace()

		if !(len(args) == 1 && args[0] == "/_") {
			if whiskErr := CheckArgs(args, 0, 0, "Namespace get",
				wski18n.T("No arguments are required.")); whiskErr != nil {
				return whiskErr
			}
		}

		actions, _, err := Client.Actions.List("", &whisk.ActionListOptions{Skip: 0, Limit: 0})
		if err != nil {
			return entityListError(err, namespace, "Actions")
		}

		packages, _, err := Client.Packages.List(&whisk.PackageListOptions{Skip: 0, Limit: 0})
		if err != nil {
			return entityListError(err, namespace, "Packages")
		}

		triggers, _, err := Client.Triggers.List(&whisk.TriggerListOptions{Skip: 0, Limit: 0})
		if err != nil {
			return entityListError(err, namespace, "Triggers")
		}

		rules, _, err := Client.Rules.List(&whisk.RuleListOptions{Skip: 0, Limit: 0})
		if err != nil {
			return entityListError(err, namespace, "Rules")
		}
		//No errors, lets attempt to retrieve the status of each rule
		for index, rule := range rules {
			ruleStatus, _, err := Client.Rules.Get(rule.Name)
			if err != nil {
				errStr := wski18n.T("Unable to get status of rule '{{.name}}': {{.err}}",
					map[string]interface{}{"name": rule.Name, "err": err})
				fmt.Println(errStr)
				werr := whisk.MakeWskErrorFromWskError(errors.New(errStr), err, whisk.EXIT_CODE_ERR_GENERAL, whisk.DISPLAY_MSG, whisk.NO_DISPLAY_USAGE)
				return werr
			}
			rules[index].Status = ruleStatus.Status
		}

		fmt.Fprintf(color.Output, wski18n.T("Entities in namespace: {{.namespace}}\n",
			map[string]interface{}{"namespace": boldString(getClientNamespace())}))
		sortByName := Flags.common.nameSort
		printList(packages, sortByName)
		printList(actions, sortByName)
		printList(triggers, sortByName)
		printList(rules, sortByName)

		return nil
	},
}

func init() {
	namespaceGetCmd.Flags().BoolVarP(&Flags.common.nameSort, "name-sort", "n", false, wski18n.T("sorts a list alphabetically by entity name; only applicable within the limit/skip returned entity block"))

	namespaceCmd.AddCommand(
		namespaceListCmd,
		namespaceGetCmd,
	)
}
