// +build native

package tests

import (
	"testing"
	"os"
	"github.com/stretchr/testify/assert"
	"github.com/openwhisk/openwhisk-cli/tests/src/integration/common"
)

var wsk *common.Wsk = common.NewWsk()

func TestHelpUsageInfoCommand(t *testing.T) {
	stdout, err := wsk.RunCommand("-h")
	assert.Equal(t, nil, err, "The command -h failed to run.")
	assert.Contains(t, string(stdout), "Usage:", "The output of the command -h does not contain \"Usage\".")
	assert.Contains(t, string(stdout), "Flags:", "The output of the command -h does not contain \"Flags\".")
	assert.Contains(t, string(stdout), "Available Commands:",
		"The output of the command -h does not contain \"Available Commands\".")
	assert.Contains(t, string(stdout), "--help", "The output of the command -h does not contain \"--help\".")
}

func TestHelpUsageInfoCommandLanguage(t *testing.T) {
	os.Setenv("LANG", "de_DE")
	assert.Equal(t, "de_DE", os.Getenv("LANG"), "The environment variable LANG has not been set to de_DE.")
	TestHelpUsageInfoCommand(t)
}

func TestShowCLIBuildVersion(t *testing.T) {
	stdout, err := wsk.RunCommand("property", "get", "--cliversion")
	assert.Equal(t, nil, err, "The command property get --cliversion failed to run.")
	assert.Contains(t, string(stdout), "whisk CLI version",
		"The output of the command property get --cliversion does not contain \"whisk CLI version\".")
}

func TestShowAPIVersion(t *testing.T) {
	stdout, err := wsk.RunCommand("property", "get", "--apiversion")
	assert.Equal(t, nil, err, "The command property get --apiversion failed to run.")
	assert.Contains(t, string(stdout), "whisk API version",
		"The output of the command property get --apiversion does not contain \"whisk API version\".")
}




