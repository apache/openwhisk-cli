// +build integration

package tests

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/openwhisk/openwhisk-cli/tests/src/integration/common"
    "strconv"
)

var wskAction *common.Wsk = common.NewWsk()
var defaultAction = common.GetTestActionFilename("hello.js")
var invalidActionArgs []common.InvalidActionArg
const BAD_REQUEST = 144
const NOT_ALLOWED = 149
const MISUSE_EXIT = 2

func initInvalidActionNames() {
    invalidActionArgs = []common.InvalidActionArg{
        common.InvalidActionArg{
            Name: "",
            ErrCode: NOT_ALLOWED,
        },
        common.InvalidActionArg{
            Name: " ",
            ErrCode: BAD_REQUEST,
        },
        common.InvalidActionArg{
            Name: "hi+there",
            ErrCode: BAD_REQUEST,
        },
        common.InvalidActionArg{
            Name: "$hola",
            ErrCode: BAD_REQUEST,
        },
        common.InvalidActionArg{
            Name: "dora?",
            ErrCode: BAD_REQUEST,
        },
        common.InvalidActionArg{
            Name: "|dora|dora?",
            ErrCode: BAD_REQUEST,
        },
    }
}

// Test case to reject creating entities with invalid names.
func TestRejectEntitiesInvalidNames(t *testing.T) {
    initInvalidActionNames()
    for _, invalidArg := range invalidActionArgs {
        _, err := wskAction.CreateAction(invalidArg.Name, defaultAction)
        assert.NotEqual(t, nil, err, "The command should fail to run.")
        expectedErrCode := "exit status " + strconv.Itoa(invalidArg.ErrCode)
        assert.Equal(t, expectedErrCode, err.Error(), "The error should be " + expectedErrCode + ".")
    }
}

// Test case to reject creating with missing file.
func TestRejectMissingFile(t *testing.T) {
    stdout, err := wskAction.CreateAction("missingFile", "notfound")
    outputString := string(stdout)
    expectedErrCode := "exit status " + strconv.Itoa(MISUSE_EXIT)
    assert.Equal(t, expectedErrCode, err.Error(), "The error should be " + expectedErrCode + ".")
    assert.Contains(t, outputString, "not a valid file",
        "The output of the command does not contain \"not a valid file\".")
}

// Test case to reject action update when specified file is missing.
func TestRejectUpdateFileMissing(t *testing.T) {
    name := "updateMissingFile"
    file := common.GetTestActionFilename("hello.js")
    _, err := wskAction.CreateAction(name, file)
    assert.Equal(t, nil, err, "The file failed to be created.")
    _, err = wskAction.UpdateAction(name, "notfound")
    expectedErrCode := "exit status " + strconv.Itoa(MISUSE_EXIT)
    assert.Equal(t, expectedErrCode, err.Error(), "The error should be " + expectedErrCode + ".")
    _, err = wskAction.DeleteAction(name)
    assert.Equal(t, nil, err, "The file failed to be deleted.")
}
