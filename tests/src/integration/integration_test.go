// +build integration

package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/openwhisk/openwhisk-cli/tests/src/integration/common"
	"os"
)

var wsk *common.Wsk = common.NewWsk()
var tmpProp = os.Getenv("GOPATH") + "/src/github.com/openwhisk/openwhisk-cli/wskprops.tmp"

func TestSetAPIHostAuthNamespace(t *testing.T) {
	common.CreateFile(tmpProp)
	common.WriteFile(tmpProp, []string{})

	os.Setenv("WSK_CONFIG_FILE", tmpProp)
	assert.Equal(t, os.Getenv("WSK_CONFIG_FILE"), tmpProp, "The environment variable WSK_CONFIG_FILE has not been set to de_DE.")

	//expectedNamespace, _ := wsk.ListNamespaces()
	//fmt.Println(string(expectedNamespace))
	wskprops := common.GetWskprops()
	if (wskprops.APIHost != "" && wskprops.APIHost != "") {
		stdout, err := wsk.RunCommand("property", "set", "-i", "--apihost", wskprops.APIHost, "--auth", wskprops.AuthKey)
		//,
		//	"--namespace", expectedNamespace)
		ouputString := string(stdout)
		assert.Equal(t, nil, err, "The command property set --apihost --auth --namespace failed to run.")
		assert.Contains(t, ouputString, "ok: whisk auth set to " + wskprops.AuthKey,
			"The output of the command property get --apiversion does not contain \"whisk auth key setting\".")
		assert.Contains(t, ouputString, "ok: whisk API host set to " + wskprops.APIHost,
			"The output of the command property get --apiversion does not contain \"whisk API host setting\".")
		//assert.Contains(t, ouputString, "ok: whisk namespace set to " + expectedNamespace,
		//	"The output of the command does not contain \"whisk namespace setting\".")
	}
	common.DeleteFile(tmpProp)
}
