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
    "testing"
    "github.com/stretchr/testify/assert"
    "github.com/apache/incubator-openwhisk-cli/tests/src/integration/common"
    "os"
    "strings"
)

var invalidArgs []common.InvalidArg
var invalidArgsMsg = "error: Invalid argument(s)"
var tooFewArgsMsg = invalidArgsMsg + "."
var tooManyArgsMsg = invalidArgsMsg + ": "
var actionNameActionReqMsg = "An action name and action are required."
var actionNameReqMsg = "An action name is required."
var actionOptMsg = "An action is optional."
var packageNameReqMsg = "A package name is required."
var packageNameBindingReqMsg = "A package name and binding name are required."
var ruleNameReqMsg = "A rule name is required."
var ruleTriggerActionReqMsg = "A rule, trigger and action name are required."
var activationIdReq = "An activation ID is required."
var triggerNameReqMsg = "A trigger name is required."
var optNamespaceMsg = "An optional namespace is the only valid argument."
var optPayloadMsg = "A payload is optional."
var noArgsReqMsg = "No arguments are required."
var invalidArg = "invalidArg"
var apiCreateReqMsg = "Specify a swagger file or specify an API base path with an API path, an API verb, and an action name."
var apiGetReqMsg = "An API base path or API name is required."
var apiDeleteReqMsg = "An API base path or API name is required.  An optional API relative path and operation may also be provided."
var apiListReqMsg = "Optional parameters are: API base path (or API name), API relative path and operation."
var invalidShared = "Cannot use value '" + invalidArg + "' for shared"

func initInvalidArgs() {
    invalidArgs = []common.InvalidArg{
        common.InvalidArg {
            Cmd: []string{"api-experimental", "create"},
            Err: tooFewArgsMsg + " " + apiCreateReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"api-experimental", "create", "/basepath", "/path", "GET", "action", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + apiCreateReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"api-experimental", "get"},
            Err: tooFewArgsMsg + " " + apiGetReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"api-experimental", "get", "/basepath", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + apiGetReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"api-experimental", "delete"},
            Err: tooFewArgsMsg + " " + apiDeleteReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"api-experimental", "delete", "/basepath", "/path", "GET", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + apiDeleteReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"api-experimental", "list", "/basepath", "/path", "GET", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + apiListReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "create"},
            Err: tooFewArgsMsg + " " + actionNameActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "create", "someAction"},
            Err: tooFewArgsMsg + " " + actionNameActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "create", "actionName", "artifactName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"action", "update"},
            Err: tooFewArgsMsg + " " + actionNameReqMsg + " " + actionOptMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "update", "actionName", "artifactName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + actionNameReqMsg + " " + actionOptMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "delete"},
            Err: tooFewArgsMsg + " " + actionNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "delete", "actionName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"action", "get"},
            Err: tooFewArgsMsg + " " + actionNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "get", "actionName", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"action", "list", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "invoke"},
            Err: tooFewArgsMsg + " " + actionNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"action", "invoke", "actionName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"activation", "list", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },
        common.InvalidArg {
            Cmd: []string{"activation", "get"},
            Err: tooFewArgsMsg + " " + activationIdReq,
        },
        common.InvalidArg {
            Cmd: []string{"activation", "get", "activationID", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"activation", "logs"},
            Err: tooFewArgsMsg + " " + activationIdReq,
        },
        common.InvalidArg {
            Cmd: []string{"activation", "logs", "activationID", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },

        common.InvalidArg {
            Cmd: []string{"activation", "result"},
            Err: tooFewArgsMsg + " " + activationIdReq,
        },
        common.InvalidArg {
            Cmd: []string{"activation", "result", "activationID", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"activation", "poll", "activationID", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },
        common.InvalidArg {
            Cmd: []string{"namespace", "list", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + noArgsReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"namespace", "get", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "create"},
            Err: tooFewArgsMsg + " " + packageNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "create", "packageName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"package", "create", "packageName", "--shared", invalidArg},
            Err: invalidShared,
        },
        common.InvalidArg {
            Cmd: []string{"package", "update"},
            Err: tooFewArgsMsg + " " + packageNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "update", "packageName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"package", "update", "packageName", "--shared", invalidArg},
            Err: invalidShared,
        },
        common.InvalidArg {
            Cmd: []string{"package", "get"},
            Err: tooFewArgsMsg + " " +packageNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "get", "packageName", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"package", "bind"},
            Err: tooFewArgsMsg + " " + packageNameBindingReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "bind", "packageName"},
            Err: tooFewArgsMsg + " " +packageNameBindingReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "bind", "packageName", "bindingName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"package", "list", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "delete"},
            Err: tooFewArgsMsg + " " + packageNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"package", "delete", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"package", "refresh", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "enable"},
            Err: tooFewArgsMsg + " " + ruleNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "enable", "ruleName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"rule", "disable"},
            Err: tooFewArgsMsg + " " + ruleNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "disable", "ruleName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"rule", "status"},
            Err: tooFewArgsMsg + " " + ruleNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "status", "ruleName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"rule", "create"},
            Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "create", "ruleName"},
            Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "create", "ruleName", "triggerName"},
            Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "create", "ruleName", "triggerName", "actionName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },

        common.InvalidArg {
            Cmd: []string{"rule", "update"},
            Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "update", "ruleName"},
            Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "update", "ruleName", "triggerName"},
            Err: tooFewArgsMsg + " " + ruleTriggerActionReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "update", "ruleName", "triggerName", "actionName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"rule", "get"},
            Err: tooFewArgsMsg + " " + ruleNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "get", "ruleName", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"rule", "delete"},
            Err: tooFewArgsMsg + " " + ruleNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"rule", "delete", "ruleName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"rule", "list", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },

        common.InvalidArg {
            Cmd: []string{"trigger", "fire"},
            Err: tooFewArgsMsg + " " + triggerNameReqMsg + " " + optPayloadMsg,
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "fire", "triggerName", "triggerPayload", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + triggerNameReqMsg + " " +optPayloadMsg,
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "create"},
            Err: tooFewArgsMsg + " " + triggerNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "create", "triggerName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "update"},
            Err: tooFewArgsMsg + " " + triggerNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "update", "triggerName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },

        common.InvalidArg {
            Cmd: []string{"trigger", "get"},
            Err: tooFewArgsMsg + " " + triggerNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "get", "triggerName", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "delete"},
            Err: tooFewArgsMsg + " " + triggerNameReqMsg,
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "delete", "triggerName", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ".",
        },
        common.InvalidArg {
            Cmd: []string{"trigger", "list", "namespace", invalidArg},
            Err: tooManyArgsMsg + invalidArg + ". " + optNamespaceMsg,
        },
    }
}

