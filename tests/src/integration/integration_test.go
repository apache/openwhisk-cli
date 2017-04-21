// +build integration

package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/openwhisk/openwhisk-cli/tests/src/integration/common"
	"os"
	"strings"
)

var wsk *common.Wsk = common.NewWsk()
var tmpProp = os.Getenv("GOPATH") + "/src/github.com/openwhisk/openwhisk-cli/wskprops.tmp"

// Test case to set apihost, auth, and namespace.
func TestSetAPIHostAuthNamespace(t *testing.T) {
	common.CreateFile(tmpProp)
	common.WriteFile(tmpProp, []string{})

	os.Setenv("WSK_CONFIG_FILE", tmpProp)
	assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set.")

	namespace, _ := wsk.ListNamespaces()
	namespaces := strings.Split(strings.TrimSpace(string(namespace)), "\n")
	expectedNamespace := string(namespaces[len(namespaces)-1])
	if (wsk.Wskprops.APIHost != "" && wsk.Wskprops.APIHost != "") {
		stdout, err := wsk.RunCommand("property", "set", "-i", "--apihost", wsk.Wskprops.APIHost,
			"--auth", wsk.Wskprops.AuthKey, "--namespace", expectedNamespace)
		ouputString := string(stdout)
		assert.Equal(t, nil, err, "The command property set --apihost --auth --namespace failed to run.")
		assert.Contains(t, ouputString, "ok: whisk auth set to " + wsk.Wskprops.AuthKey,
			"The output of the command property set --apihost --auth --namespace does not contain \"whisk auth key setting\".")
		assert.Contains(t, ouputString, "ok: whisk API host set to " + wsk.Wskprops.APIHost,
			"The output of the command property set --apihost --auth --namespace does not contain \"whisk API host setting\".")
		assert.Contains(t, ouputString, "ok: whisk namespace set to " + expectedNamespace,
			"The output of the command property set --apihost --auth --namespace does not contain \"whisk namespace setting\".")
	}
	common.DeleteFile(tmpProp)
}
