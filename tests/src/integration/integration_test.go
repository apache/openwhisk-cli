// +build integration

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

package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/apache/openwhisk-cli/tests/src/integration/common"
	"github.com/stretchr/testify/assert"
)

var invalidArgs []common.InvalidArg
var invalidArgsMsg = "error: Invalid argument(s)"
var tooFewArgsMsg = invalidArgsMsg + "."
var tooManyArgsMsg = invalidArgsMsg + ": "
var actionNameActionReqMsg = "An action name and code artifact are required."
var actionNameReqMsg = "An action name is required."
var actionOptMsg = "A code artifact is optional."
var packageNameReqMsg = "A package name is required."
var packageNameBindingReqMsg = "A package name and binding name are required."
var ruleNameReqMsg = "A rule name is required."
var ruleTriggerActionReqMsg = "A rule, trigger and action name are required."
var activationIdReq = "An activation ID is required."
var triggerNameReqMsg = "A trigger name is required."
var optNamespaceMsg = "An optional namespace is the only valid argument."
var noArgsReqMsg = "No arguments are required."
var invalidArg = "invalidArg"
var apiCreateReqMsg = "Specify a swagger file or specify an API base path with an API path, an API verb, and an action name."
var apiGetReqMsg = "An API base path or API name is required."
var apiDeleteReqMsg = "An API base path or API name is required.  An optional API relative path and operation may also be provided."
var apiListReqMsg = "Optional parameters are: API base path (or API name), API relative path and operation."
var invalidShared = "Cannot use value '" + invalidArg + "' for shared"