var wsk *common.Wsk = common.NewWsk()
var tmpProp = common.GetRepoPath() + "/wskprops.tmp"

// Test case to set apihost, auth, and namespace.
func TestSetAPIHostAuthNamespace(t *testing.T) {
    common.CreateFile(tmpProp)
    common.WriteFile(tmpProp, []string{})

    os.Setenv("WSK_CONFIG_FILE", tmpProp)
    assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

    namespace, _ := wsk.ListNamespaces()
    namespaces := strings.Split(strings.TrimSpace(string(namespace)), "\n")
    expectedNamespace := string(namespaces[len(namespaces) - 1])
    if (wsk.Wskprops.APIHost != "" && wsk.Wskprops.APIHost != "") {
        stdout, err := wsk.RunCommand("property", "set", "--apihost", wsk.Wskprops.APIHost,
            "--auth", wsk.Wskprops.AuthKey, "--namespace", expectedNamespace)
        ouputString := string(stdout)
        assert.Equal(t, nil, err, "The command property set --apihost --auth --namespace failed to run.")
        assert.Contains(t, ouputString, "ok: whisk auth set. Run 'wsk property get --auth' to see the new value.",
            "The output of the command property set --apihost --auth --namespace does not contain \"whisk auth set\".")
        assert.Contains(t, ouputString, "ok: whisk API host set to " + wsk.Wskprops.APIHost,
            "The output of the command property set --apihost --auth --namespace does not contain \"whisk API host set\".")
        assert.Contains(t, ouputString, "ok: whisk namespace set to " + expectedNamespace,
            "The output of the command property set --apihost --auth --namespace does not contain \"whisk namespace set\".")
    }
    common.DeleteFile(tmpProp)
}

// Test case to show api build version using property file.
func TestShowAPIBuildVersion(t *testing.T) {
    common.CreateFile(tmpProp)

    os.Setenv("WSK_CONFIG_FILE", tmpProp)
    assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

    stdout, err := wsk.RunCommand("property", "set", "--apihost", wsk.Wskprops.APIHost,
        "--apiversion", wsk.Wskprops.APIVersion)
    assert.Equal(t, nil, err, "The command property set --apihost --apiversion failed to run.")
    stdout, err = wsk.RunCommand("property", "get", "-i", "--apibuild")
    assert.Equal(t, nil, err, "The command property get -i --apibuild failed to run.")
    assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), "whisk API build Unknown",
        "The output of the command property get --apibuild does not contain \"whisk API build Unknown\".")
    assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), "Unable to obtain API build information",
        "The output of the command property get --apibuild does not contain \"Unable to obtain API build information\".")
    assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), "whisk API build 20",
        "The output of the command property get --apibuild does not contain \"whisk API build 20\".")
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
    assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), "whisk API build Unknown",
        "The output of the command property get --apibuild does not contain \"whisk API build Unknown\".")
    assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), "Unable to obtain API build information",
        "The output of the command property get --apibuild does not contain \"Unable to obtain API build information\".")
}

// Test case to show api build using http apihost.
func TestShowAPIBuildVersionHTTP(t *testing.T) {
    common.CreateFile(tmpProp)

    os.Setenv("WSK_CONFIG_FILE", tmpProp)
    assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

    apihost := "http://" + wsk.Wskprops.ControllerHost + ":" + wsk.Wskprops.ControllerPort
    stdout, err := wsk.RunCommand("property", "set", "--apihost", apihost)
    assert.Equal(t, nil, err, "The command property set --apihost failed to run.")
    stdout, err = wsk.RunCommand("property", "get", "-i", "--apibuild")
    assert.Equal(t, nil, err, "The command property get -i --apibuild failed to run.")
    assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), "whisk API build Unknown",
        "The output of the command property get --apibuild does not contain \"whisk API build Unknown\".")
    assert.NotContains(t, common.RemoveRedundentSpaces(string(stdout)), "Unable to obtain API build information",
        "The output of the command property get --apibuild does not contain \"Unable to obtain API build information\".")
    assert.Contains(t, common.RemoveRedundentSpaces(string(stdout)), "whisk API build 20",
        "The output of the command property get --apibuild does not contain \"whisk API build 20\".")
    common.DeleteFile(tmpProp)
}

// Test case to reject bad command.
func TestRejectAuthCommNoKey(t *testing.T) {
    common.CreateFile(tmpProp)

    os.Setenv("WSK_CONFIG_FILE", tmpProp)
    assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

    stdout, err := wsk.RunCommand("list", "--apihost", wsk.Wskprops.APIHost,
        "--apiversion", wsk.Wskprops.APIVersion)
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
        assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1.")
        assert.Contains(t, outputString, invalidArg.Err,
            "The output of the command does not contain " + invalidArg.Err)
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
        common.InvalidArg{
            Cmd: []string{"action", "create", "actionName", helloFile},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"action", "update", "actionName", helloFile},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"action", "invoke", "actionName"},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"package", "create", "packageName"},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"package", "update", "packageName"},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"package", "bind", "packageName", "boundPackageName"},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"trigger", "create", "triggerName"},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"trigger", "update", "triggerName"},
            Err: invalidParamArg,
        },
        common.InvalidArg{
            Cmd: []string{"trigger", "fire", "triggerName"},
            Err: invalidParamArg,
        },
    }

    var annotCmds = []common.InvalidArg{
        common.InvalidArg{
            Cmd: []string{"action", "create", "actionName", helloFile},
            Err: invalidAnnoArg,
        },
        common.InvalidArg{
            Cmd: []string{"action", "update", "actionName", helloFile},
            Err: invalidAnnoArg,
        },
        common.InvalidArg{
            Cmd: []string{"package", "create", "packageName"},
            Err: invalidAnnoArg,
        },
        common.InvalidArg{
            Cmd: []string{"package", "update", "packageName"},
            Err: invalidAnnoArg,
        },
        common.InvalidArg{
            Cmd: []string{"package", "bind", "packageName", "boundPackageName"},
            Err: invalidAnnoArg,
        },
        common.InvalidArg{
            Cmd: []string{"trigger", "create", "triggerName"},
            Err: invalidAnnoArg,
        },
        common.InvalidArg{
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
                "The output of the command does not contain " + cmd.Err + " .")
        }
        for _, invalid := range invalidJSONFiles {
            cs := cmd.Cmd
            cs = append(cs, "-P", invalid, "--apihost", wsk.Wskprops.APIHost)
            stdout, err := wsk.RunCommand(cs...)
            outputString := string(stdout)
            assert.NotEqual(t, nil, err, "The command should fail to run.")
            assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1.")
            assert.Contains(t, outputString, cmd.Err,
                "The output of the command does not contain " + cmd.Err + " .")
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
                "The output of the command does not contain " + cmd.Err + " .")
        }
        for _, invalid := range invalidJSONFiles {
            cs := cmd.Cmd
            cs = append(cs, "-A", invalid, "--apihost", wsk.Wskprops.APIHost)
            stdout, err := wsk.RunCommand(cs...)
            outputString := string(stdout)
            assert.NotEqual(t, nil, err, "The command should fail to run.")
            assert.Equal(t, "exit status 1", err.Error(), "The error should be exit status 1.")
            assert.Contains(t, outputString, cmd.Err,
                "The output of the command does not contain " + cmd.Err + " .")
        }
    }
}