func initInvalidArgs() {
	invalidArgs = []common.InvalidArg{
		{
			Cmd: []string{"action", "create"},
			Err: tooFewArgsMsg + " " + actionNameActionReqMsg,
		},
		{
			Cmd: []string{"action", "create", "someAction"},
			Err: tooFewArgsMsg + " " + actionNameActionReqMsg,
		},
		{
			Cmd: []string{"action", "create", "actionName", "artifactName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"action", "update"},
			Err: tooFewArgsMsg + " " + actionNameReqMsg + " " + actionOptMsg,
		},
		{
			Cmd: []string{"action", "update", "actionName", "artifactName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + actionNameReqMsg + " " + actionOptMsg,
		},
		{
			Cmd: []string{"action", "delete"},
			Err: tooFewArgsMsg + " " + actionNameReqMsg,
		},
		{
			Cmd: []string{"action", "delete", "actionName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"action", "get"},
			Err: tooFewArgsMsg + " " + actionNameReqMsg,
		},
		{
			Cmd: []string{"action", "get", "actionName", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"action", "list", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
		},
		{
			Cmd: []string{"action", "invoke"},
			Err: tooFewArgsMsg + " " + actionNameReqMsg,
		},
		{
			Cmd: []string{"action", "invoke", "actionName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"activation", "list", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
		},
		{
			Cmd: []string{"activation", "get"},
			Err: tooFewArgsMsg + " " + activationIdReq,
		},
		{
			Cmd: []string{"activation", "get", "activationID", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"activation", "logs"},
			Err: tooFewArgsMsg + " " + activationIdReq,
		},
		{
			Cmd: []string{"activation", "logs", "activationID", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},

		{
			Cmd: []string{"activation", "result"},
			Err: tooFewArgsMsg + " " + activationIdReq,
		},
		{
			Cmd: []string{"activation", "result", "activationID", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"activation", "poll", "activationID", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
		},
		{
			Cmd: []string{"namespace", "list", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + noArgsReqMsg,
		},
		{
			Cmd: []string{"namespace", "get", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + noArgsReqMsg,
		},
		{
			Cmd: []string{"package", "create"},
			Err: tooFewArgsMsg + " " + packageNameReqMsg,
		},
		{
			Cmd: []string{"package", "create", "packageName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"package", "create", "packageName", "--shared", invalidArg},
			Err: invalidShared,
		},
		{
			Cmd: []string{"package", "update"},
			Err: tooFewArgsMsg + " " + packageNameReqMsg,
		},
		{
			Cmd: []string{"package", "update", "packageName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"package", "update", "packageName", "--shared", invalidArg},
			Err: invalidShared,
		},
		{
			Cmd: []string{"package", "get"},
			Err: tooFewArgsMsg + " " + packageNameReqMsg,
		},
		{
			Cmd: []string{"package", "get", "packageName", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"package", "bind"},
			Err: tooFewArgsMsg + " " + packageNameBindingReqMsg,
		},
		{
			Cmd: []string{"package", "bind", "packageName"},
			Err: tooFewArgsMsg + " " + packageNameBindingReqMsg,
		},
		{
			Cmd: []string{"package", "bind", "packageName", "bindingName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"package", "list", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
		},
		{
			Cmd: []string{"package", "delete"},
			Err: tooFewArgsMsg + " " + packageNameReqMsg,
		},
		{
			Cmd: []string{"package", "delete", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"package", "refresh", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
		},
		{
			Cmd: []string{"rule", "enable"},
			Err: tooFewArgsMsg + " " + ruleNameReqMsg,
		},
		{
			Cmd: []string{"rule", "enable", "ruleName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"rule", "disable"},
			Err: tooFewArgsMsg + " " + ruleNameReqMsg,
		},
		{
			Cmd: []string{"rule", "disable", "ruleName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"rule", "status"},
			Err: tooFewArgsMsg + " " + ruleNameReqMsg,
		},
		{
			Cmd: []string{"rule", "status", "ruleName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"rule", "create"},
			Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
		},
		{
			Cmd: []string{"rule", "create", "ruleName"},
			Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
		},
		{
			Cmd: []string{"rule", "create", "ruleName", "triggerName"},
			Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
		},
		{
			Cmd: []string{"rule", "create", "ruleName", "triggerName", "actionName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},

		{
			Cmd: []string{"rule", "update"},
			Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
		},
		{
			Cmd: []string{"rule", "update", "ruleName"},
			Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
		},
		{
			Cmd: []string{"rule", "update", "ruleName", "triggerName"},
			Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
		},
		{
			Cmd: []string{"rule", "update", "ruleName", "triggerName", "actionName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"rule", "get"},
			Err: tooFewArgsMsg + " " + ruleNameReqMsg,
		},
		{
			Cmd: []string{"rule", "get", "ruleName", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"rule", "delete"},
			Err: tooFewArgsMsg + " " + ruleNameReqMsg,
		},
		{
			Cmd: []string{"rule", "delete", "ruleName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"rule", "list", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
		},

		{
			Cmd: []string{"trigger", "fire"},
			Err: tooFewArgsMsg + " " + triggerNameReqMsg,
		},
		{
			Cmd: []string{"trigger", "fire", "triggerName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + triggerNameReqMsg,
		},
		{
			Cmd: []string{"trigger", "create"},
			Err: tooFewArgsMsg + " " + triggerNameReqMsg,
		},
		{
			Cmd: []string{"trigger", "create", "triggerName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"trigger", "update"},
			Err: tooFewArgsMsg + " " + triggerNameReqMsg,
		},
		{
			Cmd: []string{"trigger", "update", "triggerName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},

		{
			Cmd: []string{"trigger", "get"},
			Err: tooFewArgsMsg + " " + triggerNameReqMsg,
		},
		{
			Cmd: []string{"trigger", "get", "triggerName", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"trigger", "delete"},
			Err: tooFewArgsMsg + " " + triggerNameReqMsg,
		},
		{
			Cmd: []string{"trigger", "delete", "triggerName", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ".",
		},
		{
			Cmd: []string{"trigger", "list", "namespace", invalidArg},
			Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
		},
	}
}

var wsk *common.Wsk = common.NewWsk()
var tmpProp = common.GetRepoPath() + "/wskprops.tmp"

// Test case to set apihost and auth.
func TestSetAPIHostAuthNamespace(t *testing.T) {
	common.CreateFile(tmpProp)
	common.WriteFile(tmpProp, []string{})

	os.Setenv("WSK_CONFIG_FILE", tmpProp)
	assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

	fmt.Println(wsk.Wskprops.APIHost)
	if wsk.Wskprops.APIHost != "" && wsk.Wskprops.AuthKey != "" {
		stdout, err := wsk.RunCommand("property", "set", "--apihost", wsk.Wskprops.APIHost, "--auth", wsk.Wskprops.AuthKey)
		ouputString := string(stdout)
		assert.Equal(t, nil, err, "The command property set --apihost --auth failed to run.")
		assert.Contains(t, ouputString, "ok: whisk auth set. Run 'wsk property get --auth' to see the new value.",
			"The output of the command property set --apihost --auth does not contain \"whisk auth set\".")
		assert.Contains(t, ouputString, "ok: whisk API host set to "+wsk.Wskprops.APIHost,
			"The output of the command property set --apihost --auth does not contain \"whisk API host set\".")
	}
	common.DeleteFile(tmpProp)
}

// Test case to show api build version using property file.
func TestShowAPIBuildVersion(t *testing.T) {
	common.CreateFile(tmpProp)

	os.Setenv("WSK_CONFIG_FILE", tmpProp)
	assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

	stdout, err := wsk.RunCommand("property", "set", "--apihost", wsk.Wskprops.APIHost,
		"--apiversion", wsk.Wskprops.Apiversion)
	assert.Equal(t, nil, err, "The command property set --apihost --apiversion failed to run.")
	stdout, err = wsk.RunCommand("property", "get", "-i", "--apibuild")
	assert.Equal(t, nil, err, "The command property get -i --apibuild failed to run.")
	println(common.RemoveRedundentSpaces(string(stdout)))
	assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), common.PropDisplayAPIBuild+" Unknown",
		"The output of the command property get --apibuild does not contain "+common.PropDisplayAPIBuild+" Unknown")
	assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), "Unable to obtain API build information",
		"The output of the command property get --apibuild does not contain \"Unable to obtain API build information\".")
	assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), common.PropDisplayAPIBuild+" 20",
		"The output of the command property get --apibuild does not contain"+common.PropDisplayAPIBuild+" 20")
	common.DeleteFile(tmpProp)
}

// Test case to fail to show api build when setting apihost to bogus value.
func TestFailShowAPIBuildVersion(t *testing.T) {
	common.CreateFile(tmpProp)

	os.Setenv("WSK_CONFIG_FILE", tmpProp)
	assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

	_, err := wsk.RunCommand("property", "set", "--apihost", "xxxx.yyyy")
	assert.Equal(t, nil, err, "The command property set --apihost failed to run.")
	stdout, err := wsk.RunCommand("property", "get", "-i", "--apibuild")
	assert.NotEqual(t, nil, err, "The command property get -i --apibuild does not raise any error.")
	assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), common.PropDisplayAPIBuild+" Unknown",
		"The output of the command property get --apibuild does not contain"+common.PropDisplayAPIBuild+" Unknown")
	assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), "Unable to obtain API build information",
		"The output of the command property get --apibuild does not contain \"Unable to obtain API build information\".")
}

// Test case to show api build using http apihost.
func TestShowAPIBuildVersionHTTP(t *testing.T) {
	common.CreateFile(tmpProp)

	os.Setenv("WSK_CONFIG_FILE", tmpProp)
	assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

	apihost := wsk.Wskprops.APIHost
	stdout, err := wsk.RunCommand("property", "set", "--apihost", apihost)
	assert.Equal(t, nil, err, "The command property set --apihost failed to run.")
	stdout, err = wsk.RunCommand("property", "get", "-i", "--apibuild")
	println(common.RemoveRedundentSpaces(string(stdout)))
	//assert.Equal(t, nil, err, "The command property get -i --apibuild failed to run.")
	assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), common.PropDisplayAPIBuild+" Unknown",
		"The output of the command property get --apibuild does not contain "+common.PropDisplayAPIBuild+" Unknown")
	assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), "Unable to obtain API build information",
		"The output of the command property get --apibuild does not contain \"Unable to obtain API build information\".")
	assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), common.PropDisplayAPIBuild+" 20",
		"The output of the command property get --apibuild does not contain "+common.PropDisplayAPIBuild+" 20")
	common.DeleteFile(tmpProp)
}

// Test case to reject bad command.
func TestRejectAuthCommNoKey(t *testing.T) {
	common.CreateFile(tmpProp)

	os.Setenv("WSK_CONFIG_FILE", tmpProp)
	assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

	stdout, err := wsk.RunCommand("list", "--apihost", wsk.Wskprops.APIHost,
		"--apiversion", wsk.Wskprops.Apiversion)
	assert.NotEqual(t, nil, err, "The command list should fail to run.")
	assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), "usage.",
		"The output of the command does not contain \"usage.\".")
	assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), "--auth is required",
		"The output of the command does not contain \"--auth is required\".")
	common.DeleteFile(tmpProp)
}

// Test case to reject commands that are executed with invalid arguments.
func TestRejectCommInvalidArgs(t *testing.T) {
	initInvalidArgs()
	for _, invalidArg := range invalidArgs {
		cs := invalidArg.Cmd
		cs = append(cs, "--apihost", wsk.Wskprops.APIHost)
		stdout, err := wsk.RunCommand(cs...)
		outputString := string(stdout)
		assert.NotEqual(t, nil, err, "The command should fail to run.")
		if err.Error() == "exit status 1" {
			assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1 or 2.")
		} else {
			assert.Equal(t, "exit status 2", err.Error(), "The error should be exit status 1 or 2.")
		}
		assert.Contains(t, outputString, invalidArg.Err,
			"The output of the command does not contain "+invalidArg.Err)
		assert.Contains(t, outputString, "Run 'wsk --help' for usage",
			"The output of the command does not contain \"Run 'wsk --help' for usage\".")
	}
}

// Test case to reject commands that are executed with invalid JSON for annotations and parameters.
func TestRejectCommInvalidJSON(t *testing.T) {
	helloFile := common.GetTestActionFilename("hello.js")
	var invalidJSONInputs = []string{
		"{\"invalid1\": }",
		"{\"invalid2\": bogus}",
		"{\"invalid1\": \"aKey\"",
		"invalid \"string\"",
		"{\"invalid1\": [1, 2, \"invalid\"\"arr\"]}",
	}
	var invalidJSONFiles = []string{
		common.GetTestActionFilename("malformed.js"),
		common.GetTestActionFilename("invalidInput1.json"),
		common.GetTestActionFilename("invalidInput2.json"),
		common.GetTestActionFilename("invalidInput3.json"),
		common.GetTestActionFilename("invalidInput4.json"),
	}
	var invalidParamArg = "Invalid parameter argument"
	var invalidAnnoArg = "Invalid annotation argument"
	var paramCmds = []common.InvalidArg{
		{
			Cmd: []string{"action", "create", "actionName", helloFile},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"action", "update", "actionName", helloFile},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"action", "invoke", "actionName"},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"package", "create", "packageName"},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"package", "update", "packageName"},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"package", "bind", "packageName", "boundPackageName"},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"trigger", "create", "triggerName"},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"trigger", "update", "triggerName"},
			Err: invalidParamArg,
		},
		{
			Cmd: []string{"trigger", "fire", "triggerName"},
			Err: invalidParamArg,
		},
	}

	var annotCmds = []common.InvalidArg{
		{
			Cmd: []string{"action", "create", "actionName", helloFile},
			Err: invalidAnnoArg,
		},
		{
			Cmd: []string{"action", "update", "actionName", helloFile},
			Err: invalidAnnoArg,
		},
		{
			Cmd: []string{"package", "create", "packageName"},
			Err: invalidAnnoArg,
		},
		{
			Cmd: []string{"package", "update", "packageName"},
			Err: invalidAnnoArg,
		},
		{
			Cmd: []string{"package", "bind", "packageName", "boundPackageName"},
			Err: invalidAnnoArg,
		},
		{
			Cmd: []string{"trigger", "create", "triggerName"},
			Err: invalidAnnoArg,
		},
		{
			Cmd: []string{"trigger", "update", "triggerName"},
			Err: invalidAnnoArg,
		},
	}

	for _, cmd := range paramCmds {
		for _, invalid := range invalidJSONInputs {
			cs := cmd.Cmd
			cs = append(cs, "-p", "key", invalid, "--apihost", wsk.Wskprops.APIHost)
			stdout, err := wsk.RunCommand(cs...)
			outputString := string(stdout)
			assert.NotEqual(t, nil, err, "The command should fail to run.")
			assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1.")
			assert.Contains(t, outputString, cmd.Err,
				"The output of the command does not contain "+cmd.Err+" .")
		}
		for _, invalid := range invalidJSONFiles {
			cs := cmd.Cmd
			cs = append(cs, "-P", invalid, "--apihost", wsk.Wskprops.APIHost)
			stdout, err := wsk.RunCommand(cs...)
			outputString := string(stdout)
			assert.NotEqual(t, nil, err, "The command should fail to run.")
			assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1.")
			assert.Contains(t, outputString, cmd.Err,
				"The output of the command does not contain "+cmd.Err+" .")
		}
	}

	for _, cmd := range annotCmds {
		for _, invalid := range invalidJSONInputs {
			cs := cmd.Cmd
			cs = append(cs, "-a", "key", invalid, "--apihost", wsk.Wskprops.APIHost)
			stdout, err := wsk.RunCommand(cs...)
			outputString := string(stdout)
			assert.NotEqual(t, nil, err, "The command should fail to run.")
			assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1.")
			assert.Contains(t, outputString, cmd.Err,
				"The output of the command does not contain "+cmd.Err+" .")
		}
		for _, invalid := range invalidJSONFiles {
			cs := cmd.Cmd
			cs = append(cs, "-A", invalid, "--apihost", wsk.Wskprops.APIHost)
			stdout, err := wsk.RunCommand(cs...)
			outputString := string(stdout)
			assert.NotEqual(t, nil, err, "The command should fail to run.")
			assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1.")
			assert.Contains(t, outputString, cmd.Err,
				"The output of the command does not contain "+cmd.Err+" .")
		}
	}
}
